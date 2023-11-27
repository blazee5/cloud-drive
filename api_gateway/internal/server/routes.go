package server

import (
	authHandler "github.com/blazee5/cloud-drive/microservices/api_gateway/internal/auth/handler/http"
	authService "github.com/blazee5/cloud-drive/microservices/api_gateway/internal/auth/service"
	fileHandler "github.com/blazee5/cloud-drive/microservices/api_gateway/internal/file/handler/http"
	fileService "github.com/blazee5/cloud-drive/microservices/api_gateway/internal/file/service"
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

	api := c.Group("/api")
	{
		fileServices := fileService.NewService(s.log)
		fileHandlers := fileHandler.NewHandler(s.log, fileServices)

		files := api.Group("/files", s.UserMiddleware)
		{
			files.GET("/", fileHandlers.GetUserFiles)
			files.POST("/", fileHandlers.UploadFile)
			files.GET("/:id", fileHandlers.DownloadFile)
			files.PUT("/:id", fileHandlers.UpdateFile)
			files.DELETE("/:id", fileHandlers.DeleteFile)
		}
	}
}
