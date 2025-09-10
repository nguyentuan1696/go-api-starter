package dto

import (
	"time"
	"go-api-starter/core/dto"

	"github.com/google/uuid"
)

type CategoryRequest struct {
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	Description string  `json:"description"`
	Thumbnail   string  `json:"thumbnail"`
	ParentID    *string `json:"parent_id"`
	SortOrder   int     `json:"sort_order"`
	IsActive    bool    `json:"is_active"`
}

type CategoryResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	Thumbnail   string    `json:"thumbnail"`
	ParentID    *string   `json:"parent_id"`
	ParentName  *string   `json:"parent_name"`
	SortOrder   int       `json:"sort_order"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PaginatedCategoryResponse = dto.Pagination[CategoryResponse]
