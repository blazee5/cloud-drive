package service

import (
	"context"
	pb "github.com/blazee5/cloud-drive-protos/auth"
	"github.com/blazee5/cloud-drive/auth/internal/storage"
	"github.com/blazee5/cloud-drive/auth/lib/auth"
	"github.com/blazee5/cloud-drive/auth/lib/http_errors"
	"go.uber.org/zap"
	"time"
)

type Service interface {
	SignUp(ctx context.Context, input *pb.SignUpRequest) (string, error)
	GenerateToken(ctx context.Context, input *pb.SignInRequest) (string, error)
	ValidateEmail(ctx context.Context, input *pb.ValidateCodeRequest) error
}

type AuthService struct {
	log     *zap.SugaredLogger
	storage storage.Storage
}

func NewAuthService(log *zap.SugaredLogger, storage storage.Storage) *AuthService {
	return &AuthService{log: log, storage: storage}
}

func (s *AuthService) SignUp(ctx context.Context, input *pb.SignUpRequest) (string, error) {
	input.Password = auth.GenerateHashPassword(input.Password)
	return s.storage.SignUp(ctx, input)
}

func (s *AuthService) GenerateToken(ctx context.Context, input *pb.SignInRequest) (string, error) {
	input.Password = auth.GenerateHashPassword(input.Password)
	user, err := s.storage.VerifyUser(ctx, input)

	if err != nil {
		return "", err
	}

	token, err := auth.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) ValidateEmail(ctx context.Context, input *pb.ValidateCodeRequest) error {
	code, err := s.storage.GetActivationCode(ctx, input.GetUserId(), input.GetCode())

	if err != nil {
		return err
	}

	if code.ExpireDate.Before(time.Now()) {
		return http_errors.CodeExpired
	}

	err = s.storage.ActivateUser(ctx, code.UserID.Hex())

	if err != nil {
		return err
	}

	return s.storage.DeleteActivationCode(ctx, code.ID.Hex())
}
