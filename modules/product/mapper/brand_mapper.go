package mapper

import (
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/entity"
)

func ToBrandEntity(req *dto.BrandRequest) *entity.Brand {
	return &entity.Brand{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		Logo:        req.Logo,
		Website:     req.Website,
		Country:     req.Country,
		FoundedYear: req.FoundedYear,
		IsActive:    req.IsActive,
		IsFeatured:  req.IsFeatured,
		SortOrder:   req.SortOrder,
	}
}

func ToBrandResponse(entity *entity.Brand) *dto.BrandResponse {
	return &dto.BrandResponse{
		ID:          entity.ID,
		Name:        entity.Name,
		Slug:        entity.Slug,
		Description: entity.Description,
		Logo:        entity.Logo,
		Website:     entity.Website,
		Country:     entity.Country,
		FoundedYear: entity.FoundedYear,
		IsActive:    entity.IsActive,
		IsFeatured:  entity.IsFeatured,
		SortOrder:   entity.SortOrder,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func ToBrandPaginationResponse(entity *entity.PaginatedBrandResponse) *dto.PaginatedBrandResponse {
	// Convert từng brand entity sang brand response
	brandResponses := make([]dto.BrandResponse, len(entity.Items))
	for i, brand := range entity.Items {
		brandResponses[i] = *ToBrandResponse(&brand)
	}

	// Tính total pages
	totalPages := 0
	if entity.PageSize > 0 {
		totalPages = (entity.TotalItems + entity.PageSize - 1) / entity.PageSize
	}

	return &dto.PaginatedBrandResponse{
		Items:      brandResponses,
		TotalItems: entity.TotalItems,
		TotalPages: totalPages,
		PageNumber: entity.PageNumber,
		PageSize:   entity.PageSize,
	}
}
