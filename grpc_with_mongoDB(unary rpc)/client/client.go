package main

import (
	"context"
	"fmt"
	proto "grpc_with_mongoDB/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err.Error())
	}

	client := proto.NewUserVerificationClient(conn)

	var U_name string
	var U_pass string
	fmt.Print("Enter your User name : ")
	fmt.Scan(&U_name)
	fmt.Print("Enter your password : ")
	fmt.Scan(&U_pass)

	request := &proto.Request{
		Name:     U_name,
		Password: U_pass,
	}

	res, err := client.ServerReply(context.TODO(), request)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res)

}
