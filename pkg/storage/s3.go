package storage

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	appcfg "github.com/fatkulnurk/gostarter/config"
	"github.com/fatkulnurk/gostarter/pkg/interfaces"
	"github.com/fatkulnurk/gostarter/pkg/logging"
)

func NewS3Client(cfg appcfg.S3) (*s3.Client, error) {
	// Load konfigurasi AWS default dari environment, file config, dsb
	awscfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(cfg.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessKey, cfg.SecretKey, cfg.Session)),
	)
	if err != nil {
		logging.Fatalf("unable to load SDK config, %v", err)
		return nil, err
	}

	client := s3.NewFromConfig(awscfg)
	return client, nil
}

type S3Storage struct {
	client     *s3.Client
	bucketName string
}

func NewS3Storage(client *s3.Client, bucketName string) interfaces.IStorage {
	return &S3Storage{
		client:     client,
		bucketName: bucketName,
	}
}

func (s *S3Storage) Upload(ctx context.Context, input interfaces.UploadInput) (string, error) {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucketName),
		Key:         aws.String(input.FileName),
		Body:        bytes.NewReader(input.Content),
		ContentType: aws.String(input.MimeType),
		ACL:         types.ObjectCannedACLPublicRead,
	})
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s.bucketName, input.FileName)
	return url, nil
}

func (s *S3Storage) Delete(ctx context.Context, path string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(path),
	})
	return err
}
