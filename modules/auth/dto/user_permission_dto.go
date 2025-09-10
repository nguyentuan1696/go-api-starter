package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserPermissionRequest struct {
	UserID       uuid.UUID  `json:"user_id"`
	PermissionID uuid.UUID  `json:"permission_id"`
	Granted      bool       `json:"granted"`
	GrantedBy    *uuid.UUID `json:"granted_by"`
	GrantedAt    time.Time  `json:"granted_at"`
	ExpiresAt    *time.Time `json:"expires_at"`
}

type UserPermissionResponse struct {
	ID           uuid.UUID  `json:"id"`
	UserID       uuid.UUID  `json:"user_id"`
	PermissionID uuid.UUID  `json:"permission_id"`
	Granted      bool       `json:"granted"`
	GrantedBy    *uuid.UUID `json:"granted_by"`
	GrantedAt    time.Time  `json:"granted_at"`
	ExpiresAt    *time.Time `json:"expires_at"`
}
