package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"

	pb "github.com/codeandcodes/subs/protos"
	square "github.com/square/square-connect-go-sdk/swagger"
)

// some specific helper methods for handling square data models

// Log errors and serialize them into a single JSON string
func LogAndSerializeModelError(errs []square.ModelError) string {
	log.Printf("LogAndSerializeModelError %v", errs)
	for _, err := range errs {
		log.Printf("Error Category: %v", err.Category)
		log.Printf("Error Code: %v", err.Code)
		log.Printf("Error Detail: %v", err.Detail)
	}
	return MarshalToString(errs)
}

func MarshalToString(a any) string {
	bytes, marshallErr := json.Marshal(a)
	if marshallErr != nil {
		log.Printf("Error marshalling json: %v", marshallErr)
		return ""
	}
	return string(bytes)
}

func MapSquareCustomerToUser(customer square.Customer) *pb.User {
	return &pb.User{
		Id:           customer.Id,
		EmailAddress: customer.EmailAddress,
		CreatedAt:    customer.CreatedAt,
		GivenName:    customer.GivenName,
		FamilyName:   customer.FamilyName,
		SquareId:     customer.Id,
	}
}

func MapSquareSubscriptionToSub(subscription square.Subscription) *pb.Subscription {
	return &pb.Subscription{
		Id:                 subscription.Id,
		PlanId:             subscription.PlanId,
		CustomerId:         subscription.CustomerId,
		StartDate:          subscription.StartDate,
		ChargedThroughDate: subscription.ChargedThroughDate,
		Status:             string(*subscription.Status),
		InvoiceIds:         subscription.InvoiceIds,
		CreatedAt:          subscription.CreatedAt,
		LocationId:         subscription.LocationId,
	}
}

func MapSquareCatalogObjectToSubscriptionCatalogObject(c square.CatalogObject) *pb.SubscriptionCatalogObject {
	return &pb.SubscriptionCatalogObject{
		Id:                   c.Id,
		UpdatedAt:            c.UpdatedAt,
		SubscriptionPlanData: MapSquareSubscriptionPlanDataToSubscriptionPlanData(*c.SubscriptionPlanData),
	}
}

func MapErrorAndHttpResponseToResponse(e error, httpResponse *http.Response) *pb.HttpResponse {
	return &pb.HttpResponse{
		Message:    "",
		Status:     fmt.Sprintf("%v", httpResponse.Status),
		StatusCode: fmt.Sprintf("%v", httpResponse.StatusCode),
		Error:      e.Error(),
	}
}

func MapSquareSubscriptionPlanDataToSubscriptionPlanData(plan square.CatalogSubscriptionPlan) *pb.SubscriptionPlanData {
	if len(plan.Phases) < 1 {
		log.Printf("Error mapping square subscription plan data. Zero phases.")
		return nil
	}

	// Phase
	phase := plan.Phases[0]
	if phase.Cadence == nil {
		log.Printf("Error converting to cadence enum from Square API. Phase cadence was nil")
		return nil
	}

	var cadence pb.SubscriptionFrequency_Cadence
	cadence, err := CadenceFromString(fmt.Sprintf("%v", *phase.Cadence))
	if err != nil {
		log.Printf("Error converting to cadence enum from Square API: %v", err)
	}

	// Convert Money
	var money int64
	if phase.RecurringPriceMoney != nil {
		money = phase.RecurringPriceMoney.Amount
	}

	return &pb.SubscriptionPlanData{
		Name:   plan.Name,
		Id:     phase.Uid,
		Amount: int32(money),
		SubscriptionFrequency: &pb.SubscriptionFrequency{
			Cadence:   cadence,
			Periods:   phase.Periods,
			IsOngoing: true,
		},
	}
}

func CadenceFromString(s string) (pb.SubscriptionFrequency_Cadence, error) {
	switch s {
	case "DAILY":
		return pb.SubscriptionFrequency_DAILY, nil
	case "WEEKLY":
		return pb.SubscriptionFrequency_WEEKLY, nil
	case "EVERY_TWO_WEEKS":
		return pb.SubscriptionFrequency_EVERY_TWO_WEEKS, nil
	case "THIRTY_DAYS":
		return pb.SubscriptionFrequency_THIRTY_DAYS, nil
	case "SIXTY_DAYS":
		return pb.SubscriptionFrequency_SIXTY_DAYS, nil
	case "NINETY_DAYS":
		return pb.SubscriptionFrequency_NINETY_DAYS, nil
	case "MONTHLY":
		return pb.SubscriptionFrequency_MONTHLY, nil
	case "EVERY_TWO_MONTHS":
		return pb.SubscriptionFrequency_EVERY_TWO_MONTHS, nil
	case "QUARTERLY":
		return pb.SubscriptionFrequency_QUARTERLY, nil
	case "EVERY_FOUR_MONTHS":
		return pb.SubscriptionFrequency_EVERY_FOUR_MONTHS, nil
	case "EVERY_SIX_MONTHS":
		return pb.SubscriptionFrequency_EVERY_SIX_MONTHS, nil
	case "ANNUAL":
		return pb.SubscriptionFrequency_ANNUAL, nil
	case "EVERY_TWO_YEARS":
		return pb.SubscriptionFrequency_EVERY_TWO_YEARS, nil
	default:
		return 0, fmt.Errorf("invalid Cadence string: %s", s)
	}
}

func ValidatePayers(in *pb.SubscriptionSetupRequest) error {
	for _, payer := range in.Payer {
		// validation
		if payer.Id == "" && payer.EmailAddress == "" {
			return ValidationError("payer.Id or payer.EmailAddress is empty")
		}
	}
	return nil
}

func GetUUID() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Printf("failed to generate UUID: %v", err)
	}
	return uuid.String()
}

// Return map of plan id: []*pb.Subscription
func ArrayToMap(list []square.Subscription) map[string][]*pb.Subscription {
	planMap := make(map[string][]*pb.Subscription)
	log.Printf("%v", list)
	for _, sub := range list {
		if _, ok := planMap[sub.PlanId]; !ok {
			planMap[sub.PlanId] = make([]*pb.Subscription, 0)
		}
		mapped := MapSquareSubscriptionToSub(sub)
		planMap[sub.PlanId] = append(planMap[sub.PlanId], mapped)
	}
	log.Printf("%v", planMap)
	return planMap
}

func NewSquareClient(squareAccessToken string) *square.APIClient {
	cfg := square.NewConfiguration()
	cfg.BasePath = "https://connect.squareupsandbox.com"
	cfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", squareAccessToken))
	cfg.AddDefaultHeader("Square-Version", "2022-09-21") //go sdk is tied to this

	square_client := square.NewAPIClient(cfg)
	return square_client
}
