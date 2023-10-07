package handler

import (
	"context"
	pb "github.com/blazee5/cloud-drive/microservices/auth/api/v1"
	"github.com/blazee5/cloud-drive/microservices/auth/internal/service"
)

type Server struct {
	service *service.Service
	pb.UnimplementedAuthServiceServer
}

func NewServer(service *service.Service) *Server {
	return &Server{service: service}
}

func (s *Server) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.UserResponse, error) {
	id, err := s.service.Auth.SignUp(ctx, in)

	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{Id: id}, nil
}

func (s *Server) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.Token, error) {
	token, err := s.service.Auth.GenerateToken(ctx, in)

	if err != nil {
		return nil, err
	}

	return &pb.Token{Token: token}, nil
}
