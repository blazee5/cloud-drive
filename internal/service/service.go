package service

import (
	"context"
	"github.com/blazee5/cloud-drive/internal/models"
	"github.com/blazee5/cloud-drive/internal/storage"
	"go.uber.org/zap"
)

type File interface {
	Upload(ctx context.Context, file models.File) (int, error)
	AddCount(ctx context.Context, fileName string) error
}

type Service struct {
	File
}

func NewService(log *zap.SugaredLogger, db *storage.Storage) *Service {
	return &Service{
		File: NewFileService(log, db),
	}
}
