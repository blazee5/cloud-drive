package service

import (
	"context"
	pb "github.com/blazee5/cloud-drive-protos/files"
	"github.com/blazee5/cloud-drive/files/internal/models"
	"github.com/blazee5/cloud-drive/files/internal/storage"
	"github.com/blazee5/cloud-drive/files/lib/http_errors"
	"go.uber.org/zap"
)

type Service interface {
	GetFilesByID(ctx context.Context, userID string) ([]*pb.FileInfo, error)
	Upload(ctx context.Context, fileName, userID string, chunk []byte, contentType string) (int, error)
	Download(ctx context.Context, userID string, ID int) (*models.File, error)
	Update(ctx context.Context, userID string, ID int, input *pb.UpdateFileRequest) error
	Delete(ctx context.Context, userID string, ID int) error
}

type FileService struct {
	log  *zap.SugaredLogger
	repo *storage.Storage
}

func NewFileService(log *zap.SugaredLogger, repo *storage.Storage) *FileService {
	return &FileService{log: log, repo: repo}
}

func (s *FileService) GetFilesByID(ctx context.Context, userID string) ([]*pb.FileInfo, error) {
	return s.repo.PostgresStorage.GetAllByID(ctx, userID)
}

func (s *FileService) Upload(ctx context.Context, fileName string, userID string, chunk []byte, contentType string) (int, error) {
	err := s.repo.AwsStorage.SaveFile(ctx, userID, fileName, contentType, chunk)

	if err != nil {
		return 0, err
	}

	return s.repo.PostgresStorage.Create(ctx, userID, fileName, contentType)
}

func (s *FileService) Download(ctx context.Context, userID string, ID int) (*models.File, error) {
	file, err := s.repo.GetByID(ctx, ID)

	if err != nil {
		return nil, err
	}

	object, err := s.repo.DownloadFile(ctx, userID, file.Name)

	if err != nil {
		return nil, err
	}

	err = s.repo.AddCount(ctx, ID)

	if err != nil {
		s.log.Infof("error while add download count: %v", err)
	}

	return &models.File{
		ID:          file.ID,
		Name:        file.Name,
		UserID:      file.UserID,
		ContentType: file.ContentType,
		Chunk:       object,
	}, nil
}

func (s *FileService) Update(ctx context.Context, userID string, ID int, input *pb.UpdateFileRequest) error {
	file, err := s.repo.GetByID(ctx, ID)

	if err != nil {
		return err
	}

	if file.UserID != userID {
		return http_errors.PermissionDenied
	}

	err = s.repo.AwsStorage.UpdateFile(ctx, userID, file.Name, input.GetName())

	if err != nil {
		return err
	}

	err = s.repo.PostgresStorage.Update(ctx, ID, input)

	if err != nil {
		return err
	}

	return nil
}

func (s *FileService) Delete(ctx context.Context, userID string, ID int) error {
	file, err := s.repo.GetByID(ctx, ID)

	if err != nil {
		return err
	}

	if file.UserID != userID {
		return http_errors.PermissionDenied
	}

	err = s.repo.AwsStorage.DeleteFile(ctx, userID, file.Name)

	if err != nil {
		return err
	}

	err = s.repo.PostgresStorage.Delete(ctx, ID)

	if err != nil {
		return err
	}

	return nil
}
