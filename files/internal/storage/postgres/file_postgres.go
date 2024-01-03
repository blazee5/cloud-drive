package postgres

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	pb "github.com/blazee5/cloud-drive-protos/files"
	"github.com/blazee5/cloud-drive/files/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math"
)

type FileStorage struct {
	db     *pgxpool.Pool
	tracer trace.Tracer
}

func NewFileStorage(db *pgxpool.Pool, tracer trace.Tracer) *FileStorage {
	return &FileStorage{db: db, tracer: tracer}
}

func (s *FileStorage) GetAllByID(ctx context.Context, userID string, input *pb.GetFilesRequest) (models.FileList, error) {
	ctx, span := s.tracer.Start(ctx, "fileStorage.GetAllByID")
	defer span.End()

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

	if input.GetOrderDir() == "desc" {
		input.OrderDir = "DESC"
	} else {
		input.OrderDir = "ASC"
	}

	offset = int((input.GetPage() - 1) * input.GetSize())

	sql, args, err := sq.
		Select("id", "name", "user_id", "download_count", "created_at").
		From("files").
		OrderBy(input.GetOrderBy() + " " + input.GetOrderDir()).
		Limit(uint64(input.GetSize())).
		Where(sq.Eq{"user_id": userID}).
		Offset(uint64(offset)).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	rows, err := s.db.Query(ctx, sql, args...)

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
	ctx, span := s.tracer.Start(ctx, "fileStorage.GetByID")
	defer span.End()

	var file models.File

	err := s.db.QueryRow(ctx, "SELECT id, name, user_id, content_type, download_count, created_at FROM files WHERE id = $1", ID).
		Scan(&file.ID, &file.Name, &file.UserID, &file.ContentType, &file.DownloadCount, &file.CreatedAt)

	if err != nil {
		return models.File{}, err
	}

	return file, nil
}

func (s *FileStorage) Create(ctx context.Context, userID string, fileName, contentType string) (int, error) {
	ctx, span := s.tracer.Start(ctx, "fileStorage.Create")
	defer span.End()

	var id int

	err := s.db.QueryRow(ctx, `INSERT INTO files (name, user_id, content_type)
		VALUES ($1, $2, $3) RETURNING id`, fileName, userID, contentType).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *FileStorage) AddCount(ctx context.Context, ID int) error {
	ctx, span := s.tracer.Start(ctx, "fileStorage.AddCount")
	defer span.End()

	_, err := s.db.Exec(ctx, "UPDATE files SET download_count = download_count + 1 WHERE id = $1", ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *FileStorage) Update(ctx context.Context, ID int, name string) error {
	ctx, span := s.tracer.Start(ctx, "fileStorage.Update")
	defer span.End()

	res, err := s.db.Exec(ctx, "UPDATE files SET name = $1 WHERE id = $2 RETURNING id, name, user_id, download_count, created_at", name, ID)

	if err != nil {
		return err
	}

	if res.RowsAffected() != 1 {
		return pgx.ErrNoRows
	}

	return nil
}

func (s *FileStorage) Delete(ctx context.Context, ID int) error {
	ctx, span := s.tracer.Start(ctx, "fileStorage.Delete")
	defer span.End()

	res, err := s.db.Exec(ctx, "DELETE FROM files WHERE id = $1", ID)

	if err != nil {
		return err
	}

	if res.RowsAffected() != 1 {
		return pgx.ErrNoRows
	}

	return nil
}
