package sss

import (
	"fmt"
	"io"
	"nippo/message/msgprops"
	"nippo/timeconv"
	"nippo/view/viewprops"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

func JpgKey(userId []uint8) (string, error) {
	id, err := uuid.FromBytes(userId)
	if err != nil {
		return "", err
	}

	now := timeconv.NowJst()
	key := fmt.Sprintf(
		"%d-%d-%d-%d-%d-%d-%d_%s.jpg",
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second(),
		now.Nanosecond(),
		id,
	)
	return key, nil
}

func PutJpg(
	props *msgprops.Props,
	userId []uint8,
	body io.Reader,
	length int64,
) (string, error) {
	key, err := JpgKey(userId)
	if err != nil {
		return "", err
	}

	s3Put := s3.PutObjectInput{
		Bucket:        props.Bucket(),
		Key:           &key,
		Body:          body,
		ContentLength: &length,
	}

	_, err = props.S3().PutObject(props.Ctx(), &s3Put)
	if err != nil {
		return "", err
	}

	return key, err
}

func GetPresigned(props *viewprops.Props, key string) (string, error) {
	s3Get := s3.GetObjectInput{
		Bucket: props.Bucket(),
		Key:    &key,
	}

	presigned, err := props.S3Presign().PresignGetObject(
		props.Ctx(),
		&s3Get,
		s3.WithPresignExpires(time.Minute*30),
	)
	if err != nil {
		return "", err
	}

	return presigned.URL, nil
}
