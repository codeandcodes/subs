package shared

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	pb "github.com/codeandcodes/subs/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SessionService struct {
	pb.UnimplementedUserServiceServer
	FsClient *firestore.Client
}

// Authentication
type FsSession struct {
	UserId       string
	CreationTime string
}

func (s *SessionService) GetSessionForUser(ctx context.Context, sessionId string) (*FsSession, error) {
	dsnap, err := s.FsClient.Collection("sessions").Doc(sessionId).Get(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, fmt.Sprintf("Invalid session token %v", sessionId))
	}

	var fsSession FsSession
	dsnap.DataTo(&fsSession)
	return &fsSession, nil
}

func (s *SessionService) WriteSessionToDb(ctx context.Context, sessionKey string, userId string) (*firestore.WriteResult, error) {
	docRef := s.FsClient.Collection("sessions").Doc(sessionKey)

	writeResult, err := docRef.Set(context.Background(), map[string]interface{}{
		"UserId":       userId,
		"CreationTime": time.Now().Format(time.RFC3339),
	})

	if err != nil {
		return nil, err
	}
	log.Printf("Wrote sessionId %v to Db at %v", sessionKey, writeResult.UpdateTime)
	return writeResult, nil
}
