package main

import (
	"fmt"
	"net"

	"Auth-Service/pkg/db"
	"Auth-Service/pkg/pb"
	"Auth-Service/pkg/services"
	"Auth-Service/pkg/utils"

	"google.golang.org/grpc"
)

func main() {
	h := db.Init("postgres://postgres:root@localhost:5432/authdb")

	jwt := utils.JwtWrapper{
		SecretKey:       "183709ehfd",
		Issuer:          "auth-service",
		ExpirationHours: 24 * 265,
	}

	//port number put it in env
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Failed to listing:", err)
	}

	s := services.Server{
		H:   h,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("Failed to serve:", err)
	}

}
