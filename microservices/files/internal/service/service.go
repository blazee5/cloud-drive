package service

import (
	"context"
	"github.com/blazee5/cloud-drive/microservices/files/internal/models"
	"github.com/blazee5/cloud-drive/microservices/files/internal/storage"
	"go.uber.org/zap"
)

type Service interface {
	Upload(ctx context.Context, fileName string, userId string, chunk []byte, contentType string) (int, error)
	Download(ctx context.Context, userId, fileName string) (*models.File, error)
}

type FileService struct {
	log  *zap.SugaredLogger
	repo *storage.Storage
}

func NewFileService(log *zap.SugaredLogger, repo *storage.Storage) *FileService {
	return &FileService{log: log, repo: repo}
}

func (s *FileService) Upload(ctx context.Context, fileName string, userId string, chunk []byte, contentType string) (int, error) {
	err := s.repo.AwsStorage.SaveFile(ctx, userId, fileName, contentType, chunk)

	if err != nil {
		return 0, err
	}

	return s.repo.PostgresStorage.Create(ctx, fileName, userId)
}

func (s *FileService) Download(ctx context.Context, userId, fileName string) (*models.File, error) {
	object, err := s.repo.DownloadFile(ctx, userId, fileName)

	if err != nil {
		return nil, err
	}

	file := models.File{}

	file.Chunk = object

	err = s.repo.AddCount(ctx, fileName)

	if err != nil {
		s.log.Infof("error while add download count: %v", err)
	}

	return &file, nil
}
