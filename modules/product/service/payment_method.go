package service

import (
	"context"
	"database/sql"
	"go-api-starter/core/constants"
	"go-api-starter/core/errors"
	"go-api-starter/core/logger"
	"go-api-starter/core/params"
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/mapper"
)

func (s *ProductService) PublicGetPaymentMethodDetail(ctx context.Context, id int) (*dto.PaymentMethodResponse, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	entity, err := s.repo.PrivateGetPaymentMethodById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("ProductService:PrivateGetPaymentMethodById:PaymentMethodNotFound: ", err)
			return nil, errors.NewAppError(errors.ErrNotFound, "payment method not found", err)
		}
		logger.Error("ProductService:GetPaymentMethodById:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get payment method failed", err)
	}

	return mapper.ToPaymentMethodDTO(entity), nil
}

func (s *ProductService) PrivateCreatePaymentMethod(ctx context.Context, req *dto.PaymentMethodRequest) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	paymentMethod := mapper.ToPaymentMethodEntity(req)

	err := s.repo.PrivateCreatePaymentMethod(ctx, paymentMethod)
	if err != nil {
		logger.Error("ProductService:PrivateCreatePaymentMethod:Error: ", err)
		return errors.NewAppError(errors.ErrCreateFailed, "create payment method failed", err)
	}

	return nil
}

func (s *ProductService) PrivateGetPaymentMethods(ctx context.Context, params params.QueryParams) (*dto.PaginatedPaymentMethodDTO, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	entity, err := s.repo.PrivateGetPaymentMethods(ctx, params)
	if err != nil {
		logger.Error("ProductService:GetPaymentMethods:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get payment methods failed", err)
	}

	return mapper.ToPaymentMethodPaginationDTO(entity), nil
}

func (s *ProductService) PrivateGetPaymentMethodById(ctx context.Context, id int) (*dto.PaymentMethodResponse, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	entity, err := s.repo.PrivateGetPaymentMethodById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("ProductService:PrivateGetPaymentMethodById:PaymentMethodNotFound: ", err)
			return nil, errors.NewAppError(errors.ErrNotFound, "payment method not found", err)
		}
		logger.Error("ProductService:GetPaymentMethodById:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get payment method failed", err)
	}

	return mapper.ToPaymentMethodDTO(entity), nil
}

func (s *ProductService) PrivateUpdatePaymentMethod(ctx context.Context, id int, req *dto.PaymentMethodRequest) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	paymentMethod := mapper.ToPaymentMethodEntity(req)

	err := s.repo.PrivateUpdatePaymentMethod(ctx, paymentMethod, id)
	if err != nil {
		logger.Error("ProductService:UpdatePaymentMethod:Error: ", err)
		return errors.NewAppError(errors.ErrUpdateFailed, "update payment method failed", err)
	}

	return nil
}

func (s *ProductService) PrivateDeletePaymentMethod(ctx context.Context, id int) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	err := s.repo.PrivateDeletePaymentMethod(ctx, id)
	if err != nil {
		logger.Error("ProductService:DeletePaymentMethod:Error: ", err)
		return errors.NewAppError(errors.ErrDeleteFailed, err.Error(), nil)
	}

	return nil
}
