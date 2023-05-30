package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v2"

	"google.golang.org/api/option"

	"github.com/codeandcodes/subs/backend/shared"
	subspb "github.com/codeandcodes/subs/protos"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"

	"github.com/gorilla/securecookie"
)

var (
	cookieName  = "onlysubs_session"
	cookieStore = securecookie.New([]byte("super-secret-key"), nil)
)

type Config struct {
	Firebase struct {
		ProjectId   string `yaml:"project_id"`
		PathToCreds string `yaml:"path_to_creds"`
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

func heartbeat() {
	seconds := 0
	for {
		time.Sleep(10 * time.Second)
		seconds += 10
		log.Printf("Server is running...%d seconds elapsed\n", seconds)

	}
}

// Middleware
type Middleware func(http.Handler) http.Handler

func ChainMiddleware(middlewares ...Middleware) Middleware {
	return func(final http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- { // iterate in reverse order
			final = middlewares[i](final)
		}
		return final
	}
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// If request is OPTIONS then just return with status OK
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Next
		next.ServeHTTP(w, r)
		return
	})
}

func echoMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Echo the request
		log.Printf("Request method: %s\n", r.Method)
		log.Printf("Request URL: %s\n", r.URL.String())
		if r.Body != nil {
			body, _ := io.ReadAll(r.Body)
			log.Printf("Request body: %s\n", string(body))
			r.Body = io.NopCloser(bytes.NewBuffer(body)) // put back the body content
		}
		// Then pass the request to the next middleware (or the final handler)
		next.ServeHTTP(w, r)
	})
}

// Authentication

type Credentials struct {
	UserId              string `json:"user_id"`
	FacebookAccessToken string `json:"facebook_access_token"`
	Username            string `json:"username"`
	Password            string `json:"password"`
}

func CreateLoginUserHandler(fsClient *firestore.Client) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		// Parse and decode the request body
		var creds Credentials
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Printf("Attempting to login user %v", creds.UserId)

		// If the request contains a Facebook access token
		if creds.FacebookAccessToken != "" {
			// Call Facebook's API to get the user ID associated with the access token
			res, err := http.Get("https://graph.facebook.com/me?access_token=" + creds.FacebookAccessToken)
			if err != nil || res.StatusCode != http.StatusOK {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// Parse the response to get the Facebook user ID
			var fbResponse struct {
				ID string `json:"id"`
			}
			err = json.NewDecoder(res.Body).Decode(&fbResponse)
			if err != nil || fbResponse.ID != creds.UserId {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		} else if creds.Username != "" && creds.Password != "" {
			log.Printf("Handling username/password authentication for %v", creds.UserId)
			// If the request contains a username and password
			// TODO: authenticate the username and password against your database
			// If authentication fails, return unauthorized status
		} else {
			// If neither authentication method is provided, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Printf("User successfully authenticated %v", creds.UserId)
		// Create a new session and session cookie
		sessionID := generateSessionID()

		// Save to DB
		log.Printf("New session token %v", sessionID)
		ss := shared.SessionService{
			FsClient: fsClient,
		}
		_, err = ss.WriteSessionToDb(context.Background(), sessionID, creds.UserId)
		if err != nil {
			log.Printf("Error writing session to Db: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Encode the session ID in a secure cookie
		encodedSessionID, err := cookieStore.Encode(cookieName, sessionID)
		if err != nil {
			log.Printf("Error encoding sessionId to cookie: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Create the cookie
		cookie := &http.Cookie{
			Name:     cookieName,
			Value:    encodedSessionID,
			Path:     "/",
			Secure:   true,                    // Enable secure flag (HTTPS)
			HttpOnly: true,                    // Disallow JavaScript access
			SameSite: http.SameSiteStrictMode, // Enforce same-site cookies
		}

		// Set the cookie in the response
		http.SetCookie(w, cookie)

	}
}

func generateSessionID() string {
	randomBytes := make([]byte, 32)
	if _, err := rand.Read(randomBytes); err != nil {
		// Handle error
	}
	return base64.URLEncoding.EncodeToString(randomBytes)
}

// Main function

func main() {

	configPath := flag.String("config", "config.yaml", "path to the config file")

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

	// Server code

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gwmux := gwruntime.NewServeMux(
		gwruntime.WithMarshalerOption(gwruntime.MIMEWildcard, &gwruntime.JSONPb{OrigName: true, EmitDefaults: true}),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err = subspb.RegisterSubscriptionServiceHandlerFromEndpoint(ctx, gwmux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Gateway Subscription Service: %v", err)
	}

	err = subspb.RegisterCustomerServiceHandlerFromEndpoint(ctx, gwmux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Gateway Customer Service: %v", err)
	}

	err = subspb.RegisterUserServiceHandlerFromEndpoint(ctx, gwmux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Gateway User Service: %v", err)
	}

	httpmux := http.NewServeMux()

	// Server swagger json - in a prod app, don't serve this
	httpmux.Handle("/static/protos/api.swagger.json", enableCORS(http.StripPrefix("/static/protos", http.FileServer(http.Dir("./static/protos")))))

	fs := http.FileServer(http.Dir("../../frontend"))

	httpmux.HandleFunc("/", fs.ServeHTTP)

	httpmux.HandleFunc("/location", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello, this is the location route!"))
	})

	// Register your loginUserHandler
	// NOTE: Comment this line out to bypass authentication
	httpmux.HandleFunc("/loginUser", CreateLoginUserHandler(fsClient))

	// Combine gRPC Gateway routes and HTTP routes on the same server.
	allMiddlewares := ChainMiddleware(enableCORS, echoMiddleware)
	mux := http.NewServeMux()

	mux.Handle("/", httpmux)                  // non-gRPC routes
	mux.Handle("/v1/", allMiddlewares(gwmux)) // gRPC routes

	go heartbeat()

	// Start the HTTP server with the mux as the default handler.
	err = http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}

}
