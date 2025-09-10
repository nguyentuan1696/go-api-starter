package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserRoleRequest struct {
	UserID     uuid.UUID  `json:"user_id"`
	RoleID     uuid.UUID  `json:"role_id"`
	AssignedBy *uuid.UUID `json:"assigned_by"`
	AssignedAt time.Time  `json:"assigned_at"`
	ExpiresAt  *time.Time `json:"expires_at"`
	IsActive   bool       `json:"is_active"`
}

type UserRoleResponse struct {
	ID         uuid.UUID  `json:"id"`
	UserID     uuid.UUID  `json:"user_id"`
	RoleID     uuid.UUID  `json:"role_id"`
	AssignedBy *uuid.UUID `json:"assigned_by"`
	AssignedAt time.Time  `json:"assigned_at"`
	ExpiresAt  *time.Time `json:"expires_at"`
	IsActive   bool       `json:"is_active"`
}
