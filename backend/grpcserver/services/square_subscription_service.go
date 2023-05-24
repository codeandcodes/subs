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

type SquareSubscriptionService struct {
	Client *square.APIClient
}

type SingleSubscriptionRequest struct {
	planId    string
	custId    string
	startDate string
}

// Handle all subscriptions for a SubscriptionSetupRequest
// TODO: this should use an acceptance flow before we start billing
func (s *SquareSubscriptionService) CreateSubscriptions(ctx context.Context, in *pb.SubscriptionSetupRequest, out *pb.SubscriptionSetupResponse) error {

	out.SubscriptionCreationResults = make(map[string]*pb.SubscriptionCreationResult)

	for custId, cust := range out.CustomerCreationResults {
		if cust.GetUser() != nil {
			ssr := SingleSubscriptionRequest{
				planId:    out.CatalogCreationResult.SubscriptionPlan.Id,
				custId:    cust.GetUser().Id,
				startDate: in.GetSubscriptionFrequency().StartDate,
			}
			res, httpResponse, err := s.CreateSubscription(ctx, &ssr)
			if err != nil {
				log.Printf("Error creating subscription for plan %v and user %v", ssr.planId, ssr.custId)
				out.SubscriptionCreationResults[custId] = &pb.SubscriptionCreationResult{
					HttpResponse: &pb.HttpResponse{
						Message:    "Error occurred creating subscription.",
						Status:     "",
						StatusCode: fmt.Sprintf("%v", httpResponse.StatusCode),
						Error:      strings.ToValidUTF8(fmt.Sprintf("%+v", err), ""),
					},
					Subscription: nil,
				}
				continue
			}

			defer httpResponse.Body.Close()
			bodyString := fmt.Sprintf("%+v", httpResponse)
			bodyBytes, err := json.Marshal(bodyString)
			if err != nil {
				log.Printf("Error marshalling json %v", err)
			}
			bodyString = string(bodyBytes)

			if httpResponse.StatusCode >= 400 {
				out.SubscriptionCreationResults[custId] = &pb.SubscriptionCreationResult{
					HttpResponse: &pb.HttpResponse{
						Message:    "Subscription creation failed.",
						Status:     strings.ToValidUTF8(bodyString, ""),
						StatusCode: fmt.Sprintf("%v", httpResponse.StatusCode),
						Error:      "",
					},
					Subscription: nil,
				}
			}

			//TODO: DRY http response generation
			out.SubscriptionCreationResults[custId] = &pb.SubscriptionCreationResult{
				HttpResponse: &pb.HttpResponse{
					Message:    "Subscription successfully created.",
					Status:     "",
					StatusCode: fmt.Sprintf("%v", httpResponse.StatusCode),
					Error:      "",
				},
				Subscription: MapSquareSubscriptionToSub(*res.Subscription),
			}
		}
	}
	return nil
}

// Create a single subscription
func (s *SquareSubscriptionService) CreateSubscription(ctx context.Context, req *SingleSubscriptionRequest) (square.CreateSubscriptionResponse, *http.Response, error) {
	return s.Client.SubscriptionsApi.CreateSubscription(ctx, square.CreateSubscriptionRequest{
		IdempotencyKey: GetUUID(),
		PlanId:         req.planId,
		CustomerId:     req.custId,
		StartDate:      req.startDate,
		LocationId:     "L9A9KDM49WV8Y", // this location comes from developer app
	})
}
