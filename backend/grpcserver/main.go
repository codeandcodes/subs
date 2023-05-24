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
	square "github.com/square/square-connect-go-sdk/swagger"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Square struct {
		Environment string `yaml:"environment"`
		AccessToken string `yaml:"access_token"`
	} `yaml:"square"`
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
	service := &services.SubscriptionService{
		CustomerService:     square_customer_service,
		CatalogService:      square_catalog_service,
		SubscriptionService: square_subscription_service,
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
