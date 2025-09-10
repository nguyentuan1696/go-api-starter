package service

import (
	"context"
	"mime/multipart"
	"go-api-starter/modules/storage/dto"
	"go-api-starter/modules/storage/repository"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type StorageService struct {
	repo     repository.StorageRepositoryInterface
	r2Client *s3.Client
}

func NewStorageService(repo repository.StorageRepositoryInterface, r2Client *s3.Client) *StorageService {
	return &StorageService{repo: repo, r2Client: r2Client}
}

type StorageServiceInterface interface {
	UploadToR2(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (string, string, error)
	SaveStorageWithRollback(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (*dto.StorageResponse, error)
	DeleteFromR2(ctx context.Context, key string) error
}
