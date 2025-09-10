package mapper

import (
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/entity"
)

func ToBenefitEntity(request *dto.BenefitRequest) entity.Benefit {

	if request == nil {
		return entity.Benefit{}
	}

	return entity.Benefit{
		Name:               request.Name,
		Slug:               request.Slug,
		Description:        request.Description,
		Category:           request.Category,
		TargetArea:         request.TargetArea,
		EffectivenessLevel: request.EffectivenessLevel,
		TimeToSeeResults:   request.TimeToSeeResults,
		Color:              request.Color,
		Icon:               request.Icon,
		IsActive:           request.IsActive,
		IsFeatured:         request.IsFeatured,
		SortOrder:          request.SortOrder,
	}
}

func ToBenefitDTO(entity *entity.Benefit) *dto.BenefitResponse {

	if entity == nil {
		return &dto.BenefitResponse{}
	}

	return &dto.BenefitResponse{
		ID:                 entity.ID,
		Name:               entity.Name,
		Slug:               entity.Slug,
		Description:        entity.Description,
		Category:           entity.Category,
		TargetArea:         entity.TargetArea,
		EffectivenessLevel: entity.EffectivenessLevel,
		TimeToSeeResults:   entity.TimeToSeeResults,
		Color:              entity.Color,
		Icon:               entity.Icon,
		IsActive:           entity.IsActive,
		IsFeatured:         entity.IsFeatured,
		SortOrder:          entity.SortOrder,
		CreatedAt:          entity.CreatedAt,
		UpdatedAt:          entity.UpdatedAt,
	}
}

func ToPaginatedBenefitDTO(entity *entity.PaginatedBenefitEntity) *dto.PaginatedBenefitDTO {

	benefitResponses := make([]dto.BenefitResponse, len(entity.Items))
	for i, benefit := range entity.Items {
		benefitResponses[i] = *ToBenefitDTO(&benefit)
	}

	// Tính total pages
	totalPages := 0
	if entity.PageSize > 0 {
		totalPages = (entity.TotalItems + entity.PageSize - 1) / entity.PageSize
	}

	return &dto.PaginatedBenefitDTO{
		Items:      benefitResponses,
		TotalItems: entity.TotalItems,
		TotalPages: totalPages,
		PageNumber: entity.PageNumber,
		PageSize:   entity.PageSize,
	}
}
