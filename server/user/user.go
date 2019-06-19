package user

import (
	"context"

	"github.com/yukpiz/go-twirp-example/pb/user"
)

type Server struct{}

func (h *Server) GetUser(ctx context.Context, msg *user.GetUserMessage) (*user.UserResponse, error) {
	return &user.UserResponse{
		UserId:    1,
		LastName:  "yuk",
		FirstName: "piz",
		Age:       28,
		Email:     "yukpiz@gmail.com",
	}, nil
}
