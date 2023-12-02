package service

import (
	"context"
	pb "github.com/blazee5/cloud-drive-protos/files"
	"github.com/blazee5/cloud-drive/api_gateway/internal/clients/file/grpc"
	"github.com/blazee5/cloud-drive/api_gateway/internal/domain"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"net/http"
)

type Service struct {
	log *zap.SugaredLogger
	api pb.FileServiceClient
}

func NewService(log *zap.SugaredLogger) *Service {
	return &Service{
		log: log,
		api: grpc.NewFileServiceClient(log),
	}
}

func (s *Service) GetFiles(ctx context.Context, userID string) ([]*pb.FileInfo, error) {
	files, err := s.api.GetFiles(ctx, &pb.UserRequest{UserId: userID})

	if err != nil {
		return files.GetFiles(), err
	}

	return files.GetFiles(), nil
}

func (s *Service) UploadFile(ctx context.Context, userID string, fileHeader *multipart.FileHeader) (int, error) {
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
	_, err := s.api.DeleteFile(ctx, &pb.FileRequest{
		Id:     int64(ID),
		UserId: userID,
	})

	if err != nil {
		return err
	}

	return nil
}
