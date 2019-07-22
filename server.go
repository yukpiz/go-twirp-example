package main

import (
	"log"
	"net/http"

	"github.com/yukpiz/go-twirp-example/pb/staff"
	"github.com/yukpiz/go-twirp-example/pb/user"
	sserver "github.com/yukpiz/go-twirp-example/server/staff"
	userver "github.com/yukpiz/go-twirp-example/server/user"
)

func main() {
	usv := &userver.Server{}
	uh := user.NewUserAPIServer(usv, nil)
	http.Handle(user.UserAPIPathPrefix, uh)

	ssv := &sserver.Server{}
	sh := staff.NewStaffAPIServer(ssv, nil)
	http.Handle(staff.StaffAPIPathPrefix, sh)

	err := http.ListenAndServe(":9991", nil)
	if err != nil {
		log.Fatal(err)
	}
}
