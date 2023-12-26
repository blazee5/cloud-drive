package grpc

import (
	"context"
	"database/sql"
	"errors"
	pb "github.com/blazee5/cloud-drive-protos/files"
	"github.com/blazee5/cloud-drive/files/internal/service"
	"github.com/blazee5/cloud-drive/files/lib/http_errors"
	"github.com/jackc/pgx/v5"
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

func (s *Server) GetFiles(ctx context.Context, input *pb.GetFilesRequest) (*pb.GetFileResponse, error) {
	if input.GetUserId() == "" {
		return &pb.GetFileResponse{}, status.Errorf(codes.InvalidArgument, "user_id is required field")
	}

	if input.GetSize() == 0 {
		input.Size = 5
	}

	if input.GetPage() == 0 {
		input.Page = 1
	}

	if input.GetOrderBy() == "" {
		input.OrderBy = "title"
	}

	if input.GetOrderDir() == "" {
		input.OrderDir = "ASC"
	}

	files, err := s.service.GetFilesByID(ctx, input.GetUserId(), input)

	if err != nil {
		s.log.Infof("error while get user files: %v", err)
		return &pb.GetFileResponse{}, status.Errorf(codes.Internal, "server error")
	}

	return &pb.GetFileResponse{
		Total:      int64(files.Total),
		TotalPages: int64(files.TotalPages),
		Page:       int64(files.Page),
		Size:       int64(files.Size),
		Files:      files.Files,
	}, nil
}

func (s *Server) UploadFile(ctx context.Context, input *pb.UploadRequest) (*pb.UploadResponse, error) {
	if input.GetFileName() == "" {
		return &pb.UploadResponse{}, status.Errorf(codes.InvalidArgument, "filename is required field")
	}

	if input.GetUserId() == "" {
		return &pb.UploadResponse{}, status.Errorf(codes.InvalidArgument, "user_id is required field")
	}

	if len(input.GetChunk()) == 0 {
		return &pb.UploadResponse{}, status.Errorf(codes.InvalidArgument, "chunk is required field")
	}

	id, err := s.service.Upload(ctx, input.GetFileName(), input.GetUserId(), input.GetFileType(), input.GetChunk())

	if err != nil {
		s.log.Infof("error while upload file: %v", err)
		return &pb.UploadResponse{}, status.Error(codes.Internal, "server error")
	}

	return &pb.UploadResponse{Id: int64(id)}, nil
}

func (s *Server) DownloadFile(ctx context.Context, input *pb.FileRequest) (*pb.File, error) {
	file, err := s.service.Download(ctx, input.GetUserId(), int(input.GetId()))

	if errors.Is(err, sql.ErrNoRows) {
		return &pb.File{}, status.Error(codes.NotFound, "file not found")
	}

	if err != nil {
		s.log.Infof("error while download file: %v", err)
		return &pb.File{}, status.Error(codes.Internal, "server error")
	}

	return &pb.File{
		Id:          int64(file.ID),
		Name:        file.Name,
		UserId:      file.UserID,
		ContentType: file.ContentType,
		Chunk:       file.Chunk,
	}, nil
}

func (s *Server) UpdateFile(ctx context.Context, input *pb.UpdateFileRequest) (*pb.SuccessResponse, error) {
	err := s.service.Update(ctx, input.GetUserId(), int(input.GetId()), input)

	if errors.Is(err, http_errors.PermissionDenied) {
		return &pb.SuccessResponse{}, status.Errorf(codes.PermissionDenied, "permission denied")
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return &pb.SuccessResponse{}, status.Error(codes.NotFound, "file not found")
	}

	if err != nil {
		s.log.Infof("error while update file: %v", err)
		return &pb.SuccessResponse{}, status.Errorf(codes.Internal, "server error")
	}

	return &pb.SuccessResponse{
		Success: "success",
	}, nil
}

func (s *Server) DeleteFile(ctx context.Context, input *pb.FileRequest) (*pb.SuccessResponse, error) {
	err := s.service.Delete(ctx, input.GetUserId(), int(input.GetId()))

	if errors.Is(err, http_errors.PermissionDenied) {
		return &pb.SuccessResponse{}, status.Errorf(codes.PermissionDenied, "permission denied")
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return &pb.SuccessResponse{}, status.Errorf(codes.NotFound, "file not found")
	}

	if err != nil {
		s.log.Infof("error while delete file: %v", err)
		return &pb.SuccessResponse{}, status.Errorf(codes.Internal, "server error")
	}

	return &pb.SuccessResponse{
		Success: "success",
	}, nil
}
