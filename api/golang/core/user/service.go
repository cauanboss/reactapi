package user

import (
	"api/core/user/dto"
	"api/core/user/model"
	"api/interface/database/mongodb"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
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

func (ur *UserRepository) FindAll() []dto.FindUser {
	ctx := context.TODO()
	users := []dto.FindUser{}
	cursor, err := ur.client.Find(ctx, bson.M{})
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
	user := model.User{Username: u.Username, Name: u.Name, Gender: u.Gender, Age: u.Age, Password: u.Password}
	ur.client.InsertOne(context.TODO(), user)
	return user
}

func (ur *UserRepository) DeleteOne(filter dto.Delete) (*mongo.DeleteResult, error) {
	result, err := ur.client.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("asd", err)
	}
	return result, nil
}
