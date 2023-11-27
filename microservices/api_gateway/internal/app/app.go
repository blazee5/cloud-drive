package app

import (
	"github.com/blazee5/cloud-drive/microservices/api_gateway/internal/server"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

func Run(log *zap.SugaredLogger) {
	router := gin.Default()

	s := server.NewServer(log)
	s.InitRoutes(router)

	router.Run(os.Getenv("PORT"))
}
