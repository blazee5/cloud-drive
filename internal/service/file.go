package service

import (
	"context"
	"github.com/blazee5/cloud-drive/internal/models"
	"github.com/blazee5/cloud-drive/internal/storage"
	"go.uber.org/zap"
)

type FileService struct {
	log *zap.SugaredLogger
	db  *storage.Storage
}

func NewFileService(log *zap.SugaredLogger, db *storage.Storage) *FileService {
	return &FileService{log: log, db: db}
}

func (s *FileService) Upload(ctx context.Context, file models.File) (int, error) {
	return s.db.File.Create(ctx, file)
}

func (s *FileService) AddCount(ctx context.Context, fileName string) error {
	return s.db.File.AddCount(ctx, fileName)
}
