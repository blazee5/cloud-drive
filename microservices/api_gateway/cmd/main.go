package main

import (
	"github.com/blazee5/cloud-drive/microservices/api_gateway/internal/server"
	"github.com/blazee5/cloud-drive/microservices/api_gateway/lib/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	log := logger.NewLogger()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	s := server.NewServer(log)
	s.InitRoutes(router)

	router.Run(os.Getenv("PORT"))
}
