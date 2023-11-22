package postgres

import (
	"context"
	"github.com/blazee5/cloud-drive/microservices/files/ent"
	"github.com/blazee5/cloud-drive/microservices/files/ent/file"
)

type FileStorage struct {
	db *ent.Client
}

func NewFileStorage(db *ent.Client) *FileStorage {
	return &FileStorage{db: db}
}

func (s *FileStorage) Create(ctx context.Context, fileName string, userId string) (int, error) {
	res, err := s.db.File.Create().SetName(fileName).SetUserID(userId).SetDownloadCount(0).Save(ctx)

	if err != nil {
		return 0, err
	}

	return res.ID, nil
}

func (s *FileStorage) AddCount(ctx context.Context, fileName string) error {
	_, err := s.db.File.Update().Where(file.Name(fileName)).AddDownloadCount(1).Save(ctx)

	if err != nil {
		return err
	}

	return nil
}
