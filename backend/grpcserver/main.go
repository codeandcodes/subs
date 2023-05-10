package main

import (
	"context"
	"log"
	"net"

	pb "github.com/codeandcodes/subs/protos"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSubscriptionServiceServer
}

func (s *server) SetupSubscription(ctx context.Context, in *pb.SubscriptionSetupRequest) (*pb.SubscriptionSetupResponse, error) {
	// TODO: Add your setup subscription logic here

	return &pb.SubscriptionSetupResponse{Message: "SetupSubscription has been called"}, nil
}

func (s *server) GetSubscriptions(ctx context.Context, in *pb.GetSubscriptionRequest) (*pb.GetSubscriptionsResponse, error) {
	// TODO: Add your get subscriptions logic here

	return &pb.GetSubscriptionsResponse{Message: "GetSubscriptions has been called"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSubscriptionServiceServer(s, &server{})
	log.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
