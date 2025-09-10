package dto

import (
	"time"
	"go-api-starter/core/dto"
)

type ShippingMethodRequest struct {
	Name                  string  `json:"name"`
	Description           string  `json:"description"`
	Provider              string  `json:"provider"`
	BaseCost              float64 `json:"base_cost"`
	CostPerKg             float64 `json:"cost_per_kg"`
	FreeShippingThreshold float64 `json:"free_shipping_threshold"`
	EstimatedDaysMin      int     `json:"estimated_days_min"`
	EstimatedDaysMax      int     `json:"estimated_days_max"`
	IsActive              bool    `json:"is_active"`
	SortOrder             int     `json:"sort_order"`
}

type ShippingMethodResponse struct {
	ID                    int       `json:"id"`
	Name                  string    `json:"name"`
	Description           string    `json:"description"`
	Provider              string    `json:"provider"`
	BaseCost              float64   `json:"base_cost"`
	CostPerKg             float64   `json:"cost_per_kg"`
	FreeShippingThreshold float64   `json:"free_shipping_threshold"`
	EstimatedDaysMin      int       `json:"estimated_days_min"`
	EstimatedDaysMax      int       `json:"estimated_days_max"`
	IsActive              bool      `json:"is_active"`
	SortOrder             int       `json:"sort_order"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type PaginatedShippingMethodDTO = dto.Pagination[ShippingMethodResponse]
