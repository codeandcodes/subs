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
	ServiceFactory *ServiceFactory
}

// The main method responsible for setting up all customers, catalog and subscriptions
func (s *SubscriptionService) SetupSubscription(ctx context.Context, in *pb.SubscriptionSetupRequest) (*pb.SubscriptionSetupResponse, error) {
	log.Printf("Calling SetupSubscription as %v", ctx.Value("UserId"))

	// Instantiate services and validate clients
	custService, err := s.ServiceFactory.NewSquareCustomerService(ctx)
	if err != nil {
		return nil, err
	}

	catService, err := s.ServiceFactory.NewSquareCatalogService(ctx)
	if err != nil {
		return nil, err
	}

	subService, err := s.ServiceFactory.NewSquareSubscriptionService(ctx)
	if err != nil {
		return nil, err
	}

	// Call Square services

	out := &pb.SubscriptionSetupResponse{
		CustomerCreationResults:     make(map[string]*pb.CustomerCreationResult),
		CatalogCreationResult:       &pb.CatalogCreationResult{},
		SubscriptionCreationResults: make(map[string]*pb.SubscriptionCreationResult),
	}

	// Step 1: Create Customers
	err = ValidatePayers(in)
	if err != nil {
		return out, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Error in validating input: %v", err))
	}
	err = custService.SearchOrCreateCustomers(ctx, in, out)
	if err != nil {
		return out, status.Errorf(codes.Internal, fmt.Sprintf("Fatal Error in creating customers: %v", err))
	}

	// Step 2: Setup Catalog
	log.Printf("Got request %v", in)
	err = catService.CreateSubscriptionPlan(ctx, in, out)
	if err != nil {
		return out, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Fatal Error in creating catalog object subscription plan: %v", err))
	}

	// Step 3: For each Customer, create a subscription
	err = subService.CreateSubscriptions(ctx, in, out)
	if err != nil {
		return out, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Fatal Error in creating subscriptions: %v", err))
	}
	return out, nil
}

// Lists subscriptions for the given user
func (s *SubscriptionService) GetSubscriptions(ctx context.Context, in *pb.GetSubscriptionRequest) (*pb.GetSubscriptionsResponse, error) {
	log.Printf("Calling GetSubscriptions as %v", ctx.Value("UserId"))

	// Instantiate services and validate clients
	catService, err := s.ServiceFactory.NewSquareCatalogService(ctx)
	if err != nil {
		return nil, err
	}

	subService, err := s.ServiceFactory.NewSquareSubscriptionService(ctx)
	if err != nil {
		return nil, err
	}

	response := &pb.GetSubscriptionsResponse{
		Subscriptions: make(map[string]*pb.SubscriptionCatalogObject),
	}

	listCatalogResponse, _, err := catService.ListCatalog(ctx)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Error occurred while retrieving catalog objects: %v", err))
	}

	subscriptionsResponse, _, err := subService.SearchSubscriptions(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error occurred while retrieving or mapping subscriptions: %v", err))
	}

	planMap := ArrayToMap(subscriptionsResponse.Subscriptions)

	for _, c := range listCatalogResponse.Objects {
		catObject := MapSquareCatalogObjectToSubscriptionCatalogObject(c)
		catObject.Subscriptions = planMap[c.Id]
		response.Subscriptions[c.Id] = catObject
	}

	return response, nil
}
