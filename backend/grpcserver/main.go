package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/codeandcodes/subs/backend/grpcserver/services"
	pb "github.com/codeandcodes/subs/protos"
	"github.com/jefflinse/square-connect-go-sdk/square"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Square struct {
		Environment string `yaml:"environment"`
		AccessToken string `yaml:"access_token"`
	} `yaml:"square"`
}

type server struct {
	pb.UnimplementedSubscriptionServiceServer
}

func heartbeat() {
	seconds := 0
	for {
		time.Sleep(10 * time.Second)
		seconds += 10
		log.Printf("Server is running...%d seconds elapsed\n", seconds)

	}
}

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

	// Create Square client with access token from config file
	cfg := square.NewConfiguration()
	cfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", config.Square.AccessToken))

	if config.Square.AccessToken == "" {
		log.Fatalf("Square access token is not provided in the config")
	}

	// Instantiate Services
	square_client := square.NewAPIClient(cfg)
	service := &services.SubscriptionService{
		Client: square_client,
	}
	log.Printf("Square API configuration: %+v\n", cfg)

	// Register services and start server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSubscriptionServiceServer(s, service)
	log.Println("Server is running on port 50051...")

	go heartbeat()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
