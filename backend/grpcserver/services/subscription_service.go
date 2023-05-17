package services

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/codeandcodes/subs/protos"
)

type SubscriptionService struct {
	pb.UnimplementedSubscriptionServiceServer
	CustomerService *SquareCustomerService
	CatalogService  *SquareCatalogService
}

func (s *SubscriptionService) SetupSubscription(ctx context.Context, in *pb.SubscriptionSetupRequest) (*pb.SubscriptionSetupResponse, error) {
	// TODO: Add your setup subscription logic here

	response := &pb.SubscriptionSetupResponse{
		CustomerCreationResults:     make(map[string]*pb.CustomerCreationResult),
		CatalogCreationResult:       &pb.CatalogCreationResult{},
		SubscriptionCreationResults: make(map[string]*pb.SubscriptionCreationResult),
	}

	// Step 1: Create Customers
	err := validatePayers(in)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Error in validating input: %v", err))
	}
	err = s.CustomerService.CreateCustomers(ctx, in, response)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Fatal Error in creating customers: %v", err))
	}

	// Step 2: Setup Catalog
	log.Printf("Got request %v", in)
	err = s.CatalogService.CreateSubscriptionPlan(ctx, in, response)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Fatal Error in creating catalog object subscription plan: %v", err))
	}

	// Step 3: For each Customer, create a subscription
	log.Printf("response: %v", response)
	return response, nil
}

func (s *SubscriptionService) GetSubscriptions(ctx context.Context, in *pb.GetSubscriptionRequest) (*pb.GetSubscriptionsResponse, error) {
	// TODO: Add your get subscriptions logic here

	listCustomerResponse, httpResponse, err := s.CustomerService.ListCustomers(ctx)
	if err != nil {
		return &pb.GetSubscriptionsResponse{Message: fmt.Sprintf("Error HTTP %v: %v", httpResponse, err)}, err
	}

	for _, v := range listCustomerResponse.Customers {
		log.Printf("Customer found: %v\n", v)
	}

	return &pb.GetSubscriptionsResponse{Message: fmt.Sprintf("%v", listCustomerResponse)}, nil
}
