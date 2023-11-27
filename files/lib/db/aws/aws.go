package aws

import (
	"github.com/blazee5/cloud-drive/microservices/files/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

func NewAWSClient(cfg *config.Config) *minio.Client {
	client, err := minio.New(cfg.AWS.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AWS.User, cfg.AWS.Password, ""),
		Secure: cfg.SSL,
	})

	if err != nil {
		log.Fatalf("error while connect to minio s3: %v", err)
	}

	return client
}
