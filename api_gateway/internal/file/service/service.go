package service

import (
	"context"
	pb "github.com/blazee5/cloud-drive-protos/files"
	"github.com/blazee5/cloud-drive/api_gateway/internal/clients/file/grpc"
	"github.com/blazee5/cloud-drive/api_gateway/internal/domain"
	"github.com/blazee5/cloud-drive/api_gateway/lib/tracer"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"net/http"
)

type Service struct {
	log    *zap.SugaredLogger
	api    pb.FileServiceClient
	tracer trace.Tracer
}

func NewService(log *zap.SugaredLogger, tracer *tracer.JaegerTracing) *Service {
	return &Service{
		log:    log,
		api:    grpc.NewFileServiceClient(log, tracer),
		tracer: tracer.Tracer,
	}
}

func (s *Service) GetFiles(ctx context.Context, userID, orderBy, orderDir string, page, size int) (domain.FileList, error) {
	ctx, span := s.tracer.Start(ctx, "filesService.GetFiles")
	defer span.End()

	filesProto, err := s.api.GetFiles(ctx, &pb.GetFilesRequest{
		UserId:   userID,
		Page:     int64(page),
		Size:     int64(size),
		OrderBy:  orderBy,
		OrderDir: orderDir,
	})

	if err != nil {
		return domain.FileList{}, err
	}

	files := make([]domain.FileInfo, 0, len(filesProto.Files))

	for _, file := range filesProto.Files {
		files = append(files, domain.FileInfo{
			ID:            int(file.Id),
			Name:          file.Name,
			UserID:        file.UserId,
			DownloadCount: int(file.DownloadCount),
			CreatedAt:     file.CreatedAt.AsTime(),
		})
	}

	return domain.FileList{
		Total:      int(filesProto.Total),
		TotalPages: int(filesProto.TotalPages),
		Page:       int(filesProto.Page),
		Size:       int(filesProto.Size),
		Files:      files,
	}, nil
}

func (s *Service) UploadFile(ctx context.Context, userID string, fileHeader *multipart.FileHeader) (int, error) {
	ctx, span := s.tracer.Start(ctx, "filesService.UploadFile")
	defer span.End()

	file, err := fileHeader.Open()
	defer file.Close()

	if err != nil {
		return 0, err
	}

	bytes, err := io.ReadAll(file)

	if err != nil {
		return 0, err
	}

	res, err := s.api.UploadFile(ctx, &pb.UploadRequest{
		UserId:   userID,
		FileName: fileHeader.Filename,
		FileType: http.DetectContentType(bytes),
		Chunk:    bytes,
	})

	if err != nil {
		return 0, err
	}

	return int(res.GetId()), nil
}

func (s *Service) DownloadFile(ctx context.Context, ID int, userID string) (*pb.File, error) {
	ctx, span := s.tracer.Start(ctx, "filesService.DownloadFile")
	defer span.End()

	file, err := s.api.DownloadFile(ctx, &pb.FileRequest{
		Id:     int64(ID),
		UserId: userID,
	})

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *Service) UpdateFile(ctx context.Context, ID int, userID string, input domain.UpdateFileInput) error {
	ctx, span := s.tracer.Start(ctx, "filesService.UpdateFile")
	defer span.End()

	_, err := s.api.UpdateFile(ctx, &pb.UpdateFileRequest{
		Id:     int64(ID),
		UserId: userID,
		Name:   input.FileName,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteFile(ctx context.Context, ID int, userID string) error {
	ctx, span := s.tracer.Start(ctx, "filesService.DeleteFile")
	defer span.End()

	_, err := s.api.DeleteFile(ctx, &pb.FileRequest{
		Id:     int64(ID),
		UserId: userID,
	})

	if err != nil {
		return err
	}

	return nil
}
