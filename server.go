package main

import (
	"net/http"

	"github.com/yukpiz/go-twirp-example/pb/user"
	usv "github.com/yukpiz/go-twirp-example/server/user"
)

func main() {
	server := &usv.Server{}
	handler := user.NewUserAPIServer(server, nil)

	http.ListenAndServe(":9991", handler)
}
