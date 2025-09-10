package mapper

import (
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/entity"
)

func ToCategoryEntity(req *dto.CategoryRequest) *entity.Category {
	return &entity.Category{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		Thumbnail:   req.Thumbnail,
		ParentID:    req.ParentID,
		SortOrder:   req.SortOrder,
		IsActive:    req.IsActive,
	}
}

func ToCategoryResponse(entity *entity.Category) *dto.CategoryResponse {

	response := &dto.CategoryResponse{
		ID:          entity.ID,
		Name:        entity.Name,
		Slug:        entity.Slug,
		Description: entity.Description,
		Thumbnail:   entity.Thumbnail,
		ParentID:    entity.ParentID,
		ParentName:  entity.ParentName,
		SortOrder:   entity.SortOrder,
		IsActive:    entity.IsActive,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}

	return response
}

func ToCategoryPaginationResponse(entity *entity.PaginatedCategoryResponse) *dto.PaginatedCategoryResponse {
	// Convert từng category entity sang category response
	categoryResponses := make([]dto.CategoryResponse, len(entity.Items))
	for i, category := range entity.Items {
		categoryResponses[i] = *ToCategoryResponse(&category)
	}

	// Tính total pages
	totalPages := 0
	if entity.PageSize > 0 {
		totalPages = (entity.TotalItems + entity.PageSize - 1) / entity.PageSize
	}

	return &dto.PaginatedCategoryResponse{
		Items:      categoryResponses,
		TotalItems: entity.TotalItems,
		TotalPages: totalPages,
		PageNumber: entity.PageNumber,
		PageSize:   entity.PageSize,
	}
}
