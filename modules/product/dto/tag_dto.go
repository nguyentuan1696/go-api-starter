package dto

import (
	"time"
	"go-api-starter/core/dto"

	"github.com/google/uuid"
)

type TagRequest struct {
	Name        string `json:"name"`
	Slug        string `json:"-"` // slug sẽ được tự động generate từ name
	Description string `json:"description"`
	Color       string `json:"color"` //  hex color code
	Icon        string `json:"icon"`
	IsActive    bool   `json:"is_active"`
}

type TagResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	Color       string    `json:"color"`
	Icon        string    `json:"icon"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PaginatedTagDTO = dto.Pagination[TagResponse]
