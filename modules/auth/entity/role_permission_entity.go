package entity

import (
	"time"

	"github.com/google/uuid"
)

type RolePermission struct {
	RoleID       uuid.UUID  `db:"role_id"`
	PermissionID uuid.UUID  `db:"permission_id"`
	GrantedBy    *uuid.UUID `db:"granted_by"`
	GrantedAt    time.Time  `db:"granted_at"`
}
