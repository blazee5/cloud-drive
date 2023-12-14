package service

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/blazee5/cloud-drive-protos/auth"
	"github.com/blazee5/cloud-drive/auth/internal/config"
	"github.com/blazee5/cloud-drive/auth/internal/domain"
	"github.com/blazee5/cloud-drive/auth/internal/rabbitmq"
	"github.com/blazee5/cloud-drive/auth/internal/storage"
	"github.com/blazee5/cloud-drive/auth/lib/auth"
	"github.com/blazee5/cloud-drive/auth/lib/http_errors"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

type Service interface {
	SignUp(ctx context.Context, input *pb.SignUpRequest) (string, error)
	GenerateToken(ctx context.Context, input *pb.SignInRequest) (string, error)
	ValidateEmail(ctx context.Context, input *pb.ValidateAccountRequest) error
}

type AuthService struct {
	log      *zap.SugaredLogger
	storage  storage.Storage
	producer *rabbitmq.Producer
	cfg      *config.Config
}

func NewAuthService(log *zap.SugaredLogger, storage storage.Storage, producer *rabbitmq.Producer, cfg *config.Config) *AuthService {
	return &AuthService{log: log, storage: storage, producer: producer, cfg: cfg}
}

func (s *AuthService) SignUp(ctx context.Context, input *pb.SignUpRequest) (string, error) {
	input.Password = auth.GenerateHashPassword(input.Password)

	code, err := uuid.NewUUID()

	if err != nil {
		return "", err
	}

	id, err := s.storage.SignUp(ctx, input, code.String())

	if err != nil {
		return "", err
	}

	email := domain.Email{
		Type:    "activate",
		To:      input.Email,
		Message: fmt.Sprintf("%s/auth/activate?code=%s", s.cfg.HttpServer.GatewayURL, code.String()),
	}

	msg, err := json.Marshal(&email)

	if err != nil {
		return "", err
	}

	err = s.producer.PublishMessage(ctx, string(msg))

	if err != nil {
		return "", err
	}

	return id, nil
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

func (s *AuthService) ValidateEmail(ctx context.Context, input *pb.ValidateAccountRequest) error {
	code, err := s.storage.GetActivationCode(ctx, input.GetCode())

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
