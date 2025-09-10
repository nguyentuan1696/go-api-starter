package service

import (
	"context"
	"go-api-starter/core/constants"
	"go-api-starter/core/errors"
	"go-api-starter/core/logger"
	"go-api-starter/core/params"
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/mapper"
)

func (s *ProductService) PublicGetProvinces(ctx context.Context, params params.QueryParams) (*dto.PaginatedProvinceDTO, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	provinces, err := s.repo.PublicGetProvinces(ctx, params)
	if err != nil {
		logger.Error("ProductService:PublicGetProvinces error: %v", err)
		return nil, errors.NewAppError(errors.ErrInternalServer, "internal server error", err)
	}

	return mapper.ToProvincePaginationDTO(provinces), nil
}

func (s *ProductService) PublicGetDistricts(ctx context.Context, params params.QueryParams) (*dto.PaginatedDistrictDTO, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	districts, err := s.repo.PublicGetDistricts(ctx, params)
	if err != nil {
		logger.Error("ProductService:PublicGetDistricts error: %v", err)
		return nil, errors.NewAppError(errors.ErrInternalServer, "internal server error", err)
	}

	return mapper.ToDistrictPaginationDTO(districts), nil
}

func (s *ProductService) PublicGetWards(ctx context.Context, params params.QueryParams) (*dto.PaginatedWardDTO, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	wards, err := s.repo.PublicGetWards(ctx, params)
	if err != nil {
		logger.Error("ProductService:PublicGetWards error: %v", err)
		return nil, errors.NewAppError(errors.ErrInternalServer, "internal server error", err)
	}

	return mapper.ToWardPaginationDTO(wards), nil
}
