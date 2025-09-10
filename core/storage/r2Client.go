package storage

import (
	"fmt"
	"go-api-starter/core/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// NewS3Client tạo S3 client cho R2
func NewS3Client(config *config.Config) (*s3.Client, error) {
	// Validate config
	if config == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}

	// Validate R2 configuration with detailed error messages
	if config.R2.AccessKeyID == "" {
		return nil, fmt.Errorf("R2 access key ID is required. Please set APP_R2_ACCESS_KEY_ID environment variable")
	}
	if config.R2.SecretAccessKey == "" {
		return nil, fmt.Errorf("R2 secret access key is required. Please set APP_R2_SECRET_ACCESS_KEY environment variable")
	}
	if config.R2.Endpoint == "" {
		return nil, fmt.Errorf("R2 endpoint is required. Please set APP_R2_ENDPOINT environment variable")
	}
	if config.R2.Region == "" {
		return nil, fmt.Errorf("R2 region is required. Please set APP_R2_REGION environment variable")
	}
	if config.R2.Bucket == "" {
		return nil, fmt.Errorf("R2 bucket is required. Please set APP_R2_BUCKET environment variable")
	}

	cfg := aws.Config{
		Region: config.R2.Region,
		Credentials: credentials.NewStaticCredentialsProvider(
			config.R2.AccessKeyID,
			config.R2.SecretAccessKey,
			"",
		),
		EndpointResolverWithOptions: aws.EndpointResolverWithOptionsFunc(func(service, region string, opt ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:           config.R2.Endpoint,
				SigningRegion: config.R2.Region,
			}, nil
		}),
	}

	client := s3.NewFromConfig(cfg)
	return client, nil
}
