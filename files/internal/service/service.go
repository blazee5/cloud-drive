package service

import (
	"context"
	pb "github.com/blazee5/cloud-drive-protos/files"
	"github.com/blazee5/cloud-drive/files/internal/domain"
	"github.com/blazee5/cloud-drive/files/internal/models"
	"github.com/blazee5/cloud-drive/files/internal/storage"
	"github.com/blazee5/cloud-drive/files/lib/http_errors"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"path/filepath"
)

type Service interface {
	GetFilesByID(ctx context.Context, userID string, input *pb.GetFilesRequest) (models.FileList, error)
	Upload(ctx context.Context, fileName, userID, contentType string, chunk []byte) (int, error)
	Download(ctx context.Context, userID string, ID int) (*domain.File, error)
	Update(ctx context.Context, userID string, ID int, input *pb.UpdateFileRequest) error
	Delete(ctx context.Context, userID string, ID int) error
}

type FileService struct {
	log    *zap.SugaredLogger
	repo   *storage.Storage
	tracer trace.Tracer
}

func NewFileService(log *zap.SugaredLogger, repo *storage.Storage, tracer trace.Tracer) *FileService {
	return &FileService{log: log, repo: repo, tracer: tracer}
}

func (s *FileService) GetFilesByID(ctx context.Context, userID string, input *pb.GetFilesRequest) (models.FileList, error) {
	ctx, span := s.tracer.Start(ctx, "fileService.GetFilesByID")
	defer span.End()

	return s.repo.PostgresStorage.GetAllByID(ctx, userID, input)
}

func (s *FileService) Upload(ctx context.Context, fileName, userID, contentType string, chunk []byte) (int, error) {
	ctx, span := s.tracer.Start(ctx, "fileService.Upload")
	defer span.End()

	err := s.repo.AwsStorage.SaveFile(ctx, userID, fileName, contentType, chunk)

	if err != nil {
		return 0, err
	}

	return s.repo.PostgresStorage.Create(ctx, userID, fileName, contentType)
}

func (s *FileService) Download(ctx context.Context, userID string, ID int) (*domain.File, error) {
	ctx, span := s.tracer.Start(ctx, "fileService.Download")
	defer span.End()

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

	return &domain.File{
		ID:          file.ID,
		Name:        file.Name,
		UserID:      file.UserID,
		ContentType: file.ContentType,
		Chunk:       object,
	}, nil
}

func (s *FileService) Update(ctx context.Context, userID string, ID int, input *pb.UpdateFileRequest) error {
	ctx, span := s.tracer.Start(ctx, "fileService.Update")
	defer span.End()

	file, err := s.repo.GetByID(ctx, ID)

	if err != nil {
		return err
	}

	if file.UserID != userID {
		return http_errors.PermissionDenied
	}

	newFileName := input.GetName() + filepath.Ext(file.Name)

	err = s.repo.AwsStorage.UpdateFile(ctx, userID, file.Name, newFileName)

	if err != nil {
		return err
	}

	err = s.repo.PostgresStorage.Update(ctx, ID, newFileName)

	if err != nil {
		return err
	}

	return nil
}

func (s *FileService) Delete(ctx context.Context, userID string, ID int) error {
	ctx, span := s.tracer.Start(ctx, "fileService.Delete")
	defer span.End()

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
