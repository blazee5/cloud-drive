package storage

import (
	"context"
	pb "github.com/blazee5/cloud-drive-protos/auth"
	"github.com/blazee5/cloud-drive/auth/internal/models"
)

type Storage interface {
	SignUp(ctx context.Context, input *pb.SignUpRequest) (string, error)
	VerifyUser(ctx context.Context, input *pb.SignInRequest) (models.User, error)
	GetActivationCode(ctx context.Context, userID, code string) (models.ActivationCode, error)
	ActivateUser(ctx context.Context, userID string) error
	DeleteActivationCode(ctx context.Context, ID string) error
}
