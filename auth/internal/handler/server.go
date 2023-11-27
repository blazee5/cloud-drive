package handler

import (
	"context"
	"github.com/blazee5/cloud-drive-protos/auth"
	"github.com/blazee5/cloud-drive/microservices/auth/internal/service"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	log     *zap.SugaredLogger
	service *service.AuthService
	auth.UnimplementedAuthServiceServer
}

func NewServer(log *zap.SugaredLogger, service *service.AuthService) *Server {
	return &Server{log: log, service: service}
}

func (s *Server) SignUp(ctx context.Context, in *auth.SignUpRequest) (*auth.UserResponse, error) {
	id, err := s.service.SignUp(ctx, in)

	if err != nil {
		s.log.Infof("error while signup: %v", err)
		return &auth.UserResponse{}, status.Errorf(codes.Internal, "server error")
	}

	return &auth.UserResponse{Id: id}, nil
}

func (s *Server) SignIn(ctx context.Context, in *auth.SignInRequest) (*auth.Token, error) {
	token, err := s.service.GenerateToken(ctx, in)

	if err != nil {
		return nil, err
	}

	return &auth.Token{Token: token}, nil
}
