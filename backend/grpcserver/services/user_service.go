package services

import (
	"context"

	"cloud.google.com/go/firestore"
	pb "github.com/codeandcodes/subs/protos"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	FsClient *firestore.Client
}

func (s *UserService) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	return nil, nil
}

func (s *UserService) AddSquareAccessToken(ctx context.Context, in *pb.AddSquareAccessTokenRequest) (*pb.AddSquareAccessTokenResponse, error) {
	return nil, nil
}
