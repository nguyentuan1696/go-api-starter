package mapper

import (
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/entity"
)

func ToTagEntity(dto *dto.TagRequest) *entity.Tag {

	if dto == nil {
		return nil
	}

	return &entity.Tag{
		Name:        dto.Name,
		Slug:        dto.Slug,
		Description: dto.Description,
		Color:       dto.Color,
		Icon:        dto.Icon,
		IsActive:    dto.IsActive,
	}
}

func ToTagDTO(entity *entity.Tag) *dto.TagResponse {
	if entity == nil {
		return nil
	}

	return &dto.TagResponse{
		ID:          entity.ID,
		Name:        entity.Name,
		Slug:        entity.Slug,
		Description: entity.Description,
		Color:       entity.Color,
		Icon:        entity.Icon,
		IsActive:    entity.IsActive,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func ToTagPaginationDTO(entity *entity.PaginatedTagEntity) *dto.PaginatedTagDTO {

	tagResponses := make([]dto.TagResponse, len(entity.Items))
	for i, tag := range entity.Items {
		tagResponses[i] = *ToTagDTO(&tag)
	}

	// Tính total pages
	totalPages := 0
	if entity.PageSize > 0 {
		totalPages = (entity.TotalItems + entity.PageSize - 1) / entity.PageSize
	}

	return &dto.PaginatedTagDTO{
		Items:      tagResponses,
		TotalItems: entity.TotalItems,
		TotalPages: totalPages,
		PageNumber: entity.PageNumber,
		PageSize:   entity.PageSize,
	}
}
