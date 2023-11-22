package aws

import (
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"os"
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

	info, err := s.client.PutObject(ctx, bucket, fileName, file, file.Size(), options)
	fmt.Println(info)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) DownloadFile(ctx context.Context, bucket string, fileName string) ([]byte, error) {
	options := minio.GetObjectOptions{}

	err := s.client.FGetObject(ctx, bucket, fileName, "./temp/"+fileName, options)

	if err != nil {
		return nil, err
	}

	var file []byte

	fileReader, err := os.Open("./temp/" + fileName)

	if err != nil {
		return nil, err
	}

	_, err = fileReader.Read(file)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *Storage) RemoveObject(ctx context.Context, bucket string, fileName string) error {
	if err := s.client.RemoveObject(ctx, bucket, fileName, minio.RemoveObjectOptions{}); err != nil {
		return err
	}

	return nil
}
