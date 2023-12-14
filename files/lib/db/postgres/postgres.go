package postgres

import (
	"context"
	"fmt"
	"github.com/blazee5/cloud-drive/files/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"log"
)

func NewPgxConn(ctx context.Context, cfg *config.Config) *pgxpool.Pool {
	db, err := pgxpool.New(ctx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName, cfg.DB.SSLMode))

	if err != nil {
		log.Fatalf("error while connect to postgres: %v", err)
	}

	return db
}
