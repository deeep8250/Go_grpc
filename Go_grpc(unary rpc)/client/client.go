package main

import (
	proto "Go_grpc/protoc"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	connention, err := grpc.NewClient("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err.Error())
	}
	client := proto.NewServiceClient(connention)

	//taking data from users
	var Uinput string
	fmt.Scan(&Uinput)
	var sl []string

	for _, char := range Uinput {
		sl = append(sl, string(char))
	}

	request := &proto.Request{
		Name: sl,
	}

	res, err := client.ServerReply(context.TODO(), request)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("reverse is ", res.GetName())
	defer connention.Close()
}

//:)
