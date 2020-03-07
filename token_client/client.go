package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	proto "../token"
)

const (
	address     = "localhost:5001"
)

func main() {
	fmt.Println("starting client")
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewAuthServiceClient(conn)

	// Contact the server and print out its response.
	name := "Auth"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if name == "Auth" {
		fmt.Println("Client calling Auth")
		r, err := c.Auth(ctx, &proto.TokenRequest{})
		if err != nil {
			log.Fatalf("could not find server: %v", err)
		}
		fmt.Println("Token is: %s", r.Token)
	} else {
		fmt.Println("no match")
	}
}

