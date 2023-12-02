package storage

import (
	"context"
	pb "github.com/blazee5/cloud-drive-protos/auth"
	"github.com/blazee5/cloud-drive/auth/internal/auth"
	"github.com/blazee5/cloud-drive/auth/internal/storage/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type Auth interface {
	SignUp(ctx context.Context, input *pb.SignUpRequest) (string, error)
	VerifyUser(ctx context.Context, input *pb.SignInRequest) (auth.User, error)
}

type Storage struct {
	Auth
}

func NewStorage(db *mongo.Database) *Storage {
	return &Storage{
		Auth: mongodb.NewAuthStorage(db),
	}
}
