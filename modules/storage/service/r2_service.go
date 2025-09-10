package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	// Thêm import này
	"go-api-starter/core/config"
	"go-api-starter/core/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func (s *StorageService) UploadToR2(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (string, string, error) {
	defer file.Close()
	configR2 := config.Get()

	ext := filepath.Ext(fileHeader.Filename)
	if ext == "" {
		ext = utils.GetExtensionFromContentType(fileHeader.Header.Get("Content-Type"))
	}

	// Generate unique name
	filename := utils.GenerateFileName(fileHeader.Filename, ext)

	// Upload
	_, err := s.r2Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(configR2.R2.Bucket),
		Key:         aws.String(filename),
		Body:        file,
		ContentType: aws.String(fileHeader.Header.Get("Content-Type")),
		ACL:         types.ObjectCannedACLPublicRead,
	})
	if err != nil {
		return "", "", fmt.Errorf("upload failed: %w", err)
	}

	// Return public URL - sử dụng custom domain
	var publicURL string
	if configR2.R2.PublicEndpoint != "" {
		publicURL = fmt.Sprintf("%s/%s", configR2.R2.PublicEndpoint, filename)
	} else {
		publicURL = fmt.Sprintf("%s/%s/%s", configR2.R2.Endpoint, configR2.R2.Bucket, filename)
	}

	return publicURL, filename, nil
}

func (s *StorageService) DeleteFromR2(ctx context.Context, key string) error {
	configR2 := config.Get()
	_, err := s.r2Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(configR2.R2.Bucket),
		Key:    aws.String(key),
	})
	return err
}
