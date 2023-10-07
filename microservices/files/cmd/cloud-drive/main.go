package main

import (
	"github.com/blazee5/cloud-drive/microservices/files/internal/app"
	"github.com/blazee5/cloud-drive/microservices/files/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	app.Run(cfg)
}
