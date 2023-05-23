package services

import (
	"encoding/json"
	"fmt"
	"log"

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
