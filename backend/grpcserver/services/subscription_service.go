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
	err := ValidatePayers(in)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Error in validating input: %v", err))
	}
	err = s.CustomerService.SearchOrCreateCustomers(ctx, in, response)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Fatal Error in creating customers: %v", err))
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
	response := &pb.GetSubscriptionsResponse{
		Subscriptions: make(map[string]*pb.SubscriptionCatalogObject),
	}

	listCatalogResponse, _, err := s.CatalogService.ListCatalog(ctx)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Error occurred while retrieving catalog objects: %v", err))
	}

	for _, c := range listCatalogResponse.Objects {

		// Phase
		phase := c.SubscriptionPlanData.Phases[0]
		var cadence pb.SubscriptionFrequency_Cadence
		if phase.Cadence != nil {
			cadence, err = CadenceFromString(fmt.Sprintf("%v", *phase.Cadence))
			if err != nil {
				log.Printf("Error converting to cadence enum from Square API: %v", err)
			}
		}

		// Convert Money
		var money int64
		if phase.RecurringPriceMoney != nil {
			money = phase.RecurringPriceMoney.Amount
		}

		response.Subscriptions[c.Id] = &pb.SubscriptionCatalogObject{
			Id:        c.Id,
			UpdatedAt: c.UpdatedAt,
			SubscriptionPlanData: &pb.SubscriptionPlanData{
				Name:   c.SubscriptionPlanData.Name,
				Id:     phase.Uid,
				Amount: int32(money),
				SubscriptionFrequency: &pb.SubscriptionFrequency{
					Cadence:   cadence,
					StartDate: "Tbd",
					Periods:   phase.Periods,
					IsOngoing: true,
				},
			},
		}
	}

	return response, nil
}
