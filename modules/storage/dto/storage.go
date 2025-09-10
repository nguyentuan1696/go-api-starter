package dto

import (
	"time"
	"go-api-starter/core/dto"
)

type StorageRequest struct {
	Name      string `json:"name"`
	URL       string `json:"url"`
	Size      int64  `json:"size"`
	Extension string `json:"extension"`
}

type StorageResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	Size      int64     `json:"size"`
	Extension string    `json:"extension"`
	Uploaded  time.Time `json:"uploaded"`
}

type PaginatedStorageResponse = dto.Pagination[StorageResponse]
