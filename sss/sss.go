package sss

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

func S3Client(ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)
	return client, nil
}

func Bucket() *string {
	bucket := os.Getenv("BUCKET")
	return &bucket
}

func JpgKey(userId []uint8) (*string, error) {
	id, err := uuid.FromBytes(userId)
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%s-%d.jpg", id.String(), time.Now().UnixMilli())
	return &key, nil
}

func PutJpg(ctx context.Context, userId []uint8, body io.Reader, length int64) (string, error) {
	s3Client, err := S3Client(ctx)
	if err != nil {
		return "", err
	}

	key, err := JpgKey(userId)
	if err != nil {
		return "", err
	}

	s3In := s3.PutObjectInput{
		Bucket:        Bucket(),
		Key:           key,
		Body:          body,
		ContentLength: &length,
	}

	_, err = s3Client.PutObject(ctx, &s3In)
	return *key, err
}
