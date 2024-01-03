package aws

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"go.opentelemetry.io/otel/trace"
	"io"
)

type Storage struct {
	client *minio.Client
	tracer trace.Tracer
}

func NewStorage(client *minio.Client, tracer trace.Tracer) *Storage {
	return &Storage{client: client, tracer: tracer}
}

func (s *Storage) SaveFile(ctx context.Context, bucket, fileName, contentType string, chunk []byte) error {
	ctx, span := s.tracer.Start(ctx, "fileAWSStorage.SaveFile")
	defer span.End()

	options := minio.PutObjectOptions{
		ContentType:  contentType,
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
	}

	file := bytes.NewReader(chunk)

	bucketExists, err := s.client.BucketExists(ctx, bucket)

	if err != nil {
		return err
	}

	if !bucketExists {
		err := s.client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})

		if err != nil {
			return err
		}
	}

	_, err = s.client.PutObject(ctx, bucket, fileName, file, file.Size(), options)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) DownloadFile(ctx context.Context, bucket string, fileName string) ([]byte, error) {
	ctx, span := s.tracer.Start(ctx, "fileAWSStorage.DownloadFile")
	defer span.End()

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

func (s *Storage) UpdateFile(ctx context.Context, bucket, oldName, newName string) error {
	ctx, span := s.tracer.Start(ctx, "fileAWSStorage.UpdateFile")
	defer span.End()

	copyDestOpts := minio.CopyDestOptions{
		Bucket: bucket,
		Object: newName,
	}

	copySrcOpts := minio.CopySrcOptions{
		Bucket: bucket,
		Object: oldName,
	}

	if _, err := s.client.CopyObject(ctx, copyDestOpts, copySrcOpts); err != nil {
		return err
	}

	if err := s.client.RemoveObject(ctx, bucket, oldName, minio.RemoveObjectOptions{}); err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteFile(ctx context.Context, bucket, fileName string) error {
	ctx, span := s.tracer.Start(ctx, "fileAWSStorage.DeleteFile")
	defer span.End()

	if err := s.client.RemoveObject(ctx, bucket, fileName, minio.RemoveObjectOptions{}); err != nil {
		return err
	}

	return nil
}
