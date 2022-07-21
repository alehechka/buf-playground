package server

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	session "github.com/alehechka/buf-playground/proto/gen/go/session/v1alpha1"
)

// GetUser retrieves a random user from the UserService.
func (s *SessionServiceServer) GetUser(ctx context.Context, req *session.GetUserRequest) (*session.GetUserResponse, error) {
	userID := req.GetUserId()
	if len(userID) == 0 {
		return nil, status.Error(codes.InvalidArgument, "no user_id provided")
	}

	log.Println("Got a request to retrieve user with ID:", userID)

	return &session.GetUserResponse{User: &session.User{
		UserId:    userID,
		FirstName: "Adam",
		LastName:  "Lehechka",
		Birthday: &session.Date{
			Day:   7,
			Month: 3,
			Year:  1998,
		},
		Gender: session.Gender_GENDER_MALE,
	}}, nil
}
