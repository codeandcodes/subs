package services

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/codeandcodes/subs/protos"
)

type CustomerService struct {
	pb.UnimplementedCustomerServiceServer
	ServiceFactory *ServiceFactory
}

// List customer information for single customer
func (s *CustomerService) GetCustomer(ctx context.Context, in *pb.GetCustomerRequest) (*pb.GetCustomerResponse, error) {
	log.Printf("Calling GetCustomer as %v", ctx.Value("UserId"))

	// Instantiate services and validate clients
	cs, err := s.ServiceFactory.NewSquareCustomerService(ctx)
	if err != nil {
		return nil, err
	}

	// Call Square services

	out := &pb.GetCustomerResponse{}

	if in.GetIncludeSubscriptions() {
		out.PaymentsFrom = make([]*pb.Subscription, 0)
		out.PaymentsTo = make([]*pb.Subscription, 0)
		//do stuff
	}

	if in.GetCustomerId() != "" {
		retrieveResponse, httpResponse, err := cs.GetCustomer(ctx, in.GetCustomerId())
		if err != nil {
			return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Error occurred while retrieving customer: %v", err))
		} else if httpResponse.StatusCode >= 400 {
			return nil, status.Errorf(codes.Code(httpResponse.StatusCode), "Error: %v", httpResponse.Status)
		}
		out.User = MapSquareCustomerToUser(*retrieveResponse.Customer)
		return out, nil
	} else if in.GetEmail() != "" {
		customer, httpResponse, err := cs.searchCustomer(ctx, in.GetEmail())
		if err != nil {
			return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Error occurred while retrieving customer: %v", err))
		} else if httpResponse.StatusCode >= 400 {
			return nil, status.Errorf(codes.Code(httpResponse.StatusCode), "Error: %v", httpResponse.Status)
		}
		out.User = MapSquareCustomerToUser(*customer)
		return out, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "Neither customer ID nor email found. Cannot fulfill request.")
}

// List all customers for a user
func (s *CustomerService) GetCustomers(ctx context.Context, in *pb.GetCustomersRequest) (*pb.GetCustomersResponse, error) {
	log.Printf("Calling GetCustomers as %v", ctx.Value("UserId"))

	// Instantiate services and validate clients
	cs, err := s.ServiceFactory.NewSquareCustomerService(ctx)
	if err != nil {
		return nil, err
	}

	// Call Square services
	listCustomersResponse, httpResponse, err := cs.ListCustomers(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while retreiving list of customers %v", err)
	} else if httpResponse.StatusCode >= 400 {
		return nil, status.Errorf(codes.Code(httpResponse.StatusCode), "Error: %v", httpResponse.Status)
	}
	payers := make([]*pb.User, 0)
	for _, v := range listCustomersResponse.Customers {
		payers = append(payers, MapSquareCustomerToUser(v))
	}
	return &pb.GetCustomersResponse{
		Users: payers,
	}, nil
}
