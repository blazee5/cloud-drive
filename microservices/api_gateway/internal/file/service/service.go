package service

import (
	"context"
	"github.com/blazee5/cloud-drive/microservices/api_gateway/internal/clients/file/grpc"
	"github.com/blazee5/cloud-drive/microservices/api_gateway/internal/domain"
	pb "github.com/blazee5/cloud-drive/microservices/api_gateway/proto/files"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"path"
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
		FileType: path.Ext(fileHeader.Filename),
		Chunk:    bytes,
	})

	if err != nil {
		return 0, err
	}

	return int(res.GetId()), nil
}

func (s *Service) UpdateFile(ctx context.Context, userID string, input domain.UpdateFileInput) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) DeleteFile(ctx context.Context, ID int, userID string) error {
	//TODO implement me
	panic("implement me")
}
