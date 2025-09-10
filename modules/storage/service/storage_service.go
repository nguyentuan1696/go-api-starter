package service

import (
	"context"
	"path/filepath"
	"time"
	"go-api-starter/core/errors"
	"go-api-starter/core/logger"
	"go-api-starter/modules/storage/dto"
	"go-api-starter/modules/storage/entity"
	"go-api-starter/modules/storage/mapper"

	"mime/multipart"

	"github.com/google/uuid"
)

func (s *StorageService) SaveStorageWithRollback(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (*dto.StorageResponse, error) {
	// 1. Prepare storage entity
	storageEntity := &entity.Storage{
		ID:        uuid.New(),
		Name:      fileHeader.Filename,
		Size:      fileHeader.Size,
		Extension: filepath.Ext(fileHeader.Filename),
		Uploaded:  time.Now(),
	}

	// 2. Upload to R2
	publicURL, uploadKey, err := s.UploadToR2(ctx, file, fileHeader)
	if err != nil {
		logger.Error("StorageService:SaveStorage", err)
		return nil, errors.NewAppError(errors.ErrUploadToR2, "upload to r2 failed", err)
	}

	storageEntity.URL = publicURL

	// 3. Save to database with rollback on failure
	errSaveStorageDb := s.repo.SaveStorage(ctx, storageEntity)
	if errSaveStorageDb != nil {
		// Rollback: Delete uploaded file from R2
		if rollbackErr := s.DeleteFromR2(ctx, uploadKey); rollbackErr != nil {
			logger.Error("StorageService:SaveStorage:Rollback", rollbackErr)
		}
		logger.Error("StorageService:SaveStorage", errSaveStorageDb)
		return nil, errors.NewAppError(errors.ErrSaveStorageDb, "save storage to db failed", errSaveStorageDb)
	}

	return mapper.ToStorageResponse(storageEntity), nil
}
