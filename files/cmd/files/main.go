package main

import (
	"github.com/blazee5/cloud-drive/files/internal/app"
	"github.com/blazee5/cloud-drive/files/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	app.Run(cfg)
}
