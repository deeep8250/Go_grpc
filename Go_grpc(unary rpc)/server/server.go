package main

import (
	proto "Go_grpc/protoc"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedServiceServer
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err.Error())
	}

	srv := grpc.NewServer()
	proto.RegisterServiceServer(srv, &server{})
	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		log.Fatal(err.Error())
	}

}

func (s *server) ServerReply(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	//recieve the input from user requests
	userReq := req
	fmt.Println("recieving request is : ", userReq.Name)
	palindrome := ""
	for i := len(userReq.Name); i > 0; i-- {
		palindrome += userReq.Name[i-1]
	}

	return &proto.Response{Name: palindrome}, nil
}
