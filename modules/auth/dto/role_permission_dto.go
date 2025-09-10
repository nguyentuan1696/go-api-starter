package dto

import (
	"time"
	"go-api-starter/core/dto"

	"github.com/google/uuid"
)

type RolePermissionRequest struct {
	RoleID       uuid.UUID   `json:"role_id"`
	PermissionID []uuid.UUID `json:"permission_id"`
	GrantedBy    uuid.UUID   `json:"-"`
}

type RolePermissionResponse struct {
	ID           uuid.UUID   `json:"id"`
	RoleID       uuid.UUID   `json:"role_id"`
	PermissionID []uuid.UUID `json:"permission_id"`
	GrantedBy    *uuid.UUID  `json:"granted_by"`
	GrantedAt    time.Time   `json:"granted_at"`
}

type RolePermissionsDTO struct {
	Permissions []PermissionResponse `json:"permissions"`
}

type PaginatedRolePermissionDTO = dto.Pagination[RolePermissionResponse]
