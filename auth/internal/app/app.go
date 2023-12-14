package app

import (
	"context"
	"fmt"
	pb "github.com/blazee5/cloud-drive-protos/auth"
	"github.com/blazee5/cloud-drive/auth/internal/config"
	"github.com/blazee5/cloud-drive/auth/internal/handler"
	producer "github.com/blazee5/cloud-drive/auth/internal/rabbitmq"
	"github.com/blazee5/cloud-drive/auth/internal/service"
	"github.com/blazee5/cloud-drive/auth/internal/storage/mongodb"
	"github.com/blazee5/cloud-drive/auth/lib/logger"
	"github.com/blazee5/cloud-drive/auth/lib/rabbitmq"
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
	db := client.Database(cfg.DBName)
	storages := mongodb.NewAuthStorage(db)
	rabbitConn := rabbitmq.NewRabbitMQConn(cfg)
	msgProducer := producer.NewProducer(log, rabbitConn)
	services := service.NewAuthService(log, storages, msgProducer, cfg)

	err := msgProducer.InitProducer(cfg)

	if err != nil {
		log.Infof("error while init producer: %v", err)
		cancel()
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.HttpServer.Port))
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

	if err = client.Disconnect(ctx); err != nil {
		log.Infof("error while close db: %s", err)
	}
}
