package service

import (
	"context"
	pb "github.com/blazee5/cloud-drive/microservices/auth/api/v1"
	"github.com/blazee5/cloud-drive/microservices/auth/internal/storage"
	"go.uber.org/zap"
)

type Auth interface {
	SignUp(ctx context.Context, input *pb.SignUpRequest) (string, error)
	GenerateToken(ctx context.Context, input *pb.SignInRequest) (string, error)
}

type Service struct {
	Auth
}

func NewService(log *zap.SugaredLogger, storage *storage.Storage) *Service {
	return &Service{
		Auth: NewAuthService(log, storage),
	}
}
