package aws

import (
	"CVSeeker/pkg/cfg"
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/viper"
)

type IS3Client interface {
	UploadFile(ctx context.Context, bucket, key string, fileData []byte) (string, error)
}

type S3Client struct {
	Client *s3.Client
}

// NewS3Client creates a new S3 client using AWS configuration loaded from environment variables or config files
func NewS3Client(cfgReader *viper.Viper) (*S3Client, error) {
	awsRegion := cfgReader.GetString(cfg.AwsRegion)
	awsAccessKeyID := cfgReader.GetString(cfg.AwsAccessKey)
	awsSecretAccessKey := cfgReader.GetString(cfg.AwsSecretKey)

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithCredentialsProvider(aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     awsAccessKeyID,
				SecretAccessKey: awsSecretAccessKey,
				Source:          "Custom Viper Source",
			}, nil
		})),
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
func (aw *S3Client) UploadFile(ctx context.Context, bucket, key string, fileData []byte) (string, error) {
	// Directly use the provided ctx which is expected to be managed by the caller
	_, err := aw.Client.PutObject(ctx, &s3.PutObjectInput{
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
