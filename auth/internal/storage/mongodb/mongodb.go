package mongodb

import (
	"context"
	"fmt"
	"github.com/blazee5/cloud-drive/microservices/auth/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func NewMongoDB(ctx context.Context, cfg *config.Config) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", cfg.DBHost, cfg.DBPort)))

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	return client
}
