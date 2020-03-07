package main

import (
	TokenDaos "./lib"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	proto "../token"
	"context"
)

const (
	port = ":5001"
)

func insertNewToken() string {
	fmt.Println("getting new token")
	return TokenDaos.CreateToken()
}

func checkTokenExists(token string) bool {
	return TokenDaos.TokenExists(token)
}

func checkPathCount(path string) int {
	return TokenDaos.PathCount(path)
}


func server_cmd_tests() {
	fmt.Println(insertNewToken())

	if checkTokenExists("1234567") {
		fmt.Println("count: ", checkPathCount("abc"))
	} else {
		fmt.Println("Unathorized")
	}

	if checkTokenExists("1") {
		fmt.Println("count: ", checkPathCount("abc"))
	} else {
		fmt.Println("Unathorized")
	}

	if checkTokenExists("1234567") {
		fmt.Println("count: ", checkPathCount("abcd"))
	} else {
		fmt.Println("Unathorized")
	}
}

type server struct {
	proto.UnimplementedAuthServiceServer
}

func (s *server) Auth(ctx context.Context, in *proto.TokenRequest) (*proto.TokenReply, error) {
	fmt.Println("hit Auth method")
	return &proto.TokenReply{Token: insertNewToken()}, nil
}

func main() {
	fmt.Println("starting server")
	listeningServer, error := net.Listen("tcp", port)
	if error != nil {
		log.Fatalf("failed to listen: %v", error)
	}
	newServer := grpc.NewServer()
	proto.RegisterAuthServiceServer(newServer, &server{})
	if error := newServer.Serve(listeningServer); error != nil {
		log.Fatalf("failed to serve: %v", error)
		fmt.Println("#{error}")
	}
	fmt.Println("started server fully")
}


