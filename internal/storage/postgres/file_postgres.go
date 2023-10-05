package postgres

import (
	"context"
	"github.com/blazee5/cloud-drive/internal/models"
	"github.com/jmoiron/sqlx"
	"time"
)

type FileStorage struct {
	db *sqlx.DB
}

func NewFileStorage(db *sqlx.DB) *FileStorage {
	return &FileStorage{db: db}
}

func (s *FileStorage) Create(ctx context.Context, file models.File) (int, error) {
	var id int

	q := `INSERT INTO files (name, created_at) VALUES ($1, $2) RETURNING id`

	err := s.db.QueryRowxContext(ctx, q, file.Name, time.Now()).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *FileStorage) AddCount(ctx context.Context, fileName string) error {
	q := `UPDATE files SET download_count = download_count + 1 WHERE name = $1`

	_, err := s.db.QueryContext(ctx, q, fileName)

	if err != nil {
		return err
	}

	return nil
}
