package postgres

import (
	"context"
	pb "github.com/blazee5/cloud-drive-protos/files"
	"github.com/blazee5/cloud-drive/files/ent"
	"github.com/blazee5/cloud-drive/files/ent/file"
)

type FileStorage struct {
	db *ent.Client
}

func NewFileStorage(db *ent.Client) *FileStorage {
	return &FileStorage{db: db}
}

func (s *FileStorage) GetAllByID(ctx context.Context, userID string) ([]*pb.FileInfo, error) {
	files, err := s.db.File.Query().Where(file.UserID(userID)).All(ctx)

	if err != nil {
		return nil, err
	}

	customFiles := make([]*pb.FileInfo, 0)

	for _, file := range files {
		fileInfo := &pb.FileInfo{
			Id:            int64(file.ID),
			Name:          file.Name,
			UserId:        file.UserID,
			DownloadCount: int64(file.DownloadCount),
		}

		customFiles = append(customFiles, fileInfo)
	}

	return customFiles, nil
}

func (s *FileStorage) GetByID(ctx context.Context, ID int) (*ent.File, error) {
	file, err := s.db.File.Get(ctx, ID)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *FileStorage) Create(ctx context.Context, userID string, fileName, contentType string) (int, error) {
	res, err := s.db.File.Create().SetName(fileName).SetUserID(userID).SetContentType(contentType).SetDownloadCount(0).Save(ctx)

	if err != nil {
		return 0, err
	}

	return res.ID, nil
}

func (s *FileStorage) AddCount(ctx context.Context, ID int) error {
	_, err := s.db.File.UpdateOneID(ID).AddDownloadCount(1).Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (s *FileStorage) Update(ctx context.Context, ID int, input *pb.UpdateFileRequest) error {
	_, err := s.db.File.UpdateOneID(ID).SetName(input.GetName()).Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (s *FileStorage) Delete(ctx context.Context, ID int) error {
	err := s.db.File.DeleteOneID(ID).Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
