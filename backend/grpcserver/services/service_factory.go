package services

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/codeandcodes/subs/backend/shared"
)

type ServiceFactory struct {
	FsClient *firestore.Client
}

func (s *ServiceFactory) NewSquareCustomerService(ctx context.Context) (*SquareCustomerService, error) {
	fsUser, err := s.validateSquareUser(ctx)
	if err != nil {
		return nil, err
	}

	return &SquareCustomerService{
		Client: NewSquareClient(fsUser.SquareAccessToken),
	}, nil
}

func (s *ServiceFactory) NewSquareCatalogService(ctx context.Context) (*SquareCatalogService, error) {
	fsUser, err := s.validateSquareUser(ctx)
	if err != nil {
		return nil, err
	}

	return &SquareCatalogService{
		Client: NewSquareClient(fsUser.SquareAccessToken),
	}, nil
}

func (s *ServiceFactory) NewSquareSubscriptionService(ctx context.Context) (*SquareSubscriptionService, error) {
	fsUser, err := s.validateSquareUser(ctx)
	if err != nil {
		return nil, err
	}

	return &SquareSubscriptionService{
		Client: NewSquareClient(fsUser.SquareAccessToken),
	}, nil
}

// Validates that user is present in DB and has a square access token
func (s *ServiceFactory) validateSquareUser(ctx context.Context) (*shared.FsUser, error) {
	us := shared.UserService{
		FsClient: s.FsClient,
	}

	fsUser, err := us.GetUser(ctx, fmt.Sprintf("%v", ctx.Value("UserId")))
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, fmt.Sprintf("Error occurred while retrieving contextual user: %v", err))
	}
	if fsUser.SquareAccessToken == "" {
		return nil, status.Errorf(codes.PermissionDenied,
			fmt.Sprintf("User %v has no associated square access token. Cannot call square services.", fmt.Sprintf("%v", ctx.Value("UserId"))))
	}
	return fsUser, nil
}
