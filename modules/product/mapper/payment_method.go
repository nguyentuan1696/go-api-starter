package mapper

import (
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/entity"
)

func ToPaymentMethodDTO(entity *entity.PaymentMethod) *dto.PaymentMethodResponse {
	return &dto.PaymentMethodResponse{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Provider:    entity.Provider,
		Type:        entity.Type,
		IsActive:    entity.IsActive,
		SortOrder:   entity.SortOrder,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func ToPaymentMethodEntity(dto *dto.PaymentMethodRequest) *entity.PaymentMethod {
	return &entity.PaymentMethod{
		Name:        dto.Name,
		Description: dto.Description,
		Provider:    dto.Provider,
		Type:        dto.Type,
		IsActive:    dto.IsActive,
		SortOrder:   dto.SortOrder,
	}
}

func ToPaymentMethodPaginationDTO(entity *entity.PaginatedPaymentMethodEntity) *dto.PaginatedPaymentMethodDTO {

	paymentMethodResponses := make([]dto.PaymentMethodResponse, len(entity.Items))
	for i, paymentMethod := range entity.Items {
		paymentMethodResponses[i] = *ToPaymentMethodDTO(&paymentMethod)
	}

	// Tính total pages
	totalPages := 0
	if entity.PageSize > 0 {
		totalPages = (entity.TotalItems + entity.PageSize - 1) / entity.PageSize
	}

	return &dto.PaginatedPaymentMethodDTO{
		Items:      paymentMethodResponses,
		TotalItems: entity.TotalItems,
		TotalPages: totalPages,
		PageNumber: entity.PageNumber,
		PageSize:   entity.PageSize,
	}
}
