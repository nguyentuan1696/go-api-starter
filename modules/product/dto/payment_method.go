package dto

import (
	"time"
	"go-api-starter/core/dto"
)

type PaymentMethodRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Provider    string `json:"provider"`
	Type        string `json:"type"`
	IsActive    bool   `json:"is_active"`
	SortOrder   int    `json:"sort_order"`
}

type PaymentMethodResponse struct {
	ID          int    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Provider    string    `json:"provider"`
	Type        string    `json:"type"`
	IsActive    bool      `json:"is_active"`
	SortOrder   int       `json:"sort_order"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PaginatedPaymentMethodDTO = dto.Pagination[PaymentMethodResponse]
