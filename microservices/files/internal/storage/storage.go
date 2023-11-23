package storage

import (
	"context"
	"github.com/blazee5/cloud-drive/microservices/files/ent"
	"github.com/blazee5/cloud-drive/microservices/files/internal/storage/aws"
	"github.com/blazee5/cloud-drive/microservices/files/internal/storage/postgres"
	"github.com/minio/minio-go/v7"
)

type Storage struct {
	PostgresStorage
	AwsStorage
}

type PostgresStorage interface {
	Create(ctx context.Context, fileName, userId string) (int, error)
	GetById(ctx context.Context, fileName string) (*ent.File, error)
	AddCount(ctx context.Context, fileName string) error
}

type AwsStorage interface {
	SaveFile(ctx context.Context, bucket, fileName, contentType string, chunk []byte) error
	DownloadFile(ctx context.Context, bucket, fileName string) ([]byte, error)
}

func NewStorage(db *ent.Client, awsClient *minio.Client) *Storage {
	return &Storage{
		PostgresStorage: postgres.NewFileStorage(db),
		AwsStorage:      aws.NewStorage(awsClient),
	}
}
