package aws

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"time"
)

type IS3Client interface {
	UploadFile(bucket, key string, fileData []byte) (string, error)
}

type S3Client struct {
	Client *s3.Client
}

// NewS3Client creates a new S3 client using default AWS configuration
func NewS3Client() (*S3Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-west-2"), // Specify the region here
	)
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	s3Client := s3.NewFromConfig(cfg)
	return &S3Client{
		Client: s3Client,
	}, nil
}

// UploadFile uploads file data to the specified S3 bucket and returns the URL of the uploaded file
func (c *S3Client) UploadFile(bucket, key string, fileData []byte) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Creating an uploader with the session and default options
	_, err := c.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(fileData),
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %v", err)
	}

	// Return the URL of the uploaded file
	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, key), nil
}
