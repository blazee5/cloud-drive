package app

import (
	"fmt"
	pb "github.com/blazee5/cloud-drive/microservices/files/api/v1"
	"github.com/blazee5/cloud-drive/microservices/files/internal/config"
	grpcServer "github.com/blazee5/cloud-drive/microservices/files/internal/grpc"
	"github.com/blazee5/cloud-drive/microservices/files/internal/service"
	"github.com/blazee5/cloud-drive/microservices/files/internal/storage"
	"github.com/blazee5/cloud-drive/microservices/files/lib/db/aws"
	"github.com/blazee5/cloud-drive/microservices/files/lib/db/postgres"
	"github.com/blazee5/cloud-drive/microservices/files/lib/logger"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	log := logger.NewLogger()

	awsClient := aws.NewAWSClient(cfg)

	db := postgres.New(cfg)

	storages := storage.NewStorage(db, awsClient)
	services := service.NewFileService(log, storages)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.HTTPServer.Port))
	if err != nil {
		log.Info("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFileServiceServer(s, grpcServer.NewServer(log, services))

	log.Info(fmt.Sprintf("server listening at %s", lis.Addr().String()))

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Infof("error while listen server: %s", err)
		}
	}()

	defer log.Sync()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	s.GracefulStop()
	if err = db.Close(); err != nil {
		log.Infof("error while close db conn: %v", err)
	}
}
