package server

import (
	"context"
	authHandler "github.com/blazee5/cloud-drive/api_gateway/internal/auth/handler/http"
	authService "github.com/blazee5/cloud-drive/api_gateway/internal/auth/service"
	fileHandler "github.com/blazee5/cloud-drive/api_gateway/internal/file/handler/http"
	fileService "github.com/blazee5/cloud-drive/api_gateway/internal/file/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

type Server struct {
	log        *zap.SugaredLogger
	httpServer *http.Server
}

func NewServer(log *zap.SugaredLogger) *Server {
	return &Server{log: log}
}

func (s *Server) Run(handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         os.Getenv("PORT"),
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) InitRoutes() *gin.Engine {
	router := gin.New()

	authServices := authService.NewService(s.log)
	authHandlers := authHandler.NewHandler(s.log, authServices)

	auth := router.Group("/auth")
	{
		auth.POST("/signup", authHandlers.SignUp)
		auth.POST("/signin", authHandlers.SignIn)
	}

	api := router.Group("/api")
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

	return router
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
