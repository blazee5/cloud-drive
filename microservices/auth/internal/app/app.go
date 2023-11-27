package app

import (
	"context"
	"fmt"
	pb "github.com/blazee5/cloud-drive/microservices/auth/api/v1"
	"github.com/blazee5/cloud-drive/microservices/auth/internal/config"
	"github.com/blazee5/cloud-drive/microservices/auth/internal/handler"
	"github.com/blazee5/cloud-drive/microservices/auth/internal/service"
	"github.com/blazee5/cloud-drive/microservices/auth/internal/storage"
	"github.com/blazee5/cloud-drive/microservices/auth/internal/storage/mongodb"
	"github.com/blazee5/cloud-drive/microservices/auth/lib/logger"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	log := logger.NewLogger()

	ctx, cancel := context.WithCancel(context.Background())

	client := mongodb.NewMongoDB(ctx, cfg)
	db := client.Database("cloud-drive")
	storages := storage.NewStorage(db)
	services := service.NewAuthService(log, storages)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		log.Info("failed to listen: %v", err)
	}

	log.Info(fmt.Sprintf("server listening at %s", lis.Addr().String()))
	s := grpc.NewServer()

	pb.RegisterAuthServiceServer(s, handler.NewServer(log, services))

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Infof("error while listen server: %s", err)
		}
	}()

	defer log.Sync()
	defer cancel()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	s.GracefulStop()

	if err := client.Disconnect(ctx); err != nil {
		log.Infof("error while close db: %s", err)
	}
}
