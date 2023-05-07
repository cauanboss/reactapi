package service

import (
	"api/core/user/model"
	"api/interface/database/mongodb"
	"api/internal/pb"
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	pb.UnimplementedUserServiceServer
	Client *mongo.Collection
}

func NewUserService(ur *mongo.Collection) *UserRepository {
	return &UserRepository{Client: ur}
}

func UserService() *UserRepository {
	return NewUserService(mongodb.GetClient())
}

func (ur *UserRepository) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error) {
	id := uuid.New().String()
	u := model.User{Name: in.Name, Username: in.Username, Password: in.Password, Age: in.Age, Gender: in.Gender}
	_, err := ur.Client.InsertOne(context.TODO(), &u)
	if err != nil {
		return nil, err
	}
	userResponse := &pb.User{
		Id:       string(id),
		Name:     u.Name,
		Username: u.Username,
		Password: u.Password,
		Age:      u.Age,
	}

	return userResponse, nil
}

func (ur *UserRepository) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	var u model.User
	err := ur.Client.FindOne(context.TODO(), model.User{Id: in.Id}).Decode(&u)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	userResponse := &pb.User{
		Id:       u.Id,
		Name:     u.Name,
		Username: u.Username,
		Password: u.Password,
		Age:      u.Age,
	}

	return userResponse, nil
}

// func (c *UserRepository) ListCategories(ctx context.Context, in *pb.Blank) (*pb.UserList, error) {
// 	categories, err := c.Client.FindAll()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var categoriesResponse []*pb.Category

// 	for _, category := range categories {
// 		categoryResponse := &pb.Category{
// 			Id:          category.ID,
// 			Name:        category.Name,
// 			Description: category.Description,
// 		}

// 		categoriesResponse = append(categoriesResponse, categoryResponse)
// 	}

// 	return &pb.CategoryList{Categories: categoriesResponse}, nil
// }
