package services

import (
	"context"
	"net/http"

	"Auth-Service/pkg/db"
	"Auth-Service/pkg/models"
	"Auth-Service/pkg/pb"
	"Auth-Service/pkg/utils"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	H   db.Handler
	Jwt utils.JwtWrapper
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

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	user.Email = req.Email
	user.Password = req.Password

	token, err := s.Jwt.GenerateToken(user)

	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}
