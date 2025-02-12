package main

import (
	"context"
	proto "gRPC/protoc"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedExampleServer
}

func main() {
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to start the server %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterExampleServer(grpcServer, &Server{})
	log.Printf("server started at %v", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to start: %v", err)
	}

}

func (s *Server) ServerReply(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	response := "Hello" + req.GetSomestring()

	return &proto.HelloResponse{Reply: response}, nil
}
