package main

import (
	"database/sql"
	"main/internal/database"
	"main/internal/pb"
	"main/internal/service"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userDB := database.NewUser(db)
	userService := service.NewUserService(*userDB)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
