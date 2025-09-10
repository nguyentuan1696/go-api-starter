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

func (s *ProductService) PrivateCreateSkinType(ctx context.Context, req *dto.SkinTypeRequest) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	skinType := mapper.ToSkinTypeEntity(req)
	skinType.Slug = utils.GenerateSlugWithName(skinType.Name)

	err := s.repo.PrivateCreateSkinType(ctx, skinType)
	if err != nil {
		logger.Error("ProductService:PrivateCreateSkinType:Error: ", err)
		return errors.NewAppError(errors.ErrCreateFailed, "create skin type failed", err)
	}

	return nil
}

func (s *ProductService) PrivateGetSkinTypes(ctx context.Context, params params.QueryParams) (*dto.PaginatedSkinTypeDTO, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	entity, err := s.repo.PrivateGetSkinTypes(ctx, params)
	if err != nil {
		logger.Error("ProductService:GetSkinTypes:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get skin types failed", err)
	}

	return mapper.ToSkinTypePaginationDTO(entity), nil
}

func (s *ProductService) PrivateGetSkinTypeById(ctx context.Context, id uuid.UUID) (*dto.SkinTypeResponse, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	entity, err := s.repo.PrivateGetSkinTypeById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("ProductService:PrivateGetSkinTypeById:SkinTypeNotFound: ", err)
			return nil, errors.NewAppError(errors.ErrNotFound, "skin type not found", err)
		}
		logger.Error("ProductService:GetSkinTypeById:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get skin type failed", err)
	}

	return mapper.ToSkinTypeDTO(entity), nil
}

func (s *ProductService) PrivateUpdateSkinType(ctx context.Context, id uuid.UUID, req *dto.SkinTypeRequest) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	skinType := mapper.ToSkinTypeEntity(req)

	err := s.repo.PrivateUpdateSkinType(ctx, skinType, id)
	if err != nil {
		logger.Error("ProductService:UpdateSkinType:Error: ", err)
		return errors.NewAppError(errors.ErrUpdateFailed, "update skin type failed", err)
	}

	return nil
}

func (s *ProductService) PrivateDeleteSkinType(ctx context.Context, id uuid.UUID) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	err := s.repo.PrivateDeleteSkinType(ctx, id)
	if err != nil {
		logger.Error("ProductService:DeleteSkinType:Error: ", err)
		return errors.NewAppError(errors.ErrDeleteFailed, "delete skin type failed", err)
	}

	return nil
}
