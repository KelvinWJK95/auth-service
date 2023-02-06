package services

import (
	"context"
	"net/http"

	"Auth-Service/pkg/db"
	"Auth-Service/pkg/models"
	"Auth-Service/pkg/pb"
)

type Server struct {
	H db.Handler
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	user.Email = req.Email
	user.Password = req.Password

	s.H.DB.Create(&user)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}
