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

func (s *ProductService) PrivateCreateIngredient(ctx context.Context, req *dto.IngredientRequest) error {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	ingredient := mapper.ToIngredientEntity(req)
	ingredient.Slug = utils.GenerateSlugWithName(req.Name)

	err := s.repo.PrivateCreateIngredient(ctx, ingredient)
	if err != nil {
		logger.Error("ProductService:PrivateCreateIngredient:Error: ", err)
		return errors.NewAppError(errors.ErrCreateFailed, "create ingredient failed", err)
	}

	return nil
}

func (s *ProductService) PrivateGetIngredients(ctx context.Context, params params.QueryParams) (*dto.PaginatedIngredientDTO, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	ingredients, err := s.repo.PrivateGetIngredients(ctx, params)
	if err != nil {
		logger.Error("ProductService:PrivateGetIngredients:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get ingredients failed", err)
	}

	return mapper.ToIngredientPaginationDTO(ingredients), nil
}

func (s *ProductService) PrivateGetIngredientById(ctx context.Context, id uuid.UUID) (*dto.IngredientResponse, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	ingredient, err := s.repo.PrivateGetIngredientById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("ProductService:PrivateGetIngredientById:IngredientNotFound: ", err)
			return nil, errors.NewAppError(errors.ErrNotFound, "ingredient not found", err)
		}
		logger.Error("ProductService:PrivateGetIngredientById:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get ingredient failed", err)
	}

	return mapper.ToIngredientDTO(ingredient), nil
}

func (s *ProductService) PrivateUpdateIngredient(ctx context.Context, id uuid.UUID, req *dto.IngredientRequest) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	ingredient := mapper.ToIngredientEntity(req)

	err := s.repo.PrivateUpdateIngredient(ctx, ingredient, id)
	if err != nil {
		logger.Error("ProductService:PrivateUpdateIngredient:Error: ", err)
		return errors.NewAppError(errors.ErrUpdateFailed, "update ingredient failed", err)
	}

	return nil
}

func (s *ProductService) PrivateDeleteIngredient(ctx context.Context, id uuid.UUID) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	err := s.repo.PrivateDeleteIngredient(ctx, id)
	if err != nil {
		logger.Error("ProductService:PrivateDeleteIngredient:Error: ", err)
		return errors.NewAppError(errors.ErrDeleteFailed, "delete ingredient failed", err)
	}

	return nil
}
