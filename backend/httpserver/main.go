package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	subspb "github.com/codeandcodes/subs/protos"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

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
	UserID              string `json:"user_id"`
	FacebookAccessToken string `json:"facebook_access_token"`
	Username            string `json:"username"`
	Password            string `json:"password"`
}

func loginUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the request body
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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
		if err != nil || fbResponse.ID != creds.UserID {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	} else if creds.Username != "" && creds.Password != "" {
		// If the request contains a username and password
		// TODO: authenticate the username and password against your database
		// If authentication fails, return unauthorized status
	} else {
		// If neither authentication method is provided, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create a new session and session cookie
	session, _ := store.Get(r, "onlysubs_session")

	// Set user as authenticated
	session.Values["authenticated"] = true
	session.Values["userID"] = creds.UserID
	session.Save(r, w)
}

// Main function

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gwmux := gwruntime.NewServeMux(
		gwruntime.WithMarshalerOption(gwruntime.MIMEWildcard, &gwruntime.JSONPb{OrigName: true, EmitDefaults: true}),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := subspb.RegisterSubscriptionServiceHandlerFromEndpoint(ctx, gwmux, "localhost:50051", opts)
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
	// httpmux.HandleFunc("/loginUser", loginUserHandler)

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
