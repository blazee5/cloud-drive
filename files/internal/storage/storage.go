package storage

import (
	"context"
	pb "github.com/blazee5/cloud-drive-protos/files"
	"github.com/blazee5/cloud-drive/files/ent"
	"github.com/blazee5/cloud-drive/files/internal/storage/aws"
	"github.com/blazee5/cloud-drive/files/internal/storage/postgres"
	"github.com/minio/minio-go/v7"
)

type Storage struct {
	PostgresStorage
	AwsStorage
}

type PostgresStorage interface {
	GetAllByID(ctx context.Context, userID string) ([]*pb.FileInfo, error)
	GetByID(ctx context.Context, ID int) (*ent.File, error)
	Create(ctx context.Context, userID, fileName, contentType string) (int, error)
	AddCount(ctx context.Context, ID int) error
	Update(ctx context.Context, ID int, input *pb.UpdateFileRequest) error
	Delete(ctx context.Context, ID int) error
}

type AwsStorage interface {
	SaveFile(ctx context.Context, bucket, fileName, contentType string, chunk []byte) error
	DownloadFile(ctx context.Context, bucket, fileName string) ([]byte, error)
	UpdateFile(ctx context.Context, bucket, oldName, newName string) error
	DeleteFile(ctx context.Context, bucket, fileName string) error
}

func NewStorage(db *ent.Client, awsClient *minio.Client) *Storage {
	return &Storage{
		PostgresStorage: postgres.NewFileStorage(db),
		AwsStorage:      aws.NewStorage(awsClient),
	}
}
