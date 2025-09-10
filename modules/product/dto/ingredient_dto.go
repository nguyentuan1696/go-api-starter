package dto

import (
	"time"
	"go-api-starter/core/dto"

	"github.com/google/uuid"
)

type IngredientRequest struct {
	Name         string `json:"name"`
	Slug         string `json:"-"`
	InciName     string `json:"inci_name"`
	Description  string `json:"description"`
	Origin       string `json:"origin"`
	Function     string `json:"function"`
	CasNumber    string `json:"cas_number"`
	EwgScore     int    `json:"ewg_score"`
	IsRestricted bool   `json:"is_restricted"`
	IsBanned     bool   `json:"is_banned"`
}

type IngredientResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Slug         string    `json:"slug"`
	InciName     string    `json:"inci_name"`
	Description  string    `json:"description"`
	Origin       string    `json:"origin"`
	Function     string    `json:"function"`
	CasNumber    string    `json:"cas_number"`
	EwgScore     int       `json:"ewg_score"`
	IsRestricted bool      `json:"is_restricted"`
	IsBanned     bool      `json:"is_banned"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type PaginatedIngredientDTO = dto.Pagination[IngredientResponse]
