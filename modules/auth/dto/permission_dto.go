package dto

import (
	"time"
	"go-api-starter/core/dto"

	"github.com/google/uuid"
)

type PermissionRequest struct {
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	Resource    string  `json:"resource"`
	Action      string  `json:"action"`
	Description *string `json:"description"`
	IsSystem    bool    `json:"is_system"`
}

type PermissionResponse struct {
	ID          uuid.UUID `json:"id,omitzero"`
	Name        string    `json:"name,omitempty"`
	Slug        string    `json:"slug,omitempty"`
	Resource    string    `json:"resource,omitempty"`
	Action      string    `json:"action,omitempty"`
	Description *string   `json:"description,omitempty"`
	IsSystem    bool      `json:"is_system,omitzero"`
	CreatedAt   time.Time `json:"created_at,omitzero"`
	UpdatedAt   time.Time `json:"updated_at,omitzero"`
}

type PaginatedPermissionDTO = dto.Pagination[PermissionResponse]
