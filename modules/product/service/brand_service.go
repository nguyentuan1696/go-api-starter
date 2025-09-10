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
	"go-api-starter/modules/product/mapper"

	"github.com/google/uuid"
)

func (s *ProductService) PrivateCreateBrand(ctx context.Context, req *dto.BrandRequest) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	brand := mapper.ToBrandEntity(req)
	brand.Slug = utils.GenerateSlugWithName(brand.Name)

	err := s.repo.PrivateCreateBrand(ctx, brand)
	if err != nil {
		logger.Error("ProductService:PrivateCreateBrand:Error: ", err)
		return errors.NewAppError(errors.ErrCreateFailed, "create brand failed", err)
	}

	return nil
}

func (s *ProductService) PrivateGetBrands(ctx context.Context, params params.QueryParams) (*dto.PaginatedBrandResponse, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	entity, err := s.repo.PrivateGetBrands(ctx, params)
	if err != nil {
		logger.Error("ProductService:GetBrands:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get brands failed", err)
	}

	return mapper.ToBrandPaginationResponse(entity), nil
}

func (s *ProductService) PrivateGetBrandById(ctx context.Context, id uuid.UUID) (*dto.BrandResponse, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	entity, err := s.repo.PrivateGetBrandById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("ProductService:PrivateGetBrandById:BrandNotFound: ", err)
			return nil, errors.NewAppError(errors.ErrNotFound, "brand not found", err)
		}
		logger.Error("ProductService:GetBrandById:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get brand failed", err)
	}

	return mapper.ToBrandResponse(entity), nil
}

func (s *ProductService) PrivateUpdateBrand(ctx context.Context, req *dto.BrandRequest, id uuid.UUID) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	brand := mapper.ToBrandEntity(req)

	err := s.repo.PrivateUpdateBrand(ctx, brand, id)
	if err != nil {
		logger.Error("ProductService:UpdateBrand:Error: ", err)
		return errors.NewAppError(errors.ErrUpdateFailed, "update brand failed", err)
	}

	return nil
}

func (s *ProductService) PrivateDeleteBrand(ctx context.Context, id uuid.UUID) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	err := s.repo.PrivateDeleteBrand(ctx, id)
	if err != nil {
		logger.Error("ProductService:DeleteBrand:Error: ", err)
		return errors.NewAppError(errors.ErrDeleteFailed, "delete brand failed", err)
	}

	return nil
}
