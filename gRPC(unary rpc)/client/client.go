package main

import (
	"context"
	"fmt"
	proto "gRPC/protoc"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := proto.NewExampleClient(conn)

	req := &proto.HelloRequest{
		Somestring: " deep",
	}

	res, err := client.ServerReply(context.TODO(), req)
	if err != nil {
		log.Fatalf("Error calling ServerReply: %v", err)
	}

	fmt.Println(res.GetReply())
}
