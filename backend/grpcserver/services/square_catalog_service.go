package services

import (
	"context"
	"fmt"
	"log"
	"net/http"

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

// Create Subscription Plan
// Single Phase, Subscription Phase based on setup request
func (s *SquareCatalogService) CreateSubscriptionPlan(ctx context.Context, in *pb.SubscriptionSetupRequest,
	response *pb.SubscriptionSetupResponse) error {

	currency := square.Currency(square.USD_Currency)
	subscription_plan := square.CatalogObjectType(square.SUBSCRIPTION_PLAN_CatalogObjectType)
	cadence := square.SubscriptionCadence(in.SubscriptionFrequency.Cadence.String())
	catalogObjectRequest := square.UpsertCatalogObjectRequest{
		IdempotencyKey: GetUUID(),
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
	createCatalogObjectResponse, httpResponse, err := s.Client.CatalogApi.UpsertCatalogObject(ctx, catalogObjectRequest)

	if err != nil || httpResponse.StatusCode >= 400 {
		sce := SquareCatalogError(fmt.Sprintf("Error occurred while calling Square API UpsertCatalogObject %+v", err))
		log.Printf("%v", sce)
		response.CatalogCreationResult = &pb.CatalogCreationResult{
			HttpResponse:     MapErrorAndHttpResponseToResponse(sce, httpResponse),
			SubscriptionPlan: nil,
		}
		return sce
	}

	response.CatalogCreationResult =
		&pb.CatalogCreationResult{
			HttpResponse: &pb.HttpResponse{
				Message:    "Catalog object and subscription plan successfully created.",
				Status:     fmt.Sprintf("%v", httpResponse.Status),
				StatusCode: fmt.Sprintf("%v", httpResponse.StatusCode),
				Error:      "",
			},
			SubscriptionPlan: MapSquareCatalogObjectToSubscriptionCatalogObject(*createCatalogObjectResponse.CatalogObject),
		}

	return nil
}

func (s *SquareCatalogService) ListCatalog(ctx context.Context) (square.ListCatalogResponse, *http.Response, error) {
	listCatalogOpts := &square.CatalogApiListCatalogOpts{}

	catalogResponse, httpResponse, err := s.Client.CatalogApi.ListCatalog(ctx, listCatalogOpts)
	return catalogResponse, httpResponse, err
}
