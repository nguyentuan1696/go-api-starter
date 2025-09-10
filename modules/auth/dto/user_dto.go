package dto

import (
	"time"
	"go-api-starter/core/dto"

	"github.com/google/uuid"
)

type RegisterRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequest struct {
	Identifier string `json:"identifiers"` // phone, username, email
	Password   string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ResetPasswordRequest struct {
	Token           string `json:"token"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type ChangePasswordRequest struct {
	Password        string `json:"password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
	OTP             string `json:"otp"`
}

type ForgotPasswordRequest struct {
	Identifier string `json:"identifier"`
}

type ForgotPasswordResponse struct {
	UserId uuid.UUID `json:"user_id"`
}

type VerifyOTPRequest struct {
	UserID uuid.UUID `json:"user_id"`
	OTP    string    `json:"otp"`
}

type VerifyOTPResponse struct {
	Token string `json:"token"`
}

type ResetPassword struct {
	Token             string `json:"token"`
	NewPassword       string `json:"new_password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type UserRequest struct {
	Email           string     `json:"email"`
	Phone           string     `json:"phone"`
	Username        *string    `json:"username"`
	Password        *string    `json:"-"`
	EmailVerifiedAt *time.Time `json:"-"`
	PhoneVerifiedAt *time.Time `json:"-"`
	LockedUntil     *time.Time `json:"-"`
	IsActive        bool       `json:"is_active"`
}

type UserResponse struct {
	ID              uuid.UUID  `json:"id"`
	Email           *string    `json:"email"`
	Phone           string     `json:"phone"`
	Username        *string    `json:"username"`
	Password        string     `json:"-"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	PhoneVerifiedAt *time.Time `json:"phone_verified_at"`
	LockedUntil     *time.Time `json:"locked_until"`
	IsActive        bool       `json:"is_active"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

type UserDetailDTO struct {
	ID          string  `json:"id"`
	Email       *string `json:"email"`
	Phone       string  `json:"phone"`
	Username    *string `json:"username"`
	IsActive    bool    `json:"is_active"`
	CreatedAt   string  `json:"created_at"`
	DisplayName *string `json:"display_name"`
	FullName    *string `json:"full_name"`
	Avatar      *string `json:"avatar"`
	DateOfBirth *string `json:"date_of_birth"`
	Gender      *string `json:"gender"`
	Roles       *string `json:"roles"`
}

type PaginatedUserDTO = dto.Pagination[UserResponse]
