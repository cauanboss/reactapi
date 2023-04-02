package user

import (
	"api/core/user/dto"
	"api/core/user/model"
	"api/interface/database/mongodb"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	client *mongo.Collection
}

func NewUserService(ur *mongo.Collection) *UserRepository {
	return &UserRepository{ur}
}

func UserService() *UserRepository {
	return NewUserService(mongodb.GetClient())
}

func (ur *UserRepository) FindOne(filter interface{}) model.User {
	ctx := context.TODO()
	u := model.User{}
	ur.client.FindOne(ctx, filter).Decode(&u)
	return u
}

func (ur *UserRepository) FindAll() []model.User {
	ctx := context.TODO()
	users := []model.User{}
	cursor, err := ur.client.Find(ctx, model.User{})
	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &users); err != nil {
		panic(err)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return users
}

func (ur *UserRepository) CreateUser(u model.User) model.User {
	fmt.Print("CreateUser")
	ur.client.InsertOne(context.TODO(), u)
	return u
}

func (ur *UserRepository) DeleteOne(filter dto.Delete) (*mongo.DeleteResult, error) {
	result, err := ur.client.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("asd", err)
	}
	return result, nil
}
