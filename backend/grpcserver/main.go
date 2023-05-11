package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

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

func (s *server) SetupSubscription(ctx context.Context, in *pb.SubscriptionSetupRequest) (*pb.SubscriptionSetupResponse, error) {
	// TODO: Add your setup subscription logic here

	return &pb.SubscriptionSetupResponse{Message: "SetupSubscription has been called"}, nil
}

func (s *server) GetSubscriptions(ctx context.Context, in *pb.GetSubscriptionRequest) (*pb.GetSubscriptionsResponse, error) {
	// TODO: Add your get subscriptions logic here

	return &pb.GetSubscriptionsResponse{Message: "GetSubscriptions has been called"}, nil
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

	log.Printf("Square API configuration: %+v\n", cfg)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSubscriptionServiceServer(s, &server{})
	log.Println("Server is running on port 50051...")

	go heartbeat()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
