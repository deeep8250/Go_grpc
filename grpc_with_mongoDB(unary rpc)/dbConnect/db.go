package dbconnect

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

func DbConnection() {
	var err error
	uri := "mongodb+srv://deep82500:deep82500@deep.jqe1i.mongodb.net/?retryWrites=true&w=majority&appName=deep"

	clientAdress := options.Client().ApplyURI(uri)
	client, err = mongo.Connect(context.TODO(), clientAdress)
	if err != nil {
		log.Println("Error connecting to database:", err)
		return
	}
	fmt.Println("sunccessfully connected with monogDB")

}

func GetCollection() *mongo.Collection {
	if client == nil {
		log.Fatal("monogo client is not inizilized")
	}

	collection = client.Database("client_data").Collection("employee_data")
	return collection

}
