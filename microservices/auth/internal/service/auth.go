package service

import (
	"context"
	pb "github.com/blazee5/cloud-drive/microservices/auth/api/v1"
	"github.com/blazee5/cloud-drive/microservices/auth/internal/storage"
	"github.com/blazee5/cloud-drive/microservices/auth/lib/auth"
	"go.uber.org/zap"
)

type AuthService struct {
	log     *zap.SugaredLogger
	storage *storage.Storage
}

func NewAuthService(log *zap.SugaredLogger, storage *storage.Storage) *AuthService {
	return &AuthService{log: log, storage: storage}
}

func (s *AuthService) SignUp(ctx context.Context, input *pb.SignUpRequest) (string, error) {
	input.Password = auth.GenerateHashPassword(input.Password)
	return s.storage.Auth.SignUp(ctx, input)
}

func (s *AuthService) GenerateToken(ctx context.Context, input *pb.SignInRequest) (string, error) {
	input.Password = auth.GenerateHashPassword(input.Password)
	user, err := s.storage.Auth.VerifyUser(ctx, input)

	if err != nil {
		return "", err
	}

	token, err := auth.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
