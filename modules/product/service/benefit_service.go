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

func (s *ProductService) PrivateCreateBenefit(ctx context.Context, req *dto.BenefitRequest) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	benefit := mapper.ToBenefitEntity(req)
	benefit.Slug = utils.GenerateSlugWithName(benefit.Name)

	err := s.repo.PrivateCreateBenefit(ctx, &benefit)
	if err != nil {
		logger.Error("ProductService:PrivateCreateBenefit:Error: ", err)
		return errors.NewAppError(errors.ErrCreateFailed, "create benefit failed", err)
	}

	return nil
}

func (s *ProductService) PrivateGetBenefits(ctx context.Context, params params.QueryParams) (*dto.PaginatedBenefitDTO, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	entity, err := s.repo.PrivateGetBenefits(ctx, params)
	if err != nil {
		logger.Error("ProductService:GetBenefits:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get benefits failed", err)
	}

	return mapper.ToPaginatedBenefitDTO(entity), nil
}

func (s *ProductService) PrivateGetBenefitById(ctx context.Context, id uuid.UUID) (*dto.BenefitResponse, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	entity, err := s.repo.PrivateGetBenefitById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("ProductService:PrivateGetBenefitById:BenefitNotFound: ", err)
			return nil, errors.NewAppError(errors.ErrNotFound, "benefit not found", err)
		}
		logger.Error("ProductService:GetBenefitById:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get benefit failed", err)
	}

	return mapper.ToBenefitDTO(entity), nil
}

func (s *ProductService) PrivateUpdateBenefit(ctx context.Context, id uuid.UUID, req *dto.BenefitRequest) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	benefit := mapper.ToBenefitEntity(req)
	err := s.repo.PrivateUpdateBenefit(ctx, &benefit, id)
	if err != nil {
		logger.Error("ProductService:UpdateBenefit:Error: ", err)
		return errors.NewAppError(errors.ErrUpdateFailed, "update benefit failed", err)
	}

	return nil
}

func (s *ProductService) PrivateDeleteBenefit(ctx context.Context, id uuid.UUID) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	err := s.repo.PrivateDeleteBenefit(ctx, id)
	if err != nil {
		logger.Error("ProductService:DeleteBenefit:Error: ", err)
		return errors.NewAppError(errors.ErrDeleteFailed, "delete benefit failed", err)
	}

	return nil
}
