package main

import (
	"fmt"
	config "github.com/blazee5/cloud-drive/internal/config"
	"github.com/blazee5/cloud-drive/internal/handler"
	"github.com/blazee5/cloud-drive/internal/service"
	"github.com/blazee5/cloud-drive/internal/storage"
	"github.com/blazee5/cloud-drive/internal/storage/postgres"
	"github.com/blazee5/cloud-drive/lib/logger"
	"github.com/gofiber/fiber/v2"
)

func main() {
	log := logger.NewLogger()

	cfg := config.LoadConfig()

	app := fiber.New()

	db := postgres.NewPostgres(cfg)
	storages := storage.NewStorage(db)
	services := service.NewService(log, storages)
	handlers := handler.NewHandler(log, services)

	handlers.InitRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)))

	defer log.Sync()
}
