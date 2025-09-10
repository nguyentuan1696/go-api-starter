package entity

import (
	"time"
	"go-api-starter/core/entity"

	"github.com/google/uuid"
)

type SocialLogin struct {
	UserID           uuid.UUID  `db:"user_id"`
	ProviderID       uuid.UUID  `db:"provider_id"`
	ProviderUserID   uuid.UUID  `db:"provider_user_id"`
	ProviderUsername *string    `db:"provider_username"`
	ProviderEmail    *string    `db:"provider_email"`
	AccessToken      *string    `db:"access_token"`
	RefreshToken     *string    `db:"refresh_token"`
	TokenExpiresAt   *time.Time `db:"token_expires_at"`
	ProviderData     any        `db:"provider_data"`
	LastLoginAt      *time.Time `db:"last_login_at"`
	IsActive         bool       `db:"is_active"`
	entity.BaseEntity
}
