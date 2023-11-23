package aws

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"io"
)

type Storage struct {
	client *minio.Client
}

func NewStorage(client *minio.Client) *Storage {
	return &Storage{client: client}
}

func (s *Storage) SaveFile(ctx context.Context, bucket, fileName, contentType string, chunk []byte) error {
	options := minio.PutObjectOptions{
		ContentType:  contentType,
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
	}

	file := bytes.NewReader(chunk)

	_, err := s.client.PutObject(ctx, bucket, fileName, file, file.Size(), options)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) DownloadFile(ctx context.Context, bucket string, fileName string) ([]byte, error) {
	options := minio.GetObjectOptions{}

	file, err := s.client.GetObject(ctx, bucket, fileName, options)

	if err != nil {
		return nil, err
	}

	chunk, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	return chunk, nil
}

func (s *Storage) RemoveObject(ctx context.Context, bucket string, fileName string) error {
	if err := s.client.RemoveObject(ctx, bucket, fileName, minio.RemoveObjectOptions{}); err != nil {
		return err
	}

	return nil
}
