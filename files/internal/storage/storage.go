package storage

import (
	"context"
	pb "github.com/blazee5/cloud-drive-protos/files"
	"github.com/blazee5/cloud-drive/files/internal/models"
	"github.com/blazee5/cloud-drive/files/internal/storage/aws"
	"github.com/blazee5/cloud-drive/files/internal/storage/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/minio/minio-go/v7"
)

type Storage struct {
	PostgresStorage
	AwsStorage
}

type PostgresStorage interface {
	GetAllByID(ctx context.Context, userID string, input *pb.GetFilesRequest) (models.FileList, error)
	GetByID(ctx context.Context, ID int) (models.File, error)
	Create(ctx context.Context, userID, fileName, contentType string) (int, error)
	AddCount(ctx context.Context, ID int) error
	Update(ctx context.Context, ID int, name string) error
	Delete(ctx context.Context, ID int) error
}

type AwsStorage interface {
	SaveFile(ctx context.Context, bucket, fileName, contentType string, chunk []byte) error
	DownloadFile(ctx context.Context, bucket, fileName string) ([]byte, error)
	UpdateFile(ctx context.Context, bucket, oldName, newName string) error
	DeleteFile(ctx context.Context, bucket, fileName string) error
}

func NewStorage(db *pgxpool.Pool, awsClient *minio.Client) *Storage {
	return &Storage{
		PostgresStorage: postgres.NewFileStorage(db),
		AwsStorage:      aws.NewStorage(awsClient),
	}
}
