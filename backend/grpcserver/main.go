package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/api/option"

	"github.com/codeandcodes/subs/backend/grpcserver/services"

	pb "github.com/codeandcodes/subs/protos"
	square "github.com/square/square-connect-go-sdk/swagger"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"

	firebase "firebase.google.com/go/v4"
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

	// Create Square client with access token from config file
	cfg := square.NewConfiguration()
	if config.Square.Environment == "sandbox" {
		log.Printf("Setting basepath to sandbox: %+v\n", SQUARE_SANDBOX)
		cfg.BasePath = SQUARE_SANDBOX
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

	cfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", credential.Square.AccessToken))
	cfg.AddDefaultHeader("Square-Version", "2022-09-21") //go sdk is tied to this

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

	// test
	_, _, err = fsClient.Collection("users").Add(context.Background(), map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	//fsClient.Close() //TODO: connect this up with the square_client so it's per user.

	// Instantiate Services
	// Eventually this needs to happen per user request
	square_client := square.NewAPIClient(cfg)
	square_customer_service := &services.SquareCustomerService{
		Client: square_client,
	}
	square_catalog_service := &services.SquareCatalogService{
		Client: square_client,
	}
	square_subscription_service := &services.SquareSubscriptionService{
		Client: square_client,
	}
	subscriptionService := &services.SubscriptionService{
		CustomerService:     square_customer_service,
		CatalogService:      square_catalog_service,
		SubscriptionService: square_subscription_service,
	}
	customerService := &services.CustomerService{
		CustomerService:     square_customer_service,
		CatalogService:      square_catalog_service,
		SubscriptionService: square_subscription_service,
	}
	userService := &services.UserService{
		FsClient: fsClient,
	}

	log.Printf("Square API configuration: %+v\n", cfg)

	// Register services and start server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSubscriptionServiceServer(s, subscriptionService)
	pb.RegisterCustomerServiceServer(s, customerService)
	pb.RegisterUserServiceServer(s, userService)

	log.Println("Server is running on port 50051...")

	go heartbeat()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
