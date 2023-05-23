package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

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

func validatePayers(in *pb.SubscriptionSetupRequest) error {
	for _, payer := range in.Payer {
		// validation
		if payer.Id == "" || payer.EmailAddress == "" {
			return ValidationError("payer.Id or payer.EmailAddress is empty")
		}
	}
	return nil
}

// Takes a SubscriptionSetupRequest
// For each payer, will find or create a customer record in Square API
func (s *SquareCustomerService) SearchOrCreateCustomers(ctx context.Context, in *pb.SubscriptionSetupRequest,
	response *pb.SubscriptionSetupResponse) error {

	for _, payer := range in.Payer {
		customer, httpResponse, err := s.SearchOrCreateCustomer(ctx, payer)
		if err != nil {
			log.Printf("%v", err)
		}

		defer httpResponse.Body.Close()
		bodyString := fmt.Sprintf("%+v", httpResponse)
		bodyBytes, marshallErr := json.Marshal(bodyString)
		if marshallErr != nil {
			log.Printf("Error marshalling json %v", err)
			bodyString = string(bodyBytes)
		}
		response.CustomerCreationResults[payer.EmailAddress] = &pb.CustomerCreationResult{
			User: customer,
			HttpResponse: &pb.HttpResponse{
				Message:    strings.ToValidUTF8(bodyString, ""),
				StatusCode: fmt.Sprintf("%v", httpResponse.StatusCode),
				Error:      strings.ToValidUTF8(fmt.Sprintf("%+v", err), ""),
			},
		}

	}

	return nil
}

// Search and retrieve user, or create a user if blank
func (s *SquareCustomerService) SearchOrCreateCustomer(ctx context.Context, payer *pb.User) (*pb.User, *http.Response, error) {
	foundUser, httpResponse, err := s.SearchCustomer(ctx, payer.EmailAddress)

	if err != nil {
		log.Printf("User not found %v", payer.EmailAddress)
	}

	if foundUser != nil {
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
func (s *SquareCustomerService) SearchCustomer(ctx context.Context, email_address string) (*square.Customer, *http.Response, error) {
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

// List all customers for a user
func (s *SquareCustomerService) ListCustomers(ctx context.Context) (square.ListCustomersResponse, *http.Response, error) {
	listCustomerOpts := &square.CustomersApiListCustomersOpts{}

	return s.Client.CustomersApi.ListCustomers(ctx, listCustomerOpts)
}
