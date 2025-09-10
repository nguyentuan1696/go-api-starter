package service

import (
	"context"
	"database/sql"
	"go-api-starter/core/constants"
	"go-api-starter/core/errors"
	"go-api-starter/core/logger"
	"go-api-starter/core/params"
	"go-api-starter/core/utils"
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/entity"
	"go-api-starter/modules/product/mapper"

	"github.com/google/uuid"
)

func (s *ProductService) PublicGetOrderDetailWithItems(ctx context.Context, orderID uuid.UUID) (*dto.OrderDetailWithItemsDTO, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultTimeout)
	defer cancel()

	result, err := s.repo.GetOrderDetailWithItems(ctx, orderID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewAppError(errors.ErrNotFound, "Order not found", nil)
		}
		logger.Error("ProductService:PublicGetOrderDetailWithItems", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "Order error", nil)
	}

	return mapper.ToOrderDetailWithItemsDTO(result), nil
}

func (s *ProductService) PrivateGetOrderDetailWithItems(ctx context.Context, orderID uuid.UUID) (*dto.OrderDetailWithItemsDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultTimeout)
	defer cancel()

	result, err := s.repo.GetOrderDetailWithItems(ctx, orderID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewAppError(errors.ErrNotFound, "Order not found", nil)
		}
		logger.Error("ProductService:PublicGetOrderDetailWithItems", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "Order error", nil)
	}

	return mapper.ToOrderDetailWithItemsDTO(result), nil
}

func (s *ProductService) PublicPlaceOrder(ctx context.Context, req *dto.PlaceOrderRequest) (*dto.PlaceOrderResponse, *errors.AppError) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 1. Kiểm tra shipping method
	shippingMethodInfo, errShippingMethodInfo := s.PublicGetShippingMethodDetail(ctx, req.ShippingMethodID)
	if errShippingMethodInfo != nil {
		if errShippingMethodInfo == sql.ErrNoRows {
			return nil, errors.NewAppError(errors.ErrNotFound, "Shipping method not found", nil)
		}
		logger.Error("ProductService:PublicPlaceOrder:PublicGetShippingMethodDetail", errShippingMethodInfo)
		return nil, errors.NewAppError(errors.ErrGetFailed, "Shipping method error", nil)
	}

	// 2. Kiểm tra payment method
	if _, errPaymentInfo := s.PublicGetPaymentMethodDetail(ctx, req.PaymentMethodID); errPaymentInfo != nil {
		if errPaymentInfo == sql.ErrNoRows {
			return nil, errors.NewAppError(errors.ErrNotFound, "Payment method not found", nil)
		}
		logger.Error("ProductService:PublicPlaceOrder:PublicGetPaymentMethodDetail", errPaymentInfo)
		return nil, errors.NewAppError(errors.ErrGetFailed, "Payment method error", nil)
	}

	// 3. Tính tổng tiền hàng
	var subTotal float64
	for _, item := range req.OrderItems {
		productDetail, errProductDetail := s.PublicGetProductDetailWithFields(ctx, item.ProductID, []string{"id", "price"})
		if errProductDetail != nil {
			if errProductDetail == sql.ErrNoRows {
				return nil, errors.NewAppError(errors.ErrNotFound, "Product not found", nil)
			}
			logger.Error("ProductService:PublicPlaceOrder:GetProductDetailWithFields", errProductDetail)
			return nil, errors.NewAppError(errors.ErrGetFailed, "Product error", nil)
		}

		subTotal += productDetail.Price * float64(item.Quantity)
	}

	// 4. Tính phí vận chuyển
	shippingCost := shippingMethodInfo.BaseCost
	if shippingMethodInfo.FreeShippingThreshold > 0 && subTotal >= shippingMethodInfo.FreeShippingThreshold {
		shippingCost = 0
	}

	// 5. Tạo và lưu đơn hàng
	orderEntity := mapper.ToOrderEntity(req)
	orderEntity.TotalAmount = subTotal + shippingCost
	orderEntity.OrderState = constants.OrderStatePending
	orderEntity.PaymentStatus = constants.PaymentStatusPending
	orderEntity.OrderNumber = utils.GenerateOrderNumber()
	orderEntity.ShippingMethodName = shippingMethodInfo.Name
	
	createdOrder, errCreateOrder := s.repo.PublicCreateOrder(ctx, orderEntity)
	if errCreateOrder != nil {
		logger.Error("ProductService:PublicPlaceOrder:PublicCreateOrder", errCreateOrder)
		return nil, errors.NewAppError(errors.ErrCreateFailed, "Create order error", nil)
	}

	// 7. Tạo và lưu chi tiết đơn hàng
	var orderItemEntities []*entity.OrderItem
	for _, item := range req.OrderItems {
		e := mapper.ToOrderItemEntity(&item)
		e.OrderID = createdOrder.ID
		e.ProductID = item.ProductID
		e.UnitPrice = item.OriginalPrice
		e.TotalPrice = float64(item.Quantity) * item.OriginalPrice
		orderItemEntities = append(orderItemEntities, e)
	}
	if err := s.repo.PrivateCreateOrderItems(ctx, orderItemEntities); err != nil {
		logger.Error("ProductService:PublicPlaceOrder:PrivateCreateOrderItems", err)
		return nil, errors.NewAppError(errors.ErrCreateFailed, "Create order items error", err)
	}

	// 6. Tạo response
	return &dto.PlaceOrderResponse{
		OrderNumber:            createdOrder.OrderNumber,
		ShippingRecipientName:  createdOrder.ShippingRecipientName,
		ShippingRecipientPhone: createdOrder.ShippingRecipientPhone,
		ShippingAddress:        createdOrder.ShippingAddress,
		TotalAmount:            createdOrder.TotalAmount,
	}, nil
}

func (r *ProductService) PrivateGetOrders(ctx context.Context, params params.QueryParams) (*dto.PaginatedOrderDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultTimeout)
	defer cancel()

	orders, err := r.repo.PrivateGetOrders(ctx, params)
	if err != nil {
		logger.Error("ProductService:PrivateGetOrders:PrivateGetOrders", err)
		return nil, err
	}

	return mapper.ToOrderPaginationDTO(orders), nil
}

func (r *ProductService) PrivateGetOrderById(ctx context.Context, id uuid.UUID) (*dto.OrderResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultTimeout)
	defer cancel()

	order, err := r.repo.PrivateGetOrderById(ctx, id)
	if err != nil {
		logger.Error("ProductService:PrivateGetOrderById:PrivateGetOrderById", err)
		return nil, err
	}

	return mapper.ToOrderDTO(order), nil
}
