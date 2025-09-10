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

func (s *ProductService) PublicCreateWishlist(ctx context.Context, req *dto.WishlistRequest) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	wishlist := mapper.ToWishListEntity(req)

	err := s.repo.PublicCreateWishlist(ctx, wishlist)
	if err != nil {
		logger.Error("ProductService:PublicCreateWishlist:Error: ", err)
		return errors.NewAppError(errors.ErrCreateFailed, "create wishlist failed", err)
	}

	return nil
}

func (s *ProductService) PublicGetWishlists(ctx context.Context, params params.QueryParams) (*dto.PaginatedWishListDTO, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	entity, err := s.repo.PublicGetWishlists(ctx, params)
	if err != nil {
		logger.Error("ProductService:PublicGetWishlists:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get wishlists failed", err)
	}

	return mapper.ToPaginatedWishlistDTO(entity), nil
}

func (s *ProductService) PublicGetWishlistById(ctx context.Context, id string) (*dto.WishlistResponse, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	entity, err := s.repo.PublicGetWishlistById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("ProductService:PublicGetWishlistById:WishlistNotFound: ", err)
			return nil, errors.NewAppError(errors.ErrNotFound, "wishlist not found", err)
		}
		logger.Error("ProductService:PublicGetWishlistById:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get wishlist failed", err)
	}

	return mapper.ToWishListDTO(entity), nil
}

func (s *ProductService) PublicGetWishlistByCustomerAndProduct(ctx context.Context, customerID string, productID string) (*dto.WishlistResponse, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	entity, err := s.repo.PublicGetWishlistByCustomerAndProduct(ctx, customerID, productID)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("ProductService:PublicGetWishlistByCustomerAndProduct:WishlistNotFound: ", err)
			return nil, errors.NewAppError(errors.ErrNotFound, "wishlist not found", err)
		}
		logger.Error("ProductService:PublicGetWishlistByCustomerAndProduct:Error: ", err)
		return nil, errors.NewAppError(errors.ErrGetFailed, "get wishlist failed", err)
	}

	return mapper.ToWishListDTO(entity), nil
}

func (s *ProductService) PublicDeleteWishlist(ctx context.Context, id string) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	err := s.repo.PublicDeleteWishlist(ctx, id)
	if err != nil {
		logger.Error("ProductService:PublicDeleteWishlist:Error: ", err)
		return errors.NewAppError(errors.ErrDeleteFailed, err.Error(), nil)
	}

	return nil
}
