package services

import (
	"context"

	pb "github.com/codeandcodes/subs/protos"
	"github.com/jefflinse/square-connect-go-sdk/square"
)

type SubscriptionService struct {
	Client *square.APIClient
	pb.UnimplementedSubscriptionServiceServer
}

func (s *SubscriptionService) SetupSubscription(ctx context.Context, in *pb.SubscriptionSetupRequest) (*pb.SubscriptionSetupResponse, error) {
	// TODO: Add your setup subscription logic here

	return &pb.SubscriptionSetupResponse{Message: "SetupSubscription has been called"}, nil
}

func (s *SubscriptionService) GetSubscriptions(ctx context.Context, in *pb.GetSubscriptionRequest) (*pb.GetSubscriptionsResponse, error) {
	// TODO: Add your get subscriptions logic here

	return &pb.GetSubscriptionsResponse{Message: "GetSubscriptions has been called"}, nil
}
