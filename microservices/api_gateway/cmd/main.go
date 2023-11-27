package main

import (
	"github.com/blazee5/cloud-drive/microservices/api_gateway/internal/app"
	"github.com/blazee5/cloud-drive/microservices/api_gateway/lib/logger"
	"github.com/joho/godotenv"
)

func main() {
	log := logger.NewLogger()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Run(log)
}
