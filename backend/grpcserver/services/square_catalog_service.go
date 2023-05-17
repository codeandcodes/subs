package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	pb "github.com/codeandcodes/subs/protos"
	"github.com/jefflinse/square-connect-go-sdk/square"
)

type SquareCatalogService struct {
	Client *square.APIClient
}

func (s *SquareCatalogService) CreateSubscriptionPlan(ctx context.Context, in *pb.SubscriptionSetupRequest,
	response *pb.SubscriptionSetupResponse) error {
	// Create Subscription Plan
	// Single Phase, Subscription Phase based on setup request
	catalogObjectRequest := square.UpsertCatalogObjectRequest{
		IdempotencyKey: "<need unique key here>",
		Object: &square.CatalogObject{
			Type_: string(square.SUBSCRIPTION_PLAN_CatalogObjectType),
			Id:    "#<need some ID here note hte hashtag>",
			SubscriptionPlanData: &square.CatalogSubscriptionPlan{
				Name: "<need some name here>",
				Phases: []square.SubscriptionPhase{
					{
						Cadence: in.SubscriptionFrequency.Cadence.String(),
						Periods: in.SubscriptionFrequency.Periods,
						RecurringPriceMoney: &square.Money{
							Amount:   int64(in.Amount),
							Currency: string(square.USD_Currency),
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

func (s *SquareCatalogService) ListCatalog(ctx context.Context) error {
	listCatalogOpts := &square.CatalogApiListCatalogOpts{}

	s.Client.CatalogApi.ListCatalog(ctx, listCatalogOpts)
	return nil
}
