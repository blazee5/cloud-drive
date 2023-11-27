package grpc

import (
	"context"
	"errors"
	pb "github.com/blazee5/cloud-drive/microservices/files/api/v1"
	"github.com/blazee5/cloud-drive/microservices/files/internal/service"
	"github.com/blazee5/cloud-drive/microservices/files/lib/http_errors"
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

func (s *Server) GetFiles(ctx context.Context, input *pb.UserRequest) (*pb.FileResponse, error) {
	if input.GetUserId() == "" {
		return &pb.FileResponse{
			Files: make([]*pb.FileInfo, 0),
		}, status.Errorf(codes.InvalidArgument, "user_id is required field")
	}

	files, err := s.service.GetFilesByID(ctx, input.GetUserId())

	if err != nil {
		s.log.Infof("error while get user files: %v", err)
		return &pb.FileResponse{
			Files: make([]*pb.FileInfo, 0),
		}, status.Errorf(codes.Internal, "server error")
	}

	return &pb.FileResponse{
		Files: files,
	}, nil
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
	file, err := s.service.Download(ctx, input.GetUserId(), int(input.GetId()))

	if err != nil {
		s.log.Infof("error while download file: %v", err)
		return &pb.File{}, status.Error(codes.Internal, "server error")
	}

	return &pb.File{
		Id:     int64(file.ID),
		Name:   file.Name,
		UserId: file.UserID,
		Chunk:  file.Chunk,
	}, nil
}

func (s *Server) UpdateFile(ctx context.Context, input *pb.UpdateFileRequest) (*pb.SuccessResponse, error) {
	err := s.service.Update(ctx, input.GetUserId(), int(input.GetId()), input)

	if errors.Is(err, http_errors.PermissionDenied) {
		return &pb.SuccessResponse{
			Success: "",
		}, status.Errorf(codes.PermissionDenied, "permission denied")
	}

	if err != nil {
		s.log.Infof("error while update file: %v", err)
		return &pb.SuccessResponse{
			Success: "",
		}, status.Errorf(codes.Internal, "server error")
	}

	return &pb.SuccessResponse{
		Success: "success",
	}, nil
}

func (s *Server) DeleteFile(ctx context.Context, input *pb.FileRequest) (*pb.SuccessResponse, error) {
	err := s.service.Delete(ctx, input.GetUserId(), int(input.GetId()))

	if errors.Is(err, http_errors.PermissionDenied) {
		return &pb.SuccessResponse{
			Success: "",
		}, status.Errorf(codes.PermissionDenied, "permission denied")
	}

	if err != nil {
		s.log.Infof("error while delete file: %v", err)
		return &pb.SuccessResponse{
			Success: "",
		}, status.Errorf(codes.Internal, "server error")
	}

	return &pb.SuccessResponse{
		Success: "success",
	}, nil
}
