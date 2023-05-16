package services

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/antihax/optional"
	pb "github.com/codeandcodes/subs/protos"
	"github.com/jefflinse/square-connect-go-sdk/square"
)

type SubscriptionService struct {
	Client *square.APIClient
	pb.UnimplementedSubscriptionServiceServer
}

func (s *SubscriptionService) SetupSubscription(ctx context.Context, in *pb.SubscriptionSetupRequest) (*pb.SubscriptionSetupResponse, error) {
	// TODO: Add your setup subscription logic here

	// Step 1: Create Customers
	createCustomerResponse, httpResponse, err := s.createCustomer(ctx, s.Client, in)
	if err != nil {
		return &pb.SubscriptionSetupResponse{Message: fmt.Sprintf("Error HTTP %v: %v", httpResponse, err)}, err
	}

	// Step 2: Setup Catalog

	// Step 3: For each Customer, create a subscription

	return &pb.SubscriptionSetupResponse{Message: fmt.Sprintf("%v", createCustomerResponse)}, nil
}

func (s *SubscriptionService) GetSubscriptions(ctx context.Context, in *pb.GetSubscriptionRequest) (*pb.GetSubscriptionsResponse, error) {
	// TODO: Add your get subscriptions logic here

	listCustomerResponse, httpResponse, err := s.listCustomers(ctx, s.Client)
	if err != nil {
		return &pb.GetSubscriptionsResponse{Message: fmt.Sprintf("Error HTTP %v: %v", httpResponse, err)}, err
	}

	for _, v := range listCustomerResponse.Customers {
		log.Printf("Customer found: %v\n", v)
	}

	return &pb.GetSubscriptionsResponse{Message: fmt.Sprintf("%v", listCustomerResponse)}, nil
}

/**
* Takes a SubscriptionSetupRequest
* Maps it to a CreateCustomerRequest
* Calls service and returns result
 */
func (s *SubscriptionService) createCustomer(ctx context.Context, client *square.APIClient, request *pb.SubscriptionSetupRequest) (square.CreateCustomerResponse, *http.Response, error) {

	createCustomerRequest := square.CreateCustomerRequest{
		IdempotencyKey: "rocketegg+1@gmail.com",
		EmailAddress:   "rocketegg+1@gmail.com",
		GivenName:      "Rocket Egg",
	}

	return client.CustomersApi.CreateCustomer(ctx, createCustomerRequest)
}

func (s *SubscriptionService) listCustomers(ctx context.Context, client *square.APIClient) (square.ListCustomersResponse, *http.Response, error) {
	listCustomerOpts := &square.CustomersApiListCustomersOpts{
		SortField: optional.NewString("DEFAULT"),
	}

	return client.CustomersApi.ListCustomers(ctx, listCustomerOpts)
}
