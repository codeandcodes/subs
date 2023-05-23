package services

import (
	"encoding/json"
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
