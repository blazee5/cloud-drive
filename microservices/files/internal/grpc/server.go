package grpc

import (
	"context"
	pb "github.com/blazee5/cloud-drive/microservices/files/api/v1"
	"github.com/blazee5/cloud-drive/microservices/files/internal/service"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	log     *zap.SugaredLogger
	service service.Service
	pb.UnimplementedFileServiceServer
}

func NewServer(log *zap.SugaredLogger, service service.Service) *Server {
	return &Server{log: log, service: service}
}

func (s *Server) UploadFile(ctx context.Context, input *pb.UploadRequest) (*pb.UploadResponse, error) {
	if input.GetFileName() == "" {
		return &pb.UploadResponse{
			Id: 0,
		}, status.Errorf(codes.InvalidArgument, "filename is required field")
	}

	if input.GetUserId() == "" {
		return &pb.UploadResponse{
			Id: 0,
		}, status.Errorf(codes.InvalidArgument, "user_id is required field")
	}

	if len(input.GetChunk()) == 0 {
		return &pb.UploadResponse{
			Id: 0,
		}, status.Errorf(codes.InvalidArgument, "chunk is required field")
	}

	id, err := s.service.Upload(ctx, input.GetFileName(), input.GetUserId(), input.GetChunk(), input.GetFileType())

	if err != nil {
		s.log.Infof("error while upload file: %v", err)
		return &pb.UploadResponse{Id: 0}, status.Error(codes.Internal, "server error")
	}

	return &pb.UploadResponse{Id: int64(id)}, nil
}

func (s *Server) DownloadFile(ctx context.Context, input *pb.FileRequest) (*pb.File, error) {
	file, err := s.service.Download(ctx, input.GetUserId(), input.GetFileName())

	if err != nil {
		s.log.Infof("error while download file: %v", err)
		return &pb.File{}, status.Error(codes.Internal, "server error")
	}

	return &pb.File{
		Id:     int64(file.Id),
		Name:   file.Name,
		UserId: file.UserId,
		Chunk:  file.Chunk,
	}, nil
}
