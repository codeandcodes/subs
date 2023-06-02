package shared

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	pb "github.com/codeandcodes/subs/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	FsClient *firestore.Client
}

type FsUser struct {
	EmailAddress      string
	FbUserId          string
	OsUserId          string
	SquareAccessToken string
	DisplayName       string
	PhotoUrl          string
}

type UserNotFoundError string
type FirestoreError string

func (e UserNotFoundError) Error() string {
	return fmt.Sprintf("User not found in firestore! %v", string(e))
}

func (e FirestoreError) Error() string {
	return fmt.Sprintf("Something went wrong with firestore: %v", string(e))
}

// See TODO in main.go: hitting this with facebook access token should go through validation flow first.
// For now assume it's verified
func (s *UserService) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	log.Printf("Registering user: %v", in.EmailAddress)

	resp, err := s.GetUser(ctx, &pb.GetUserRequest{
		EmailAddress: in.EmailAddress,
	})
	if err != nil {
		return nil, err
	}

	// existing user
	var doc *firestore.DocumentRef
	var message string
	if resp != nil {
		message = fmt.Sprintf("Existing user found for email %v at %v. Updating user and returning.", in.EmailAddress, resp.OsUserId)
		log.Println(message)
		// update user fields token
		doc = s.FsClient.Collection("users").Doc(resp.GetOsUserId())

		_, err := doc.Set(ctx, map[string]interface{}{
			"FbUserId":     in.FbUserId,
			"EmailAddress": in.EmailAddress,
			"DisplayName":  in.DisplayName,
			"PhotoUrl":     in.PhotoUrl,
		}, firestore.MergeAll)

		if err != nil {
			// Handle any errors in an appropriate way, such as returning them.
			log.Printf("An error has occurred updating user: %s", err)
			return nil, err
		}
	} else {
		message = fmt.Sprintf("No user found with email %v. Creating from scratch.", in.EmailAddress)
		log.Println(message)
		_, _, err := s.FsClient.Collection("users").Add(context.Background(), map[string]interface{}{
			"FbUserId":     in.FbUserId,
			"EmailAddress": in.EmailAddress,
			"DisplayName":  in.DisplayName,
			"PhotoUrl":     in.PhotoUrl,
		})
		if err != nil {
			log.Printf("Failed adding user to firestore: %v", err)
			return nil, err
		}
	}

	fsUser, err := s.GetUserWithId(ctx, doc.ID)
	if err != nil {
		log.Printf("Failed retrieving just user from firestore: %v", err)
		return nil, err
	}

	return &pb.RegisterUserResponse{
		OsUserId:     doc.ID,
		EmailAddress: fsUser.EmailAddress,
		FbUserId:     fsUser.FbUserId,
		HttpResponse: &pb.HttpResponse{
			Message:    message,
			StatusCode: fmt.Sprintf("%d", http.StatusOK),
		},
	}, nil
}

func (s *UserService) AddSquareAccessToken(ctx context.Context, in *pb.AddSquareAccessTokenRequest) (*pb.AddSquareAccessTokenResponse, error) {
	log.Printf("Calling AddSquareAccessToken as %v", ctx.Value("UserId"))

	// TODO: get this from the context instead of directly from the request once auth is in place
	osUserId := fmt.Sprintf("%v", ctx.Value("UserId"))

	doc := s.FsClient.Collection("users").Doc(osUserId)

	_, err := doc.Set(ctx, map[string]interface{}{
		"SquareAccessToken": in.SquareAccessToken,
	}, firestore.MergeAll)

	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred storing square access token: %s", err)
		return nil, err
	}

	return &pb.AddSquareAccessTokenResponse{
		HttpResponse: &pb.HttpResponse{
			Message:    "Successfully associated token.",
			StatusCode: fmt.Sprintf("%d", http.StatusOK),
		},
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	iter := s.FsClient.Collection("users").Where("EmailAddress", "==", in.EmailAddress).Documents(ctx)

	dsnaps, err := iter.GetAll()
	if err != nil {
		return nil, FirestoreError(fmt.Sprintf("%v", err))
	}

	switch true {
	//user not found
	case len(dsnaps) == 0:
		log.Printf("User not found with email. Returning nil but no error.")
		return &pb.GetUserResponse{}, status.Error(codes.NotFound, fmt.Sprintf("User %v not found", in.EmailAddress))
	// return the user
	case len(dsnaps) == 1:
		var fsUser FsUser
		dsnaps[0].DataTo(&fsUser)
		return &pb.GetUserResponse{
			OsUserId: dsnaps[0].Ref.ID,
		}, nil

	// more than one user found
	default:
		for _, dsnap := range dsnaps {
			var fsUser FsUser
			dsnap.DataTo(&fsUser)
			log.Printf("Duplicate user %v email %v", fsUser.OsUserId, fsUser.EmailAddress)
		}
		return nil, FirestoreError(fmt.Sprintf("Duplicate user found for %v", in.EmailAddress))
	}
}

func (s *UserService) GetUserWithId(ctx context.Context, userId string) (*FsUser, error) {
	dsnap, err := s.FsClient.Collection("users").Doc(userId).Get(ctx)
	if err != nil {
		return nil, FirestoreError(fmt.Sprintf("%v", err))
	}

	if !dsnap.Exists() {
		return nil, UserNotFoundError(userId)
	}

	var fsUser FsUser
	dsnap.DataTo(&fsUser)
	return &fsUser, nil
}
