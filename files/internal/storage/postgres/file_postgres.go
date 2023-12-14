package postgres

import (
	"context"
	pb "github.com/blazee5/cloud-drive-protos/files"
	"github.com/blazee5/cloud-drive/files/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math"
)

type FileStorage struct {
	db *pgxpool.Pool
}

func NewFileStorage(db *pgxpool.Pool) *FileStorage {
	return &FileStorage{db: db}
}

func (s *FileStorage) GetAllByID(ctx context.Context, userID string, input *pb.GetFilesRequest) (models.FileList, error) {
	var total int

	if err := s.db.QueryRow(ctx, "SELECT COUNT(*) FROM files WHERE user_id = $1", userID).Scan(&total); err != nil {
		return models.FileList{}, err
	}

	if total == 0 {
		return models.FileList{
			Files: make([]*pb.FileInfo, 0),
		}, nil
	}

	var offset int

	if input.GetPage() == 0 {
		offset = 0
	}
	offset = int((input.GetPage() - 1) * input.GetSize())

	rows, err := s.db.Query(ctx, "SELECT id, name, user_id, download_count, created_at FROM files WHERE user_id = $1 ORDER BY id LIMIT $2 OFFSET $3", userID, input.GetSize(), offset)

	if err != nil {
		return models.FileList{}, err
	}

	files := make([]models.FileInfo, 0, input.Size)

	for rows.Next() {
		var file models.FileInfo

		if err = rows.Scan(&file.ID, &file.Name, &file.UserID, &file.DownloadCount, &file.CreatedAt); err != nil {
			return models.FileList{}, err
		}

		files = append(files, file)
	}

	filesList := make([]*pb.FileInfo, 0, len(files))
	for _, file := range files {
		filesList = append(filesList, &pb.FileInfo{
			Id:            int64(file.ID),
			Name:          file.Name,
			UserId:        file.UserID,
			DownloadCount: int64(file.DownloadCount),
			CreatedAt:     timestamppb.New(file.CreatedAt),
		})
	}

	return models.FileList{
		Total:      total,
		TotalPages: int(math.Ceil(float64(total) / float64(input.Size))),
		Page:       int(input.Page),
		Size:       int(input.Size),
		Files:      filesList,
	}, nil
}

func (s *FileStorage) GetByID(ctx context.Context, ID int) (models.File, error) {
	var file models.File

	err := s.db.QueryRow(ctx, "SELECT id, name, user_id, content_type, download_count, created_at WHERE id = $1", ID).
		Scan(&file.ID, &file.Name, &file.UserID, &file.ContentType, &file.DownloadCount, &file.CreatedAt)

	if err != nil {
		return models.File{}, err
	}

	return file, nil
}

func (s *FileStorage) Create(ctx context.Context, userID string, fileName, contentType string) (int, error) {
	var id int

	err := s.db.QueryRow(ctx, `INSERT INTO files (name, user_id, content_type)
		VALUES ($1, $2, $3) RETURNING id`, fileName, userID, contentType).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *FileStorage) AddCount(ctx context.Context, ID int) error {
	_, err := s.db.Exec(ctx, "UPDATE files SET download_count = download_count + 1 WHERE id = $1", ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *FileStorage) Update(ctx context.Context, ID int, input *pb.UpdateFileRequest) error {
	var file models.FileInfo

	err := s.db.QueryRow(ctx, "UPDATE files SET name = $1 WHERE id = $1 RETURNING id, name, user_id, download_count, created_at", ID, input.Name).
		Scan(&file.ID, &file.Name, &file.UserID, &file.DownloadCount, &file.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (s *FileStorage) Delete(ctx context.Context, ID int) error {
	_, err := s.db.Exec(ctx, "DELETE FROM files WHERE id = $1", ID)

	if err != nil {
		return err
	}

	return nil
}
