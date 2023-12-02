package main

import (
	"github.com/blazee5/cloud-drive/auth/internal/app"
	"github.com/blazee5/cloud-drive/auth/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	app.Run(cfg)
}
