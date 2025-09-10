package entity

import (
	"time"
	"go-api-starter/core/entity"
)

type User struct {
	Email           *string    `db:"email"`
	Phone           string     `db:"phone"`
	Username        *string    `db:"username"`
	Password        string     `db:"password"`
	EmailVerifiedAt *time.Time `db:"email_verified_at"`
	PhoneVerifiedAt *time.Time `db:"phone_verified_at"`
	LockedUntil     *time.Time `db:"locked_until"`
	IsActive        bool       `db:"is_active"`
	entity.BaseEntity
}

type UserDetail struct {
	ID          string  `db:"id"`
	Email       *string `db:"email"`
	Phone       string  `db:"phone"`
	Username    *string `db:"username"`
	IsActive    bool    `db:"is_active"`
	CreatedAt   string  `db:"created_at"`
	DisplayName *string `db:"display_name"`
	FullName    *string `db:"full_name"`
	Avatar      *string `db:"avatar"`
	DateOfBirth *string `db:"date_of_birth"`
	Gender      *string `db:"gender"`
	Roles       *string `db:"roles"`
}

type PaginatedUserEntity = entity.Pagination[User]
