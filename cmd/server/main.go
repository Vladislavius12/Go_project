package main

import (
	"log"
	"net"

	"github.com/Vladislavius12/Go_project/pkg/api"
	"github.com/Vladislavius12/Go_project/pkg/auth"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &auth.AUTHServer{}
	api.RegisterAuthServer(s, srv)
	auth.Connect()

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
