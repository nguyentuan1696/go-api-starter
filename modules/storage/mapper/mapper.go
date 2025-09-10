package mapper

import (
	"go-api-starter/modules/storage/dto"
	"go-api-starter/modules/storage/entity"
)

func ToStorageResponse(storage *entity.Storage) *dto.StorageResponse {
	return &dto.StorageResponse{
		ID:        storage.ID.String(),
		Name:      storage.Name,
		URL:       storage.URL,
		Size:      storage.Size,
		Extension: storage.Extension,
		Uploaded:  storage.Uploaded,
	}
}

func ToStorageEntity(storage *dto.StorageRequest) *entity.Storage {
	return &entity.Storage{
		Name:      storage.Name,
		URL:       storage.URL,
		Size:      storage.Size,
		Extension: storage.Extension,
	}
}

func ToStoragePaginationResponse(entity *entity.PaginatedStorageResponse) *dto.PaginatedStorageResponse {
	// Convert từng category entity sang category response
	storageResponses := make([]dto.StorageResponse, len(entity.Items))
	for i, storage := range entity.Items {
		storageResponses[i] = *ToStorageResponse(&storage)
	}

	// Tính total pages
	totalPages := 0
	if entity.PageSize > 0 {
		totalPages = (entity.TotalItems + entity.PageSize - 1) / entity.PageSize
	}

	return &dto.PaginatedStorageResponse{
		Items:      storageResponses,
		TotalItems: entity.TotalItems,
		TotalPages: totalPages,
		PageNumber: entity.PageNumber,
		PageSize:   entity.PageSize,
	}
}
