package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserPermission struct {
	ID           uuid.UUID  `db:"id"`
	UserID       uuid.UUID  `db:"user_id"`
	PermissionID uuid.UUID  `db:"permission_id"`
	Granted      bool       `db:"granted"`
	GrantedBy    *uuid.UUID `db:"granted_by"`
	GrantedAt    time.Time  `db:"granted_at"`
	ExpiresAt    *time.Time `db:"expires_at"`
}
