package services

import (
	"context"
	"fmt"
	"log"
	"net/http"

	pb "github.com/codeandcodes/subs/protos"
	square "github.com/square/square-connect-go-sdk/swagger"
)

type SquareCustomerService struct {
	Client *square.APIClient
}

type ValidationError string
type CustomerCreationError string

func (e ValidationError) Error() string {
	return fmt.Sprintf("Error in validating data input: %v", string(e))
}

func (e CustomerCreationError) Error() string {
	return fmt.Sprintf("Unexpected error in creation of customer: %v", string(e))
}

// Takes a SubscriptionSetupRequest
// For each payer, will find or create a customer record in Square API
func (s *SquareCustomerService) SearchOrCreateCustomers(ctx context.Context, in *pb.SubscriptionSetupRequest,
	response *pb.SubscriptionSetupResponse) error {
	log.Printf("Calling SearchOrCreateCustomers as %v", ctx.Value("UserId"))

	for _, payer := range in.Payer {
		customer, httpResponse, err := s.searchOrCreateCustomer(ctx, payer)
		if err != nil || httpResponse.StatusCode >= 400 {
			cce := CustomerCreationError(fmt.Sprintf("Error searching or creating customer: %v", err))
			response.CustomerCreationResults[payer.Id] = &pb.CustomerCreationResult{
				User:         nil,
				HttpResponse: MapErrorAndHttpResponseToResponse(cce, httpResponse),
			}
		}

		response.CustomerCreationResults[payer.Id] = &pb.CustomerCreationResult{
			User: customer,
			HttpResponse: &pb.HttpResponse{
				Message:    "Successfully created or found customer.",
				Status:     fmt.Sprintf("%v", httpResponse.Status),
				StatusCode: fmt.Sprintf("%v", httpResponse.StatusCode),
				Error:      "",
			},
		}

	}

	return nil
}

// Search and retrieve user, or create a user if blank
func (s *SquareCustomerService) searchOrCreateCustomer(ctx context.Context, payer *pb.User) (*pb.User, *http.Response, error) {
	log.Printf("Calling searchOrCreateCustomer as %v on %v", ctx.Value("UserId"), payer.EmailAddress)
	foundUser, httpResponse, err := s.searchCustomer(ctx, payer.EmailAddress)
	if err != nil {
		log.Printf("User not found %v", payer.EmailAddress)
	}

	if foundUser != nil {
		log.Printf("User found, mapping customer for %v", payer.EmailAddress)
		return MapSquareCustomerToUser(*foundUser), httpResponse, nil
	}

	createCustomerRequest := square.CreateCustomerRequest{
		IdempotencyKey: payer.Id,
		EmailAddress:   payer.EmailAddress,
		GivenName:      payer.GivenName,
		FamilyName:     payer.FamilyName,
	}

	log.Printf("Since user not found, creating customer for %v", payer.EmailAddress)
	createCustomerResponse, httpResponse, cErr := s.Client.CustomersApi.CreateCustomer(ctx, createCustomerRequest)

	if cErr != nil {
		return nil, httpResponse, cErr
	}

	if createCustomerResponse.Customer == nil {
		return nil, httpResponse, CustomerCreationError(
			fmt.Sprintf("Unexpected error for %v. Customer in response is nil", createCustomerRequest.IdempotencyKey))
	}

	return MapSquareCustomerToUser(*createCustomerResponse.Customer), httpResponse, nil

}

// Search for a single customer based on email.
// First searches, gets ID, then returns square customer directory customer object
func (s *SquareCustomerService) searchCustomer(ctx context.Context, email_address string) (*square.Customer, *http.Response, error) {
	log.Printf("Calling searchCustomer as %v on %v", ctx.Value("UserId"), email_address)
	searchResponse, httpResponse, err := s.Client.CustomersApi.SearchCustomers(ctx, square.SearchCustomersRequest{
		Limit: 1, // TODO: if there are more than 1, something is wrong, we should fix this
		Query: &square.CustomerQuery{
			Filter: &square.CustomerFilter{
				EmailAddress: &square.CustomerTextFilter{
					Exact: email_address,
				},
			},
		},
	})

	if err != nil {
		log.Printf("Error while searching for user %v. %v", email_address, httpResponse)
		return nil, httpResponse, err
	}

	if len(searchResponse.Customers) == 0 {
		return nil, httpResponse, nil
	}

	retrieveResponse, httpResponse, err := s.Client.CustomersApi.RetrieveCustomer(ctx, searchResponse.Customers[0].Id)

	if err != nil {
		log.Printf("Error while retrieving user %v. %v", searchResponse.Customers[0].Id, httpResponse)
		return nil, httpResponse, err
	}

	return retrieveResponse.Customer, httpResponse, nil

}

// Retrieves customer by Id
func (s *SquareCustomerService) GetCustomer(ctx context.Context, id string) (square.RetrieveCustomerResponse, *http.Response, error) {
	return s.Client.CustomersApi.RetrieveCustomer(ctx, id)
}

// List all customers for a user
func (s *SquareCustomerService) ListCustomers(ctx context.Context) (square.ListCustomersResponse, *http.Response, error) {
	listCustomerOpts := &square.CustomersApiListCustomersOpts{}

	return s.Client.CustomersApi.ListCustomers(ctx, listCustomerOpts)
}
