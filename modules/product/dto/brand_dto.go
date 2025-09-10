package dto

import (
	"time"
	"go-api-starter/core/dto"

	"github.com/google/uuid"
)

type BrandRequest struct {
	Name        string `json:"name"`
	Slug        string `json:"-"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Website     string `json:"website"`
	Country     string `json:"country"`
	FoundedYear int    `json:"founded_year"`
	IsActive    bool   `json:"is_active"`
	IsFeatured  bool   `json:"is_featured"`
	SortOrder   int    `json:"sort_order"`
}

type BrandResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	Logo        string    `json:"logo"`
	Website     string    `json:"website"`
	Country     string    `json:"country"`
	FoundedYear int       `json:"founded_year"`
	IsActive    bool      `json:"is_active"`
	IsFeatured  bool      `json:"is_featured"`
	SortOrder   int       `json:"sort_order"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PaginatedBrandResponse = dto.Pagination[BrandResponse]
