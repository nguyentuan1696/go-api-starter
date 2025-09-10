package mapper

import (
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/entity"
)

func ToSkinTypeEntity(req *dto.SkinTypeRequest) *entity.SkinType {
	return &entity.SkinType{
		Name:            req.Name,
		Slug:            req.Slug,
		Description:     req.Description,
		Characteristics: req.Characteristics,
		CareTips:        req.CareTips,
		Color:           req.Color,
		Icon:            req.Icon,
		IsActive:        req.IsActive,
	}
}

func ToSkinTypeDTO(entity *entity.SkinType) *dto.SkinTypeResponse {
	return &dto.SkinTypeResponse{
		ID:              entity.ID,
		Name:            entity.Name,
		Slug:            entity.Slug,
		Description:     entity.Description,
		Characteristics: entity.Characteristics,
		CareTips:        entity.CareTips,
		Color:           entity.Color,
		Icon:            entity.Icon,
		IsActive:        entity.IsActive,
	}
}

func ToSkinTypePaginationDTO(entity *entity.PaginatedSkinTypeEntity) *dto.PaginatedSkinTypeDTO {

	skinTypeResponses := make([]dto.SkinTypeResponse, len(entity.Items))
	for i, tag := range entity.Items {
		skinTypeResponses[i] = *ToSkinTypeDTO(&tag)
	}

	// Tính total pages
	totalPages := 0
	if entity.PageSize > 0 {
		totalPages = (entity.TotalItems + entity.PageSize - 1) / entity.PageSize
	}

	return &dto.PaginatedSkinTypeDTO{
		Items:      skinTypeResponses,
		TotalItems: entity.TotalItems,
		TotalPages: totalPages,
		PageNumber: entity.PageNumber,
		PageSize:   entity.PageSize,
	}
}
