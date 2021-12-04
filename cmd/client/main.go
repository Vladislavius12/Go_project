package main

import (
	"context"
	"flag"
	"log"

	"github.com/Vladislavius12/Go_project/pkg/api"
	"google.golang.org/grpc"
)

func main() {

	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("not enough arguments")
	}

	x := flag.Arg(0)

	y := flag.Arg(1)

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewAuthClient(conn)
	res, err := c.Auth(context.Background(), &api.AuthRequest{Login: x, Password: y})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetRes())

}
