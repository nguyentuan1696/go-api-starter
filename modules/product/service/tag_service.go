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

func (s *ProductService) PrivateCreateTag(ctx context.Context, req *dto.TagRequest) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	tag := mapper.ToTagEntity(req)
	tag.Slug = utils.GenerateSlugWithName(tag.Name)

	err := s.repo.PrivateCreateTag(ctx, tag)
	if err != nil {
		logger.Error("ProductService:PrivateCreateTag:Error: ", err)
		return errors.NewAppError(errors.ErrCreateFailed, "create tag failed", err)
	}

	return nil
}

func (s *ProductService) PrivateGetTags(ctx context.Context, params params.QueryParams) (*dto.PaginatedTagDTO, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	entity, err := s.repo.PrivateGetTags(ctx, params)
	if err != nil {
		logger.Error("ProductService:GetTags:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get tags failed", err)
	}

	return mapper.ToTagPaginationDTO(entity), nil
}

func (s *ProductService) PrivateGetTagById(ctx context.Context, id uuid.UUID) (*dto.TagResponse, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	entity, err := s.repo.PrivateGetTagById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("ProductService:PrivateGetTagById:TagNotFound: ", err)
			return nil, errors.NewAppError(errors.ErrNotFound, "tag not found", err)
		}
		logger.Error("ProductService:GetTagById:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get tag failed", err)
	}

	return mapper.ToTagDTO(entity), nil
}

func (s *ProductService) PrivateUpdateTag(ctx context.Context, id uuid.UUID, req *dto.TagRequest) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	tag := mapper.ToTagEntity(req)

	err := s.repo.PrivateUpdateTag(ctx, tag, id)
	if err != nil {
		logger.Error("ProductService:UpdateTag:Error: ", err)
		return errors.NewAppError(errors.ErrUpdateFailed, "update tag failed", err)
	}

	return nil
}

func (s *ProductService) PrivateDeleteTag(ctx context.Context, id uuid.UUID) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	err := s.repo.PrivateDeleteTag(ctx, id)
	if err != nil {
		logger.Error("ProductService:DeleteTag:Error: ", err)
		return errors.NewAppError(errors.ErrDeleteFailed, err.Error(), nil)
	}

	return nil
}
