package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"google.golang.org/api/option"

	"github.com/codeandcodes/subs/backend/grpcserver/services"

	"github.com/codeandcodes/subs/backend/shared"
	pb "github.com/codeandcodes/subs/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gopkg.in/yaml.v2"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"

	"github.com/gorilla/securecookie"
)

var (
	cookieName  = "onlysubs_session"
	cookieStore = securecookie.New([]byte("super-secret-key"), nil)
)

type Config struct {
	Square struct {
		Environment string `yaml:"environment"`
	} `yaml:"square"`
	Firebase struct {
		ProjectId   string `yaml:"project_id"`
		PathToCreds string `yaml:"path_to_creds"`
	}
}

type Credential struct {
	Square struct {
		AccessToken string `yaml:"access_token"`
	} `yaml:"square"`
}

const SQUARE_SANDBOX = "https://connect.squareupsandbox.com"

func heartbeat() {
	seconds := 0
	for {
		time.Sleep(10 * time.Second)
		seconds += 10
		log.Printf("Server is running...%d seconds elapsed\n", seconds)
	}
}

func configureFirebase(cfg Config) (*firebase.App, error) {
	// Use the application default credentials
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: cfg.Firebase.ProjectId}
	opt := option.WithCredentialsFile(cfg.Firebase.PathToCreds)
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln(err)
	}
	return app, err
}

// Authentication
type FsSession struct {
	UserId       string
	CreationTime string
}

// Returns closure that generates unary inteceptor with context populated
func CreateUnaryInterceptor(fsClient *firestore.Client) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Don't enforce auth on RegisterUser endpoint
		if info.FullMethod == "/subs.UserService/RegisterUser" {
			return handler(ctx, req)
		}

		md, _ := metadata.FromIncomingContext(ctx)
		log.Printf("Got metadata %v", md)

		if len(md.Get("grpcgateway-cookie")) == 0 {
			return nil, status.Error(codes.Unauthenticated, "Missing session token")
		}

		if cookies := md.Get("grpcgateway-cookie"); len(cookies) > 0 {
			header := http.Header{}
			header.Add("Cookie", cookies[0])
			request := http.Request{Header: header}
			sessionId, err := getSessionIDFromCookie(&request)
			//sessionId, err := getSessionIDFromCookie(&request, sessionIdUtil)
			if err != nil {
				return nil, status.Error(codes.Unauthenticated, "Error decoding session ID")
			}
			// You have your cookie!
			// Now you can look up the session and check the user's authentication.
			log.Printf("Decoded sessionID %v from cookie %v", sessionId, cookies)

			dsnap, err := fsClient.Collection("sessions").Doc(sessionId).Get(ctx)
			if err != nil {
				return nil, status.Error(codes.Unauthenticated, fmt.Sprintf("Invalid session token %v", sessionId))
			}

			var fsSession FsSession
			dsnap.DataTo(&fsSession)

			// Note, could do some validation with CreationTime here
			ctx = context.WithValue(ctx, "UserId", fsSession.UserId)

			return handler(ctx, req)
		}
		return nil, status.Error(codes.Unauthenticated, "Missing session token")
	}
}

func getSessionIDFromCookie(r *http.Request) (string, error) {
	// Retrieve the cookie from the request
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		// Handle error (cookie not found)
		return "", err
	}

	// Decode the cookie value (session ID)
	var sessionID string
	err = cookieStore.Decode(cookieName, cookie.Value, &sessionID)
	if err != nil {
		// Handle error (failed to decode cookie)
		return "", err
	}

	return sessionID, nil
}

func main() {

	configPath := flag.String("config", "config.yaml", "path to the config file")
	credentialPath := flag.String("creds", "credential.yaml", "path to the credential file")
	flag.Parse()

	// Read config file
	data, err := os.ReadFile(*configPath)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	// Unmarshal config file into Config struct
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("failed to unmarshal config file: %v", err)
	}

	// Read credential file
	data, err = os.ReadFile(*credentialPath)
	if err != nil {
		log.Fatalf("failed to read credential file: %v", err)
	}

	// Unmarshal config file into Config struct
	var credential Credential
	if err := yaml.Unmarshal(data, &credential); err != nil {
		log.Fatalf("failed to unmarshal credential file: %v", err)
	}

	if credential.Square.AccessToken == "" {
		log.Fatalf("Square access token is not provided in the credential file")
	}

	// Configure Firebase
	fsApp, err := configureFirebase(config)
	if err != nil {
		log.Fatalln("Failed to connect to firestore project. Terminating.")
	}
	fsClient, err := fsApp.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("Connected to firestore db: %v", config.Firebase.ProjectId)
	}
	defer fsClient.Close()

	// Instantiate Services
	// Eventually this needs to happen per user request
	subscriptionService := &services.SubscriptionService{
		ServiceFactory: &services.ServiceFactory{
			FsClient: fsClient,
		},
	}
	customerService := &services.CustomerService{
		ServiceFactory: &services.ServiceFactory{
			FsClient: fsClient,
		},
	}
	userService := &shared.UserService{
		FsClient: fsClient,
	}

	// Register services and start server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var s *grpc.Server

	log.Printf("Starting grpc server in auth mode (authentication on)")
	s = grpc.NewServer(grpc.UnaryInterceptor(CreateUnaryInterceptor(fsClient)))

	pb.RegisterSubscriptionServiceServer(s, subscriptionService)
	pb.RegisterCustomerServiceServer(s, customerService)
	pb.RegisterUserServiceServer(s, userService)

	log.Println("Server is running on port 50051...")

	go heartbeat()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
