package entity

import (
	"time"
	"go-api-starter/core/entity"

	"github.com/google/uuid"
)

type UserProfile struct {
	UserID      uuid.UUID  `db:"user_id"`
	FullName    *string    `db:"full_name"`
	DisplayName *string    `db:"display_name"`
	Avatar      *string    `db:"avatar"`
	DateOfBirth *time.Time `db:"date_of_birth"`
	Gender      *string    `db:"gender"`
	entity.BaseEntity
}
