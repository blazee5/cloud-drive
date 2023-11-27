package handler

import (
	"context"
	pb "github.com/blazee5/cloud-drive/microservices/auth/api/v1"
	"github.com/blazee5/cloud-drive/microservices/auth/internal/service"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	log     *zap.SugaredLogger
	service *service.AuthService
	pb.UnimplementedAuthServiceServer
}

func NewServer(log *zap.SugaredLogger, service *service.AuthService) *Server {
	return &Server{log: log, service: service}
}

func (s *Server) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.UserResponse, error) {
	id, err := s.service.SignUp(ctx, in)

	if err != nil {
		s.log.Infof("error while signup: %v", err)
		return &pb.UserResponse{}, status.Errorf(codes.Internal, "server error")
	}

	return &pb.UserResponse{Id: id}, nil
}

func (s *Server) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.Token, error) {
	token, err := s.service.GenerateToken(ctx, in)

	if err != nil {
		return nil, err
	}

	return &pb.Token{Token: token}, nil
}
