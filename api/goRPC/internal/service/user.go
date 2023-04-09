package service

import (
	"context"
	"main/internal/database"
	"main/internal/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	UserDB database.User
}

func NewUserService(userDB database.User) *UserService {
	return &UserService{UserDB: userDB}
}

func (c *UserService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error) {
	user, err := c.UserDB.Create(in.Name, in.Email, in.Password)
	if err != nil {
		return nil, err
	}
	userResponse := &pb.User{
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return userResponse, nil
}

func (c *UserService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	user, err := c.UserDB.FindOne(in.Id)
	if err != nil {
		return nil, err
	}
	userResponse := &pb.User{
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return userResponse, nil
}

func (c *UserService) ListUsers(ctx context.Context, in *pb.Blank) (*pb.UserList, error) {
	users, err := c.UserDB.FindAll()
	if err != nil {
		return nil, err
	}
	var usersResponse []*pb.User
	for _, user := range users {
		userResponse := &pb.User{
			Id:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		}
		usersResponse = append(usersResponse, userResponse)
	}
	return &pb.UserList{Users: usersResponse}, nil
}
