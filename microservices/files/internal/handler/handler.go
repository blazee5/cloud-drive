package handler

import (
	"github.com/blazee5/cloud-drive/microservices/files/internal/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Handler struct {
	log     *zap.SugaredLogger
	service *service.Service
}

func NewHandler(log *zap.SugaredLogger, service *service.Service) *Handler {
	return &Handler{log: log, service: service}
}

func (h *Handler) InitRoutes(app *fiber.App) {
	api := fiber.Router(app).Group("/api")
	{
		api.Post("/upload", h.Upload)
		api.Get("/download/:filename", h.Download)
	}
}
