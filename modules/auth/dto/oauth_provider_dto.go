package dto

import (
	"time"
	"go-api-starter/core/dto"

	"github.com/google/uuid"
)

type OAuthProviderRequest struct {
	Name         string   `json:"name"`
	DisplayName  string   `json:"display_name"`
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	RedirectURI  *string  `json:"redirect_uri"`
	Scopes       []string `json:"scopes"`
	AuthURL      *string  `json:"auth_url"`
	TokenURL     *string  `json:"token_url"`
	UserInfoURL  *string  `json:"user_info_url"`
	IsActive     bool     `json:"is_active"`
}

type OAuthProviderResponse struct {
	ID           uuid.Domain `json:"id"`
	Name         string      `json:"name"`
	DisplayName  string      `json:"display_name"`
	ClientID     string      `json:"client_id"`
	ClientSecret string      `json:"client_secret"`
	RedirectURI  *string     `json:"redirect_uri"`
	Scopes       []string    `json:"scopes"`
	AuthURL      *string     `json:"auth_url"`
	TokenURL     *string     `json:"token_url"`
	UserInfoURL  *string     `json:"user_info_url"`
	IsActive     bool        `json:"is_active"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

type PaginatedOAuthProviderDTO = dto.Pagination[OAuthProviderResponse]
