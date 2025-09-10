package repository

import (
	"context"
	"go-api-starter/core/logger"
	"go-api-starter/core/params"
	"go-api-starter/modules/storage/entity"
)

func (r *StorageRepository) GetStorages(ctx context.Context, params params.QueryParams) (*entity.PaginatedStorageResponse, error) {
	return nil, nil
}

func (r *StorageRepository) SaveStorage(ctx context.Context, storage *entity.Storage) error {
	query := `
		INSERT INTO storages (id, name, url, size, extension, uploaded)
		VALUES (:id, :name, :url, :size, :extension, :uploaded)
	`
	_, err := r.DB.NamedExecContext(ctx, query, storage)
	if err != nil {
		logger.Error("StorageRepository:SaveStorage:", err)
		return err
	}

	return nil
}
