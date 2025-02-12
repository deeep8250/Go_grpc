package main

import (
	proto "Go_grpc_Server-streaming/protoc"
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to connect to the server ", err.Error())
	}

	client := proto.NewNewsModelClient(conn)
	req := proto.Request{News: "Todays news"}
	stream, err := client.TodaysNews(context.Background(), &req)
	if err != nil {
		log.Fatal("something went wrong when recieving the data from server", err.Error())
	}

	for {

		res, err := stream.Recv()

		if err == io.EOF {
			log.Println("✅ All news received for today.")
			break
		}
		if err != nil {
			log.Fatal("something went wrong while recieving the data from stream ", err.Error())
		}

		fmt.Println(res.GetHeadlines())

	}
	defer conn.Close()

}
