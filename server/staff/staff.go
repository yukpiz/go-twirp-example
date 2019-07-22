package staff

import (
	"context"

	"github.com/yukpiz/go-twirp-example/pb/staff"
)

type Server struct{}

func (s *Server) GetStaff(ctx context.Context, msg *staff.GetStaffMessage) (*staff.StaffResponse, error) {
	return &staff.StaffResponse{
		StaffId:   1,
		FirstName: "yuk",
		LastName:  "piz",
		Age:       28,
		Email:     "yukpiz@gmail.com",
	}, nil
}
