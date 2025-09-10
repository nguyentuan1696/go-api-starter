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

func (s *ProductService) PublicGetProducts(ctx context.Context, params params.QueryParams) (*dto.PaginatedProductDTO, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	entity, err := s.repo.PrivateGetProducts(ctx, params)
	if err != nil {
		logger.Error("ProductService:PublicGetProducts:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get products failed", err)
	}

	return mapper.ToProductPaginationDTO(entity), nil
}

func (s *ProductService) PublicGetProductDetailWithFields(ctx context.Context, id uuid.UUID, field []string) (*dto.ProductDetailResponse, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	entity, err := s.repo.PublicGetProductDetail(ctx, id, field)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("ProductService:PublicGetProductDetail:ProductNotFound: ", err)
			return nil, errors.NewAppError(errors.ErrNotFound, "product not found", err)
		}
		logger.Error("ProductService:PublicGetProductDetail:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get product detail failed", err)
	}

	return mapper.ToProductDetailDTO(entity), nil
}

func (s *ProductService) PublicGetProductDetail(ctx context.Context, slug string) (*dto.ProductDetailResponse, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	entity, err := s.repo.PublicGetProductDetailBySlug(ctx, slug, nil)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("ProductService:PublicGetProductDetail:ProductNotFound: ", err)
			return nil, errors.NewAppError(errors.ErrNotFound, "product not found", err)
		}
		logger.Error("ProductService:PublicGetProductDetail:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get product detail failed", err)
	}

	return mapper.ToProductDetailDTO(entity), nil
}

func (s *ProductService) PrivateCreateProduct(ctx context.Context, req *dto.ProductRequest) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	product := mapper.ToProductEntity(req)
	product.Slug = utils.GenerateSlugWithName(product.Name)

	err := s.repo.PrivateCreateProduct(ctx, product)
	if err != nil {
		logger.Error("ProductService:PrivateCreateProduct:Error: ", err)
		return errors.NewAppError(errors.ErrCreateFailed, "create product failed", err)
	}

	return nil
}

func (s *ProductService) PrivateGetProducts(ctx context.Context, params params.QueryParams) (*dto.PaginatedProductDTO, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	entity, err := s.repo.PrivateGetProducts(ctx, params)
	if err != nil {
		logger.Error("ProductService:PrivateGetProducts:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get products failed", err)
	}

	return mapper.ToProductPaginationDTO(entity), nil
}

func (s *ProductService) PrivateGetProductById(ctx context.Context, id uuid.UUID) (*dto.ProductResponse, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	entity, err := s.repo.PrivateGetProductById(ctx, id)
	if err != nil {
		logger.Error("ProductService:PrivateGetProductById:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get product failed", err)
	}

	return mapper.ToProductDTO(entity), nil
}

func (s *ProductService) PrivateUpdateProduct(ctx context.Context, id uuid.UUID, req *dto.ProductRequest) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	product := mapper.ToProductEntity(req)

	err := s.repo.PrivateUpdateProduct(ctx, product, id)
	if err != nil {
		logger.Error("ProductService:PrivateUpdateProduct:Error: ", err)
		return errors.NewAppError(errors.ErrUpdateFailed, "update product failed", err)
	}

	return nil
}

func (s *ProductService) PrivateDeleteProduct(ctx context.Context, id uuid.UUID) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	err := s.repo.PrivateDeleteProduct(ctx, id)
	if err != nil {
		logger.Error("ProductService:PrivateDeleteProduct:Error: ", err)
		return errors.NewAppError(errors.ErrDeleteFailed, "delete product failed", err)
	}

	return nil
}
