package main

import (
	"context"
	dbconnect "grpc_with_mongoDB/dbConnect"
	proto "grpc_with_mongoDB/proto"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedUserVerificationServer
}

func main() {
	dbconnect.DbConnection()
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err.Error())
	}

	srv := grpc.NewServer()
	proto.RegisterUserVerificationServer(srv, &server{})

	if err := srv.Serve(listen); err != nil {
		log.Fatal(err.Error())
	}

}

func (s *server) ServerReply(ctx context.Context, req *proto.Request) (*proto.Response, error) {

	var coll = dbconnect.GetCollection()

	filter := bson.M{
		"name":     req.Name,
		"password": req.Password,
	}

	err := coll.FindOne(ctx, filter).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			_, err := coll.InsertOne(ctx, filter)
			if err != nil {
				log.Fatal(err.Error())
			}
			return &proto.Response{Reply: "Not a verified user but this data is successfully added into db"}, nil
		} else {
			return &proto.Response{Reply: err.Error()}, nil
		}
	}
	return &proto.Response{Reply: "Verified user"}, nil

}
