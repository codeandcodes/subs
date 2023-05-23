package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"

	pb "github.com/codeandcodes/subs/protos"
	square "github.com/square/square-connect-go-sdk/swagger"
)

type SquareCatalogService struct {
	Client *square.APIClient
}

type SquareCatalogError string

func (e SquareCatalogError) Error() string {
	return fmt.Sprintf("Error in calling Square Catalog API: %v", string(e))
}

func (s *SquareCatalogService) CreateSubscriptionPlan(ctx context.Context, in *pb.SubscriptionSetupRequest,
	response *pb.SubscriptionSetupResponse) error {
	// Create Subscription Plan
	// Single Phase, Subscription Phase based on setup request
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Printf("failed to generate UUID: %v", err)
	}

	currency := square.Currency(square.USD_Currency)
	subscription_plan := square.CatalogObjectType(square.SUBSCRIPTION_PLAN_CatalogObjectType)
	cadence := square.SubscriptionCadence(in.SubscriptionFrequency.Cadence.String())
	catalogObjectRequest := square.UpsertCatalogObjectRequest{
		IdempotencyKey: uuid.String(),
		Object: &square.CatalogObject{
			Type_: &subscription_plan,
			Id:    fmt.Sprintf("#%v", in.Name),
			SubscriptionPlanData: &square.CatalogSubscriptionPlan{
				Name: in.Description,
				Phases: []square.SubscriptionPhase{
					{
						Cadence: &cadence,
						Periods: in.SubscriptionFrequency.Periods,
						RecurringPriceMoney: &square.Money{
							Amount:   int64(in.Amount),
							Currency: &currency,
						},
					},
				},
			},
		},
	}
	createCatalogObjectResponse, httpResponse, cErr := s.Client.CatalogApi.UpsertCatalogObject(ctx, catalogObjectRequest)
	defer httpResponse.Body.Close()
	bodyString := fmt.Sprintf("%+v", httpResponse)
	bodyBytes, err := json.Marshal(bodyString)
	if err != nil {
		log.Printf("Error marshalling json %v", err)
	}
	bodyString = string(bodyBytes)

	log.Println(bodyString)

	if cErr != nil {
		log.Printf("Error occurred while calling Square API UpsertCatalogObject %+v, %+v", createCatalogObjectResponse, cErr)
	}

	response.CatalogCreationResult =
		&pb.CatalogCreationResult{
			HttpResponse: &pb.HttpResponse{
				Message:    strings.ToValidUTF8(fmt.Sprintf("%+v", createCatalogObjectResponse), ""),
				Status:     strings.ToValidUTF8(bodyString, ""),
				StatusCode: fmt.Sprintf("%v", httpResponse.StatusCode),
				Error:      strings.ToValidUTF8(fmt.Sprintf("%+v", cErr), ""),
			},
		}

	return nil
}

func (s *SquareCatalogService) ListCatalog(ctx context.Context) (square.ListCatalogResponse, *http.Response, error) {
	listCatalogOpts := &square.CatalogApiListCatalogOpts{}

	catalogResponse, httpResponse, err := s.Client.CatalogApi.ListCatalog(ctx, listCatalogOpts)
	return catalogResponse, httpResponse, err
}
