package entity

import (
	"go-api-starter/core/entity"
)

type OAuthProvider struct {
	Name         string   `db:"name"`
	DisplayName  string   `db:"display_name"`
	ClientID     string   `db:"client_id"`
	ClientSecret string   `db:"client_secret"`
	RedirectURI  *string  `db:"redirect_uri"`
	Scopes       []string `db:"scopes"`
	AuthURL      *string  `db:"auth_url"`
	TokenURL     *string  `db:"token_url"`
	UserInfoURL  *string  `db:"user_info_url"`
	IsActive     bool     `json:"is_active"`
	entity.BaseEntity
}
