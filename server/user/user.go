package user

import (
	"context"

	"github.com/yukpiz/go-twirp-example/pb/user"
)

type Server struct{}

func (s *Server) GetUser(ctx context.Context, msg *user.GetUserMessage) (*user.UserResponse, error) {

	// ...

	return &user.UserResponse{
		UserId:    1,
		LastName:  "yuk",
		FirstName: "piz",
		Age:       28,
		Email:     "yukpiz@gmail.com",
	}, nil
}

func (s *Server) CreateUser(ctx context.Context, msg *user.CreateUserMessage) (*user.UserResponse, error) {

	// ...

	return &user.UserResponse{
		UserId:    1,
		LastName:  "yuk",
		FirstName: "piz",
		Age:       28,
		Email:     "yukpiz@gmail.com",
	}, nil
}

func (s *Server) UpdateUser(ctx context.Context, msg *user.UpdateUserMessage) (*user.UserResponse, error) {

	// ...

	return &user.UserResponse{
		UserId:    1,
		LastName:  "yuk",
		FirstName: "piz",
		Age:       28,
		Email:     "yukpiz@gmail.com",
	}, nil
}

func (s *Server) DeleteUser(ctx context.Context, msg *user.DeleteUserMessage) (*user.EmptyResponse, error) {

	// ...

	return nil, nil
}
