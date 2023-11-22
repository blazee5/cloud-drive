package postgres

import (
	"context"
	"fmt"
	"github.com/blazee5/cloud-drive/microservices/files/ent"
	"github.com/blazee5/cloud-drive/microservices/files/internal/config"
	_ "github.com/lib/pq"
	"log"
)

func New(cfg *config.Config) *ent.Client {
	client, err := ent.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName))

	if err != nil {
		log.Fatalf("error while connect to postgres: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
