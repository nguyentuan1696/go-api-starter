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

func (s *ProductService) PrivateCreateCategory(ctx context.Context, req *dto.CategoryRequest) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	category := mapper.ToCategoryEntity(req)
	category.Slug = utils.GenerateSlugWithName(category.Name)

	err := s.repo.PrivateCreateCategory(ctx, category)
	if err != nil {
		return errors.NewAppError(errors.ErrCreateFailed, "create category failed", err)
	}
	return nil
}

func (s *ProductService) PrivateGetCategoryById(ctx context.Context, id uuid.UUID) (*dto.CategoryResponse, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	category, err := s.repo.PrivateGetCategoryById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("ProductService:PrivateGetCategoryById:CategoryNotFound: ", err)
			return nil, errors.NewAppError(errors.ErrNotFound, "category not found", err)
		}
		return nil, errors.NewAppError(errors.ErrGetFailed, "get category failed", err)
	}
	return mapper.ToCategoryResponse(category), nil
}

func (s *ProductService) PrivateGetCategories(ctx context.Context, params params.QueryParams) (*dto.PaginatedCategoryResponse, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	categories, err := s.repo.PrivateGetCategories(ctx, params)
	if err != nil {
		return nil, errors.NewAppError(errors.ErrGetFailed, "get categories failed", err)
	}
	return mapper.ToCategoryPaginationResponse(categories), nil
}

func (s *ProductService) PrivateUpdateCategory(ctx context.Context, req *dto.CategoryRequest, id uuid.UUID) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	category := mapper.ToCategoryEntity(req)

	err := s.repo.PrivateUpdateCategory(ctx, category, id)
	if err != nil {
		return errors.NewAppError(errors.ErrGetFailed, "get category failed", err)
	}

	return nil
}

func (s *ProductService) PrivateDeleteCategory(ctx context.Context, id uuid.UUID) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	err := s.repo.PrivateDeleteCategory(ctx, id)
	if err != nil {
		return errors.NewAppError(errors.ErrGetFailed, "get category failed", err)
	}

	return nil
}
