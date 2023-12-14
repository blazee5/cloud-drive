package file

import (
	"context"
	pb "github.com/blazee5/cloud-drive-protos/files"
	"github.com/blazee5/cloud-drive/api_gateway/internal/domain"
	"mime/multipart"
)

type Service interface {
	GetFiles(ctx context.Context, userID string, page, size int) (*pb.GetFileResponse, error)
	UploadFile(ctx context.Context, userID string, file *multipart.FileHeader) (int, error)
	DownloadFile(ctx context.Context, ID int, userID string) (*pb.File, error)
	UpdateFile(ctx context.Context, ID int, userID string, input domain.UpdateFileInput) error
	DeleteFile(ctx context.Context, ID int, userID string) error
}
