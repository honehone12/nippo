package sss

import (
	"fmt"
	"io"
	"nippo/message/property"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

func JpgKey(userId []uint8) (*string, error) {
	id, err := uuid.FromBytes(userId)
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%s-%d.jpg", id.String(), time.Now().UnixMilli())
	return &key, nil
}

func PutJpg(
	props *property.MsgingProps,
	userId []uint8,
	body io.Reader,
	length int64,
) (string, error) {
	key, err := JpgKey(userId)
	if err != nil {
		return "", err
	}

	s3In := s3.PutObjectInput{
		Bucket:        props.Bucket(),
		Key:           key,
		Body:          body,
		ContentLength: &length,
	}

	_, err = props.S3().PutObject(props.Ctx(), &s3In)
	return *key, err
}

/*
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"
    "github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func main() {
    // AWSセッションを作成
    cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1")) // 適切なリージョンを指定
    if err != nil {
        fmt.Println("Error loading configuration:", err)
        return
    }

    // S3クライアントを作成
    client := s3.NewFromConfig(cfg)

    // Presignクライアントを作成
    presignClient := s3.NewPresignClient(client)

    // バケット名とオブジェクトキーを指定
    bucketName := "your-bucket-name"
    objectKey := "your-object-key"

    // GetObjectInputを作成
    getObjectInput := &s3.GetObjectInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(objectKey),
    }

    // 事前署名付きURLを生成（有効期限を設定）
    presignedURL, err := presignClient.PresignGetObject(context.TODO(), getObjectInput, func(options *s3.PresignOptions) {
        options.Expires = 15 * time.Minute // 有効期限15分
    })
    if err != nil {
        fmt.Println("Error generating presigned URL:", err)
        return
    }

    fmt.Println("Presigned URL:", presignedURL.URL)
}

*/
