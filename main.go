package main

import (
	"fmt"
	"net"

	"Auth-Service/pkg/db"
	"Auth-Service/pkg/pb"
	"Auth-Service/pkg/services"

	"google.golang.org/grpc"
)

func main() {
	h := db.Init("postgres://postgres:root@localhost:5432/testdb")

	//port number put it in env
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Failed to listing:", err)
	}

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("Failed to serve:", err)
	}

}
