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
	Scs *SquareCustomerService
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
	s.Scs.createCustomers(ctx, in, response)

	// Step 2: Setup Catalog
	log.Printf("Got request %v", in)

	// Step 3: For each Customer, create a subscription
	log.Printf("response: %v", response)
	return response, nil
}

func (s *SubscriptionService) GetSubscriptions(ctx context.Context, in *pb.GetSubscriptionRequest) (*pb.GetSubscriptionsResponse, error) {
	// TODO: Add your get subscriptions logic here

	listCustomerResponse, httpResponse, err := s.Scs.listCustomers(ctx)
	if err != nil {
		return &pb.GetSubscriptionsResponse{Message: fmt.Sprintf("Error HTTP %v: %v", httpResponse, err)}, err
	}

	for _, v := range listCustomerResponse.Customers {
		log.Printf("Customer found: %v\n", v)
	}

	return &pb.GetSubscriptionsResponse{Message: fmt.Sprintf("%v", listCustomerResponse)}, nil
}
