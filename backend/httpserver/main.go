package main

import (
	"context"
	"log"
	"net/http"

	subspb "github.com/codeandcodes/subs/protos"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gwmux := gwruntime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := subspb.RegisterSubscriptionServiceHandlerFromEndpoint(ctx, gwmux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC Gateway: %v", err)
	}

	httpmux := http.NewServeMux()
	httpmux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello, this is the root route!"))
	})

	httpmux.HandleFunc("/location", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello, this is the location route!"))
	})

	// Combine gRPC Gateway routes and HTTP routes on the same server.
	mux := http.NewServeMux()
	mux.Handle("/", httpmux)  // non-gRPC routes
	mux.Handle("/v1/", gwmux) // gRPC routes

	// Start the HTTP server with the mux as the default handler.
	err = http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
