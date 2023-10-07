package app

import (
	"fmt"
	"github.com/blazee5/cloud-drive/microservices/files/internal/config"
	"github.com/blazee5/cloud-drive/microservices/files/internal/handler"
	"github.com/blazee5/cloud-drive/microservices/files/internal/service"
	"github.com/blazee5/cloud-drive/microservices/files/internal/storage"
	"github.com/blazee5/cloud-drive/microservices/files/internal/storage/postgres"
	"github.com/blazee5/cloud-drive/microservices/files/lib/logger"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	log := logger.NewLogger()

	app := fiber.New()

	db := postgres.NewPostgres(cfg)
	storages := storage.NewStorage(db)
	services := service.NewService(log, storages)
	handlers := handler.NewHandler(log, services)

	handlers.InitRoutes(app)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		log.Info("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	log.Info(fmt.Sprintf("server listening at %s", lis.Addr().String()))

	go func() {
		if err := server.Serve(lis); err != nil {
			log.Infof("error while listen server: %s", err)
		}
	}()

	defer log.Sync()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	server.GracefulStop()
	if err := db.Close(); err != nil {
		log.Infof("error while close db: %s", err)
	}
}
