package main

import (
	TokenDaos "./lib"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	proto "../token"
	"context"
	"strconv"
)

const (
	port = ":5001"
)

func insertNewToken() string {
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
	proto.UnimplementedTokenServiceServer
}

func (s *server) Auth(ctx context.Context, in *proto.TokenRequest) (*proto.TokenReply, error) {
	return &proto.TokenReply{Token: insertNewToken()}, nil
}

func (s *server) CheckPath(ctx context.Context, in *proto.RequestPath) (*proto.CountReply, error) {
	if checkTokenExists(in.Token) {
		count := checkPathCount(in.Path)
		fmt.Println("count: ", count)
		s1 := strconv.FormatInt(int64(count), 10)
		fmt.Println("count String: ", s1)
		return &proto.CountReply{Count: s1}, nil
	} else {
		fmt.Println("Unathorized")
		return &proto.CountReply{Count: "Unauthorized"}, nil
	}
}

func main() {
	fmt.Println("starting server")
	listeningServer, error := net.Listen("tcp", port)
	if error != nil {
		log.Fatalf("failed to listen: %v", error)
	}
	fmt.Println("starting Auth Server")
	authServer := grpc.NewServer()
	proto.RegisterTokenServiceServer(authServer, &server{})
	if error := authServer.Serve(listeningServer); error != nil {
		log.Fatalf("failed to serve: %v", error)
		fmt.Println("#{error}")
	}

}


