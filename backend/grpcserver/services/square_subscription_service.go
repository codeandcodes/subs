package services

import (
	"context"
	"fmt"
	"log"
	"net/http"

	pb "github.com/codeandcodes/subs/protos"
	square "github.com/square/square-connect-go-sdk/swagger"
)

type SubscriptionCreationError string

func (e SubscriptionCreationError) Error() string {
	return fmt.Sprintf("Unexpected error in creation of subscription: %v", string(e))
}

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
	log.Printf("Calling CreateSubscriptions as %v", ctx.Value("UserId"))
	out.SubscriptionCreationResults = make(map[string]*pb.SubscriptionCreationResult)

	for custId, cust := range out.CustomerCreationResults {
		if cust.GetUser() != nil {
			ssr := SingleSubscriptionRequest{
				planId:    out.CatalogCreationResult.SubscriptionPlan.Id,
				custId:    cust.GetUser().Id,
				startDate: in.GetSubscriptionFrequency().StartDate,
			}
			res, httpResponse, err := s.createSubscription(ctx, &ssr)
			if err != nil || httpResponse.StatusCode >= 400 {
				sce := SubscriptionCreationError(fmt.Sprintf("Error creating subscription for plan %v and user %v", ssr.planId, ssr.custId))
				out.SubscriptionCreationResults[custId] = &pb.SubscriptionCreationResult{
					HttpResponse: MapErrorAndHttpResponseToResponse(sce, httpResponse),
					Subscription: nil,
				}
				continue
			}

			//TODO: DRY http response generation
			out.SubscriptionCreationResults[custId] = &pb.SubscriptionCreationResult{
				HttpResponse: &pb.HttpResponse{
					Message:    "Subscription successfully created.",
					Status:     fmt.Sprintf("%v", httpResponse.Status),
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
func (s *SquareSubscriptionService) createSubscription(ctx context.Context, req *SingleSubscriptionRequest) (square.CreateSubscriptionResponse, *http.Response, error) {
	log.Printf("Calling createSubscription as %v for customer: %v", ctx.Value("UserId"), req.custId)
	return s.Client.SubscriptionsApi.CreateSubscription(ctx, square.CreateSubscriptionRequest{
		IdempotencyKey: GetUUID(),
		PlanId:         req.planId,
		CustomerId:     req.custId,
		StartDate:      req.startDate,
		LocationId:     "L9A9KDM49WV8Y", // this location comes from developer app
	})
}

func (s *SquareSubscriptionService) SearchSubscriptions(ctx context.Context) (square.SearchSubscriptionsResponse, *http.Response, error) {
	log.Printf("Calling SearchSubscriptions as %v", ctx.Value("UserId"))
	return s.Client.SubscriptionsApi.SearchSubscriptions(ctx, square.SearchSubscriptionsRequest{
		Limit: 1000,
	})
}
