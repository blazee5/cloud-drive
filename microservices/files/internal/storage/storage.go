package storage

import (
	"context"
	"github.com/blazee5/cloud-drive/microservices/files/internal/models"
	"github.com/blazee5/cloud-drive/microservices/files/internal/storage/postgres"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	File
}

type File interface {
	Create(ctx context.Context, input models.File) (int, error)
	AddCount(ctx context.Context, fileName string) error
}

func NewStorage(db *sqlx.DB) *Storage {

	return &Storage{
		File: postgres.NewFileStorage(db),
	}
}
