package mongodb

import (
	"api/core/user/model"
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func createClient() (*mongo.Collection, error) {
	co := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), co)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Error")
	}

	var e = client.Ping(context.TODO(), nil)

	if e != nil {
		fmt.Println(e)
		return nil, errors.New("Error")
	}

	if collection != nil {
		return collection, nil
	}
	collection = client.Database("Udemy").Collection("users")

	filter := bson.M{"name": "Cauan"}

	u := model.User{}

	collection.FindOne(context.TODO(), filter).Decode(&u)

	fmt.Println(u)

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
