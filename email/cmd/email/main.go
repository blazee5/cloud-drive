package main

import (
	"github.com/blazee5/cloud-drive/email/internal/app"
	"github.com/blazee5/cloud-drive/email/lib/logger"
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
