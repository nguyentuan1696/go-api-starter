package repository

import (
	"context"
	"go-api-starter/core/database"
	"go-api-starter/modules/storage/entity"
)

type StorageRepository struct {
	DB database.Database
}

func NewStorageRepository(db database.Database) *StorageRepository {
	return &StorageRepository{DB: db}
}

type StorageRepositoryInterface interface {
	SaveStorage(ctx context.Context, storage *entity.Storage) error
}
