package dto

import (
	"go-api-starter/core/dto"

	"github.com/google/uuid"
)

type SkinTypeRequest struct {
	Name            string `json:"name"`
	Slug            string `json:"-"` // slug will be generated from name
	Description     string `json:"description"`
	Characteristics string `json:"characteristics"`
	CareTips        string `json:"care_tips"`
	Color           string `json:"color"`
	Icon            string `json:"icon"`
	IsActive        bool   `json:"is_active"`
}

type SkinTypeResponse struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	Slug            string    `json:"slug"`
	Description     string    `json:"description"`
	Characteristics string    `json:"characteristics"`
	CareTips        string    `json:"care_tips"`
	Color           string    `json:"color"`
	Icon            string    `json:"icon"`
	IsActive        bool      `json:"is_active"`
}

type PaginatedSkinTypeDTO = dto.Pagination[SkinTypeResponse]
