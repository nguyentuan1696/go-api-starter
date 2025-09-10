package dto

import (
	"time"
	"go-api-starter/core/dto"

	"github.com/google/uuid"
)

type SocialLoginRequest struct {
	UserID           uuid.UUID  `json:"user_id"`
	ProviderID       uuid.UUID  `json:"provider_id"`
	ProviderUserID   uuid.UUID  `json:"provider_user_id"`
	ProviderUsername *string    `json:"provider_username"`
	ProviderEmail    *string    `json:"provider_email"`
	AccessToken      *string    `json:"access_token"`
	RefreshToken     *string    `json:"refresh_token"`
	TokenExpiresAt   *time.Time `json:"token_expires_at"`
	ProviderData     any        `json:"provider_data"`
	LastLoginAt      *time.Time `json:"last_login_at"`
	IsActive         bool       `json:"is_active"`
}

type SocialLoginResponse struct {
	ID               uuid.UUID  `json:"id"`
	UserID           uuid.UUID  `json:"user_id"`
	ProviderID       uuid.UUID  `json:"provider_id"`
	ProviderUserID   uuid.UUID  `json:"provider_user_id"`
	ProviderUsername *string    `json:"provider_username"`
	ProviderEmail    *string    `json:"provider_email"`
	AccessToken      *string    `json:"access_token"`
	RefreshToken     *string    `json:"refresh_token"`
	TokenExpiresAt   *time.Time `json:"token_expires_at"`
	ProviderData     any        `json:"provider_data"`
	LastLoginAt      *time.Time `json:"last_login_at"`
	IsActive         bool       `json:"is_active"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

type PaginatedSocialLoginDTO = dto.Pagination[SocialLoginResponse]
