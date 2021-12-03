package main

import (
	"log"
	"net"

	"github.com/.../GoProj/pkg/api"
	"github.com/.../GoProj/pkg/auth"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &auth.AUTHServer{}
	api.RegisterAuthServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
