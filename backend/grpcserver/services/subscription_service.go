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
	CustomerService     *SquareCustomerService
	CatalogService      *SquareCatalogService
	SubscriptionService *SquareSubscriptionService
}

// The main method responsible for setting up all customers, catalog and subscriptions
func (s *SubscriptionService) SetupSubscription(ctx context.Context, in *pb.SubscriptionSetupRequest) (*pb.SubscriptionSetupResponse, error) {
	out := &pb.SubscriptionSetupResponse{
		CustomerCreationResults:     make(map[string]*pb.CustomerCreationResult),
		CatalogCreationResult:       &pb.CatalogCreationResult{},
		SubscriptionCreationResults: make(map[string]*pb.SubscriptionCreationResult),
	}

	// Step 1: Create Customers
	err := ValidatePayers(in)
	if err != nil {
		return out, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Error in validating input: %v", err))
	}
	err = s.CustomerService.SearchOrCreateCustomers(ctx, in, out)
	if err != nil {
		return out, status.Errorf(codes.Internal, fmt.Sprintf("Fatal Error in creating customers: %v", err))
	}

	// Step 2: Setup Catalog
	log.Printf("Got request %v", in)
	err = s.CatalogService.CreateSubscriptionPlan(ctx, in, out)
	if err != nil {
		return out, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Fatal Error in creating catalog object subscription plan: %v", err))
	}

	// Step 3: For each Customer, create a subscription
	err = s.SubscriptionService.CreateSubscriptions(ctx, in, out)
	if err != nil {
		return out, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Fatal Error in creating subscriptions: %v", err))
	}
	return out, nil
}

// Lists subscriptions for the given user
func (s *SubscriptionService) GetSubscriptions(ctx context.Context, in *pb.GetSubscriptionRequest) (*pb.GetSubscriptionsResponse, error) {
	response := &pb.GetSubscriptionsResponse{
		Subscriptions: make(map[string]*pb.SubscriptionCatalogObject),
	}

	listCatalogResponse, _, err := s.CatalogService.ListCatalog(ctx)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Error occurred while retrieving catalog objects: %v", err))
	}

	for _, c := range listCatalogResponse.Objects {
		response.Subscriptions[c.Id] = MapSquareCatalogObjectToSubscriptionCatalogObject(c)
	}

	return response, nil
}
