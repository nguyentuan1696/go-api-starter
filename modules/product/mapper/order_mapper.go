package mapper

import (
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/entity"
)

func ToOrderItemEntity(req *dto.OrderItemRequest) *entity.OrderItem {
	return &entity.OrderItem{
		ProductID:  req.ProductID,
		Name:       req.ProductName,
		UnitPrice:  req.OriginalPrice,
		Quantity:   req.Quantity,
		TotalPrice: req.Total,
	}
}

func ToOrderDTO(entity *entity.Order) *dto.OrderResponse {
	return &dto.OrderResponse{
		ID:                     entity.ID,
		OrderNumber:            entity.OrderNumber,
		CustomerID:             entity.CustomerID,
		CustomerEmail:          entity.CustomerEmail,
		CustomerPhone:          entity.CustomerPhone,
		CustomerName:           entity.CustomerName,
		ShippingRecipientName:  entity.ShippingRecipientName,
		ShippingRecipientPhone: entity.ShippingRecipientPhone,
		ShippingAddress:        entity.ShippingAddress,
		ShippingWardName:       entity.ShippingWardName,
		ShippingDistrictName:   entity.ShippingDistrictName,
		ShippingProvinceName:   entity.ShippingProvinceName,
		OrderState:             entity.OrderState,
		PaymentStatus:          entity.PaymentStatus,
		Subtotal:               entity.Subtotal,
		ShippingCost:           entity.ShippingCost,
		TaxAmount:              entity.TaxAmount,
		DiscountAmount:         entity.DiscountAmount,
		TotalAmount:            entity.TotalAmount,
		ShippingMethodID:       entity.ShippingMethodID,
		PaymentMethodID:        entity.PaymentMethodID,
		CouponCode:             entity.CouponCode,
		Notes:                  entity.Notes,
		AdminNotes:             entity.AdminNotes,
		CouponID:               entity.CouponID,
	}
}

func ToOrderEntity(req *dto.PlaceOrderRequest) *entity.Order {
	return &entity.Order{
		OrderNumber:            req.OrderNumber,
		CustomerID:             req.CustomerID,
		CustomerEmail:          req.CustomerEmail,
		CustomerPhone:          req.CustomerPhone,
		CustomerName:           req.CustomerName,
		ShippingRecipientName:  req.ShippingRecipientName,
		ShippingRecipientPhone: req.ShippingRecipientPhone,
		ShippingAddress:        req.ShippingAddress,
		ShippingWardName:       req.ShippingWardName,
		ShippingDistrictName:   req.ShippingDistrictName,
		ShippingProvinceName:   req.ShippingProvinceName,
		OrderState:             req.OrderState,
		PaymentStatus:          req.PaymentStatus,
		Subtotal:               req.Subtotal,
		ShippingCost:           req.ShippingCost,
		TaxAmount:              req.TaxAmount,
		DiscountAmount:         req.DiscountAmount,
		TotalAmount:            req.TotalAmount,
		ShippingMethodID:       req.ShippingMethodID,
		PaymentMethodID:        req.PaymentMethodID,
		CouponCode:             req.CouponCode,
		Notes:                  req.Notes,
		AdminNotes:             req.AdminNotes,
	}
}

func ToOrderPaginationDTO(entity *entity.PaginatedOrderEntity) *dto.PaginatedOrderDTO {

	orderResponses := make([]dto.OrderResponse, len(entity.Items))
	for i, order := range entity.Items {
		orderResponses[i] = *ToOrderDTO(&order)
	}

	// Tính total pages
	totalPages := 0
	if entity.PageSize > 0 {
		totalPages = (entity.TotalItems + entity.PageSize - 1) / entity.PageSize
	}

	return &dto.PaginatedOrderDTO{
		Items:      orderResponses,
		TotalItems: entity.TotalItems,
		TotalPages: totalPages,
		PageNumber: entity.PageNumber,
		PageSize:   entity.PageSize,
	}
}

func ToOrderDetailWithItemsDTO(entity *entity.OrderDetailWithItems) *dto.OrderDetailWithItemsDTO {
	return &dto.OrderDetailWithItemsDTO{
		OrderID:                entity.OrderID,
		OrderNumber:            entity.OrderNumber,
		CustomerID:             entity.CustomerID,
		CustomerName:           entity.CustomerName,
		CustomerEmail:          entity.CustomerEmail,
		CustomerPhone:          entity.CustomerPhone,
		ShippingRecipientName:  entity.ShippingRecipientName,
		ShippingRecipientPhone: entity.ShippingRecipientPhone,
		ShippingAddress:        entity.ShippingAddress,
		ShippingWardName:       entity.ShippingWardName,
		ShippingDistrictName:   entity.ShippingDistrictName,
		ShippingProvinceName:   entity.ShippingProvinceName,
		PaymentMethodID:        entity.PaymentMethodID,
		PaymentMethodName:      entity.PaymentMethodName,
		ShippingMethodID:       entity.ShippingMethodID,
		ShippingMethodName:     entity.ShippingMethodName,
		Notes:                  entity.Notes,
		AdminNotes:             entity.AdminNotes,
		OrderState:             entity.OrderState,
		PaymentStatus:          entity.PaymentStatus,
		Subtotal:               entity.Subtotal,
		ShippingCost:           entity.ShippingCost,
		TaxAmount:              entity.TaxAmount,
		DiscountAmount:         entity.DiscountAmount,
		TotalAmount:            entity.TotalAmount,
		ConfirmedAt:            entity.ConfirmedAt,
		ShippedAt:              entity.ShippedAt,
		DeliveredAt:            entity.DeliveredAt,
		CancelledAt:            entity.CancelledAt,
		Items:                  entity.Items,
	}
}
