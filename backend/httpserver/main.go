package main

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"time"

	subspb "github.com/codeandcodes/subs/protos"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
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
		log.Fatalf("Failed to register gRPC Gateway: %v", err)
	}

	httpmux := http.NewServeMux()

	// Server swagger json - in a prod app, don't serve this
	httpmux.Handle("/static/protos/api.swagger.json", enableCORS(http.StripPrefix("/static/protos", http.FileServer(http.Dir("./static/protos")))))

	httpmux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello, this is the root route!"))
	})

	httpmux.HandleFunc("/location", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello, this is the location route!"))
	})

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
