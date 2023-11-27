package main

import (
	"github.com/blazee5/cloud-drive/microservices/auth/internal/app"
	"github.com/blazee5/cloud-drive/microservices/auth/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	app.Run(cfg)
}
