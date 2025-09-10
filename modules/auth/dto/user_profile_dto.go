package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserProfileRequest struct {
	UserID      uuid.UUID  `json:"user_id"`
	FirstName   *string    `json:"first_name"`
	LastName    *string    `json:"last_name"`
	DisplayName *string    `json:"display_name"`
	Avatar      *string    `json:"avatar"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	Gender      *string    `json:"gender"`
}

type UserProfileResponse struct {
	UserID      uuid.UUID  `json:"user_id"`
	FirstName   *string    `json:"first_name"`
	LastName    *string    `json:"last_name"`
	DisplayName *string    `json:"display_name"`
	Avatar      *string    `json:"avatar"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	Gender      *string    `json:"gender"`
}

