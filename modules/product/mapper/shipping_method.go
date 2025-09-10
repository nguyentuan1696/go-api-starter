package mapper

import (
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/entity"
)

func ToShippingMethodEntity(req *dto.ShippingMethodRequest) *entity.ShippingMethod {
	return &entity.ShippingMethod{
		Name:                  req.Name,
		Description:           req.Description,
		Provider:              req.Provider,
		BaseCost:              req.BaseCost,
		CostPerKg:             req.CostPerKg,
		FreeShippingThreshold: req.FreeShippingThreshold,
		EstimatedDaysMin:      req.EstimatedDaysMin,
		EstimatedDaysMax:      req.EstimatedDaysMax,
		IsActive:              req.IsActive,
		SortOrder:             req.SortOrder,
	}
}

func ToShippingMethodDTO(entity *entity.ShippingMethod) *dto.ShippingMethodResponse {
	return &dto.ShippingMethodResponse{
		ID:                    entity.ID,
		Name:                  entity.Name,
		Description:           entity.Description,
		Provider:              entity.Provider,
		BaseCost:              entity.BaseCost,
		CostPerKg:             entity.CostPerKg,
		FreeShippingThreshold: entity.FreeShippingThreshold,
		EstimatedDaysMin:      entity.EstimatedDaysMin,
		EstimatedDaysMax:      entity.EstimatedDaysMax,
		IsActive:              entity.IsActive,
		SortOrder:             entity.SortOrder,
		CreatedAt:             entity.CreatedAt,
		UpdatedAt:             entity.UpdatedAt,
	}
}

func ToShippingMethodPaginationDTO(entity *entity.PaginatedShippingMethodEntity) *dto.PaginatedShippingMethodDTO {

	shippingMethodResponses := make([]dto.ShippingMethodResponse, len(entity.Items))
	for i, shippingMethod := range entity.Items {
		shippingMethodResponses[i] = *ToShippingMethodDTO(&shippingMethod)
	}

	// Tính total pages
	totalPages := 0
	if entity.PageSize > 0 {
		totalPages = (entity.TotalItems + entity.PageSize - 1) / entity.PageSize
	}

	return &dto.PaginatedShippingMethodDTO{
		Items:      shippingMethodResponses,
		TotalItems: entity.TotalItems,
		TotalPages: totalPages,
		PageNumber: entity.PageNumber,
		PageSize:   entity.PageSize,
	}
}
