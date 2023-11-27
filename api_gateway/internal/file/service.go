package file

import (
	"context"
	"github.com/blazee5/cloud-drive/microservices/api_gateway/internal/domain"
	pb "github.com/blazee5/cloud-drive/microservices/api_gateway/proto/files"
	"mime/multipart"
)

type Service interface {
	GetFiles(ctx context.Context, userID string) ([]*pb.FileInfo, error)
	UploadFile(ctx context.Context, userID string, file *multipart.FileHeader) (int, error)
	DownloadFile(ctx context.Context, ID int, userID string) (*pb.File, error)
	UpdateFile(ctx context.Context, ID int, userID string, input domain.UpdateFileInput) error
	DeleteFile(ctx context.Context, ID int, userID string) error
}
