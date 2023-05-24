package services

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/codeandcodes/subs/protos"
	square "github.com/square/square-connect-go-sdk/swagger"
)

type CustomerService struct {
	pb.UnimplementedCustomerServiceServer
	CustomerService     *SquareCustomerService
	CatalogService      *SquareCatalogService
	SubscriptionService *SquareSubscriptionService
}

// The main method responsible for setting up all customers, catalog and subscriptions
func (s *CustomerService) GetCustomer(ctx context.Context, in *pb.GetCustomerRequest) (*pb.GetCustomerResponse, error) {
	log.Printf("Calling GetCustomer on %v", in)

	out := &pb.GetCustomerResponse{}

	if in.GetIncludeSubscriptions() {
		out.PaymentsFrom = make([]*pb.Subscription, 0)
		out.PaymentsTo = make([]*pb.Subscription, 0)
		//do stuff
	}

	if in.GetCustomerId() != "" {
		retrieveResponse, _, err := s.CustomerService.GetCustomer(ctx, in.GetCustomerId())
		if err != nil {
			return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Error occurred while retrieving customer: %v", err))
		}
		out.User = MapSquareCustomerToUser(*retrieveResponse.Customer)
		return out, nil
	} else if in.GetEmail() != "" {
		customer, _, err := s.CustomerService.searchCustomer(ctx, in.GetEmail())
		if err != nil {
			return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Error occurred while retrieving customer: %v", err))
		}
		out.User = MapSquareCustomerToUser(*customer)
		return out, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "Neither customer ID nor email found. Cannot fulfill request.")
}

// List all customers for a user
func (s *CustomerService) ListCustomers(ctx context.Context) (square.ListCustomersResponse, *http.Response, error) {
	return s.CustomerService.ListCustomers(ctx)
}
