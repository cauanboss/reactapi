package mongodb

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

func createClient() (*mongo.Collection, error) {
	co := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err := mongo.Connect(ctx, co)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Error")
	}

	var e = client.Ping(ctx, nil)

	if e != nil {
		fmt.Println(e)
		return nil, errors.New("Error")
	}

	if collection != nil {
		fmt.Print("already connected!!")
		return collection, nil
	}
	collection = client.Database("reactapi").Collection("users")
	fmt.Print("OK")
	return collection, err
}

func GetClient() *mongo.Collection {

	c, err := createClient()
	if err != nil {
		log.Fatal("Err", err)
		return nil
	}

	return c
}
