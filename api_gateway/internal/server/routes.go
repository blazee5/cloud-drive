package server

import (
	"context"
	authHandler "github.com/blazee5/cloud-drive/api_gateway/internal/auth/handler/http"
	authService "github.com/blazee5/cloud-drive/api_gateway/internal/auth/service"
	fileHandler "github.com/blazee5/cloud-drive/api_gateway/internal/file/handler/http"
	fileService "github.com/blazee5/cloud-drive/api_gateway/internal/file/service"
	"github.com/blazee5/cloud-drive/api_gateway/lib/tracer"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

type Server struct {
	log        *zap.SugaredLogger
	httpServer *http.Server
	tracer     *tracer.JaegerTracing
}

func NewServer(log *zap.SugaredLogger, tracer *tracer.JaegerTracing) *Server {
	return &Server{log: log, tracer: tracer}
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
	router.Use(cors.Default())

	authServices := authService.NewService(s.log, s.tracer)
	authHandlers := authHandler.NewHandler(s.log, authServices, s.tracer.Tracer)

	auth := router.Group("/auth")
	{
		auth.POST("/signup", authHandlers.SignUp)
		auth.POST("/signin", authHandlers.SignIn)
		auth.GET("/activate", authHandlers.ActivateAccount)
	}

	api := router.Group("/api")
	{
		fileServices := fileService.NewService(s.log, s.tracer)
		fileHandlers := fileHandler.NewHandler(s.log, fileServices, s.tracer.Tracer)

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
