package service

import (
	"context"
	"github.com/blazee5/cloud-drive/microservices/api_gateway/internal/clients/auth/grpc"
	"github.com/blazee5/cloud-drive/microservices/api_gateway/internal/domain"
	pb "github.com/blazee5/cloud-drive/microservices/api_gateway/proto/auth"
	"go.uber.org/zap"
)

type Service struct {
	log *zap.SugaredLogger
	api pb.AuthServiceClient
}

func NewService(log *zap.SugaredLogger) *Service {
	return &Service{
		log: log,
		api: grpc.NewAuthServiceClient(log),
	}
}

func (s *Service) SignUp(ctx context.Context, input domain.SignUpRequest) (string, error) {
	res, err := s.api.SignUp(ctx, &pb.SignUpRequest{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	})

	if err != nil {
		return "", err
	}

	return res.GetId(), nil
}

func (s *Service) SignIn(ctx context.Context, input domain.SignInRequest) (string, error) {
	res, err := s.api.SignIn(ctx, &pb.SignInRequest{
		Email:    input.Email,
		Password: input.Password,
	})

	if err != nil {
		return "", err
	}

	return res.GetToken(), nil
}

func (s *Service) ValidateUser(ctx context.Context, token string) (string, error) {
	res, err := s.api.ValidateUser(ctx, &pb.TokenRequest{
		Token: token,
	})

	if err != nil {
		return "", err
	}

	return res.GetId(), nil
}
