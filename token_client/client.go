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
	c := proto.NewTokenServiceClient(conn)

	// Contact the server and print out its response.
	path := "Auth"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fmt.Println("Client calling Auth")
	r, err := c.Auth(ctx, &proto.TokenRequest{})
	if err != nil {
		log.Fatalf("could not find server: %v", err)
	}
	newToken := r.Token
	fmt.Println("Token is: %s", newToken)

	fmt.Println("Client calling Path")
	newPath := &proto.RequestPath{}
	newPath.Token = newToken
	newPath.Path = path
	r2, err2 := c.CheckPath(ctx, newPath)
	if err2 != nil {
		log.Fatalf("could not find server: %v", err2)
	}
	fmt.Println(r2.Count)
}

