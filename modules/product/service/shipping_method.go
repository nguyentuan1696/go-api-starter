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

func (s *ProductService) PublicGetShippingMethodDetail(ctx context.Context, id int) (*dto.ShippingMethodResponse, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	entity, err := s.repo.PrivateGetShippingMethodById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("ProductService:PrivateGetShippingMethodById:ShippingMethodNotFound: ", err)
			return nil, errors.NewAppError(errors.ErrNotFound, "shipping method not found", err)
		}
		logger.Error("ProductService:GetShippingMethodById:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get shipping method failed", err)
	}

	return mapper.ToShippingMethodDTO(entity), nil
}

func (s *ProductService) PrivateCreateShippingMethod(ctx context.Context, req *dto.ShippingMethodRequest) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	shippingMethod := mapper.ToShippingMethodEntity(req)

	err := s.repo.PrivateCreateShippingMethod(ctx, shippingMethod)
	if err != nil {
		logger.Error("ProductService:PrivateCreateShippingMethod:Error: ", err)
		return errors.NewAppError(errors.ErrCreateFailed, "create shipping method failed", err)
	}

	return nil
}

func (s *ProductService) PrivateGetShippingMethods(ctx context.Context, params params.QueryParams) (*dto.PaginatedShippingMethodDTO, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	entity, err := s.repo.PrivateGetShippingMethods(ctx, params)
	if err != nil {
		logger.Error("ProductService:GetShippingMethods:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get shipping methods failed", err)
	}

	return mapper.ToShippingMethodPaginationDTO(entity), nil
}

func (s *ProductService) PrivateGetShippingMethodById(ctx context.Context, id int) (*dto.ShippingMethodResponse, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	entity, err := s.repo.PrivateGetShippingMethodById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("ProductService:PrivateGetShippingMethodById:ShippingMethodNotFound: ", err)
			return nil, errors.NewAppError(errors.ErrNotFound, "shipping method not found", err)
		}
		logger.Error("ProductService:GetShippingMethodById:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get shipping method failed", err)
	}

	return mapper.ToShippingMethodDTO(entity), nil
}

func (s *ProductService) PrivateUpdateShippingMethod(ctx context.Context, id int, req *dto.ShippingMethodRequest) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	shippingMethod := mapper.ToShippingMethodEntity(req)

	err := s.repo.PrivateUpdateShippingMethod(ctx, shippingMethod, id)
	if err != nil {
		logger.Error("ProductService:UpdateShippingMethod:Error: ", err)
		return errors.NewAppError(errors.ErrUpdateFailed, "update shipping method failed", err)
	}

	return nil
}

func (s *ProductService) PrivateDeleteShippingMethod(ctx context.Context, id int) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	err := s.repo.PrivateDeleteShippingMethod(ctx, id)
	if err != nil {
		logger.Error("ProductService:DeleteShippingMethod:Error: ", err)
		return errors.NewAppError(errors.ErrDeleteFailed, err.Error(), nil)
	}

	return nil
}
