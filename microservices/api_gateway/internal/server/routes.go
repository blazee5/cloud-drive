package server

import (
	authHandler "github.com/blazee5/cloud-drive/microservices/api_gateway/internal/auth/handler/http"
	authService "github.com/blazee5/cloud-drive/microservices/api_gateway/internal/auth/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	log *zap.SugaredLogger
}

func NewServer(log *zap.SugaredLogger) *Server {
	return &Server{log: log}
}

func (s *Server) InitRoutes(c *gin.Engine) {
	authServices := authService.NewService(s.log)
	authHandlers := authHandler.NewHandler(s.log, authServices)

	auth := c.Group("/auth")
	{
		auth.POST("/signup", authHandlers.SignUp)
		auth.POST("/signin", authHandlers.SignIn)
	}
}
