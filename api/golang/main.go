package main

import (
	"api/interface/database/mongodb"
	"api/interface/routing"
	"api/internal/pb"
	"api/internal/service"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func handleRequest() {
	http.Handle("/", routing.GetRoutes())
	userDB := mongodb.GetClient()
	userService := service.NewUserService(userDB)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userService)
	reflection.Register(grpcServer)

	go http.ListenAndServe(":8101", nil)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

func main() {
	handleRequest()
}
