package handler

import (
	"context"
	"errors"
	"github.com/blazee5/cloud-drive-protos/auth"
	"github.com/blazee5/cloud-drive/auth/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	log     *zap.SugaredLogger
	service service.Service
	auth.UnimplementedAuthServiceServer
}

func NewServer(log *zap.SugaredLogger, service service.Service) *Server {
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

	if errors.Is(err, mongo.ErrNoDocuments) {
		return &auth.Token{}, status.Errorf(codes.NotFound, "invalid credentials")
	}

	if err != nil {
		s.log.Infof("error while signin: %v", err)
		return &auth.Token{}, status.Errorf(codes.Internal, "server error")
	}

	return &auth.Token{Token: token}, nil
}

func (s *Server) ValidateAccount(ctx context.Context, in *auth.ValidateAccountRequest) (*auth.ValidateAccountResponse, error) {
	err := s.service.ValidateEmail(ctx, in)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return &auth.ValidateAccountResponse{}, status.Errorf(codes.NotFound, "invalid code")
	}

	if err != nil {
		s.log.Infof("error while validate code: %v", err)
		return &auth.ValidateAccountResponse{}, status.Errorf(codes.Internal, "server error")
	}

	return &auth.ValidateAccountResponse{
		Status: "success",
	}, nil
}
