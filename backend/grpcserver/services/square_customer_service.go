package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/antihax/optional"
	pb "github.com/codeandcodes/subs/protos"
	"github.com/jefflinse/square-connect-go-sdk/square"
)

type SquareCustomerService struct {
	Client          *square.APIClient
	CreateCustomers func(context.Context, *pb.SubscriptionSetupRequest, *pb.SubscriptionSetupResponse) error
	ListCustomers   func(context.Context, *pb.SubscriptionSetupRequest, *pb.SubscriptionSetupResponse) (square.ListCustomersResponse, *http.Response, error)
}

type ValidationError string

func (e ValidationError) Error() string {
	return fmt.Sprintf("Error in validating data input: %v", string(e))
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

/**
* Takes a SubscriptionSetupRequest
* Maps it to a CreateCustomerRequest
* Calls service and returns result
 */
func (s *SquareCustomerService) createCustomers(ctx context.Context, in *pb.SubscriptionSetupRequest,
	response *pb.SubscriptionSetupResponse) error {

	for _, payer := range in.Payer {
		createCustomerRequest := square.CreateCustomerRequest{
			IdempotencyKey: payer.Id,
			EmailAddress:   payer.EmailAddress,
			GivenName:      payer.GivenName,
		}

		log.Printf("Creating customer request for %v", createCustomerRequest)
		createCustomerResponse, httpResponse, cErr := s.Client.CustomersApi.CreateCustomer(ctx, createCustomerRequest)

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

func (s *SquareCustomerService) listCustomers(ctx context.Context) (square.ListCustomersResponse, *http.Response, error) {
	listCustomerOpts := &square.CustomersApiListCustomersOpts{
		SortField: optional.NewString("DEFAULT"),
	}

	return s.Client.CustomersApi.ListCustomers(ctx, listCustomerOpts)
}
