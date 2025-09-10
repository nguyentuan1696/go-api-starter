package dto

import (
	"time"
	"go-api-starter/core/dto"

	"github.com/google/uuid"
)

type BenefitRequest struct {
	Name               string `json:"name"`
	Slug               string `json:"-"`
	Description        string `json:"description"`
	ShortDescription   string `json:"short_description"`
	Category           string `json:"category"`
	TargetArea         string `json:"target_area"`
	EffectivenessLevel string `json:"effectiveness_level"`
	TimeToSeeResults   string `json:"time_to_see_results"`
	Color              string `json:"color"`
	Icon               string `json:"icon"`
	IsActive           bool   `json:"is_active"`
	IsFeatured         bool   `json:"is_featured"`
	SortOrder          int    `json:"sort_order"`
}

type BenefitResponse struct {
	ID                 uuid.UUID `json:"id"`
	Name               string    `json:"name"`
	Slug               string    `json:"slug"`
	Description        string    `json:"description"`
	ShortDescription   string    `json:"short_description"`
	Category           string    `json:"category"`
	TargetArea         string    `json:"target_area"`
	EffectivenessLevel string    `json:"effectiveness_level"`
	TimeToSeeResults   string    `json:"time_to_see_results"`
	Color              string    `json:"color"`
	Icon               string    `json:"icon"`
	IsActive           bool      `json:"is_active"`
	IsFeatured         bool      `json:"is_featured"`
	SortOrder          int       `json:"sort_order"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type PaginatedBenefitDTO = dto.Pagination[BenefitResponse]
