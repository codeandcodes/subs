package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	s.createCustomers(ctx, s.Client, in, response)

	// Step 2: Setup Catalog
	log.Printf("Got request %v", in)

	// Step 3: For each Customer, create a subscription
	log.Printf("response: %v", response)
	return response, nil
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

type ValidationError string

func (e ValidationError) Error() string {
	return fmt.Sprintf("Error in validating data input: %v", string(e))
}

/**
* Takes a SubscriptionSetupRequest
* Maps it to a CreateCustomerRequest
* Calls service and returns result
 */
func (s *SubscriptionService) createCustomers(ctx context.Context, client *square.APIClient,
	in *pb.SubscriptionSetupRequest, response *pb.SubscriptionSetupResponse) error {

	for _, payer := range in.Payer {
		createCustomerRequest := square.CreateCustomerRequest{
			IdempotencyKey: payer.Id,
			EmailAddress:   payer.EmailAddress,
			GivenName:      payer.GivenName,
		}

		log.Printf("Creating customer request for %v", createCustomerRequest)
		createCustomerResponse, httpResponse, cErr := client.CustomersApi.CreateCustomer(ctx, createCustomerRequest)

		defer httpResponse.Body.Close()
		bodyString := fmt.Sprintf("%+v", httpResponse)
		bodyBytes, err := json.Marshal(bodyString)
		if err != nil {
			log.Printf("Error marshalling json %v", err)
		}
		bodyString = string(bodyBytes)
		// httpResponse.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		log.Println(bodyString)

		if cErr != nil {
			log.Printf("Error occurred while calling CreateCustomer %+v, %+v", createCustomerResponse, cErr)
		}

		response.CustomerCreationResults[payer.EmailAddress] =
			&pb.CustomerCreationResult{
				HttpResponse: &pb.HttpResponse{
					Message:    strings.ToValidUTF8(fmt.Sprintf("%+v", createCustomerResponse), ""),
					Status:     strings.ToValidUTF8(bodyString, ""),
					StatusCode: fmt.Sprintf("%v", httpResponse.StatusCode),
					Error:      strings.ToValidUTF8(fmt.Sprintf("%+v", cErr), ""),
				},
			}
	}

	return nil
}

func validatePayers(in *pb.SubscriptionSetupRequest) error {
	for _, payer := range in.Payer {
		// validation
		if payer.Id == "" || payer.EmailAddress == "" {
			return ValidationError("payer.Id or payer.EmailAddress is empty")
		}
	}
	return nil
}

func (s *SubscriptionService) listCustomers(ctx context.Context, client *square.APIClient) (square.ListCustomersResponse, *http.Response, error) {
	listCustomerOpts := &square.CustomersApiListCustomersOpts{
		SortField: optional.NewString("DEFAULT"),
	}

	return client.CustomersApi.ListCustomers(ctx, listCustomerOpts)
}
