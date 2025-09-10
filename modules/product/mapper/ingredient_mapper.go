package mapper

import (
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/entity"
)

func ToIngredientEntity(dto *dto.IngredientRequest) *entity.Ingredient {
	if dto == nil {
		return nil
	}
	return &entity.Ingredient{
		Name:         dto.Name,
		Slug:         dto.Slug,
		InciName:     dto.InciName,
		Description:  dto.Description,
		Origin:       dto.Origin,
		Function:     dto.Function,
		CasNumber:    dto.CasNumber,
		EwgScore:     dto.EwgScore,
		IsRestricted: dto.IsRestricted,
		IsBanned:     dto.IsBanned,
	}
}

func ToIngredientDTO(entity *entity.Ingredient) *dto.IngredientResponse {
	if entity == nil {
		return nil
	}
	return &dto.IngredientResponse{
		ID:           entity.ID,
		Name:         entity.Name,
		Slug:         entity.Slug,
		InciName:     entity.InciName,
		Description:  entity.Description,
		Origin:       entity.Origin,
		Function:     entity.Function,
		CasNumber:    entity.CasNumber,
		EwgScore:     entity.EwgScore,
		IsRestricted: entity.IsRestricted,
		IsBanned:     entity.IsBanned,
	}
}

func ToIngredientPaginationDTO(entity *entity.PaginatedIngredientEntity) *dto.PaginatedIngredientDTO {

	ingredientResponses := make([]dto.IngredientResponse, len(entity.Items))
	for i, ingredient := range entity.Items {
		ingredientResponses[i] = *ToIngredientDTO(&ingredient)
	}

	// Tính total pages
	totalPages := 0
	if entity.PageSize > 0 {
		totalPages = (entity.TotalItems + entity.PageSize - 1) / entity.PageSize
	}

	return &dto.PaginatedIngredientDTO{
		Items:      ingredientResponses,
		TotalItems: entity.TotalItems,
		TotalPages: totalPages,
		PageNumber: entity.PageNumber,
		PageSize:   entity.PageSize,
	}
}
