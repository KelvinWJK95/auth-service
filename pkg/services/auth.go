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

	if len(req.Email) == 0 {
		return &pb.RegisterResponse{
			Status: http.StatusOK,
			Error:  "Email must be provided",
		}, nil
	} else if result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusOK,
			Error:  "Email already in use",
		}, nil
	}

	user.Email = req.Email
	user.Password = req.Password

	s.H.DB.Create(&user)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	if len(req.Email) == 0 {
		return &pb.LoginResponse{
			Status: http.StatusOK,
			Error:  "Email must be provided",
		}, nil
	} else if result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error != nil {
		return &pb.LoginResponse{
			Status: http.StatusOK,
			Error:  "This email is not registered",
		}, nil
	}

	//need add has keys
	if req.Password != user.Password {
		return &pb.LoginResponse{
			Status: http.StatusOK,
			Error:  "Incorrect Password",
		}, nil
	}

	user.Email = req.Email
	user.Password = req.Password

	token, err := s.Jwt.GenerateToken(user)

	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusOK,
			Error:  err.Error(),
		}, nil
	}

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}
