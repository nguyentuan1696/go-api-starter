package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserRole struct {
	ID         uuid.UUID  `db:"id"`
	UserID     uuid.UUID  `db:"user_id"`
	RoleID     uuid.UUID  `db:"role_id"`
	AssignedBy *uuid.UUID `db:"assigned_by"`
	AssignedAt time.Time  `db:"assigned_at"`
	ExpiresAt  *time.Time `db:"expires_at"`
	IsActive   bool       `db:"is_active"`
}
