package dto

import (
	"time"
	"go-api-starter/core/dto"

	"github.com/google/uuid"
)

type RoleRequest struct {
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	Description *string `json:"description"`
	IsSystem    bool    `json:"is_system"`
	IsActive    bool    `json:"is_active"`
}

type RoleResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description *string   `json:"description"`
	IsSystem    bool      `json:"is_system"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PaginatedRoleDTO = dto.Pagination[RoleResponse]
