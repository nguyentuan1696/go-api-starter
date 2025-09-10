package mapper

import (
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/entity"
)

func ToWishListEntity(req *dto.WishlistRequest) *entity.Wishlist {
	return &entity.Wishlist{
		CustomerID: req.CustomerID,
		ProductID:  req.ProductID,
	}
}

func ToWishListDTO(wishlist *entity.Wishlist) *dto.WishlistResponse {
	return &dto.WishlistResponse{
		ID:         wishlist.ID,
		CustomerID: wishlist.CustomerID,
		ProductID:  wishlist.ProductID,
		AddedAt:    wishlist.AddedAt,
	}
}

func ToPaginatedWishlistDTO(entity *entity.PaginatedWishlistEntity) *dto.PaginatedWishListDTO {

	tagResponses := make([]dto.WishlistResponse, len(entity.Items))
	for i, tag := range entity.Items {
		tagResponses[i] = *ToWishListDTO(&tag)
	}

	// Tính total pages
	totalPages := 0
	if entity.PageSize > 0 {
		totalPages = (entity.TotalItems + entity.PageSize - 1) / entity.PageSize
	}

	return &dto.PaginatedWishListDTO{
		Items:      tagResponses,
		TotalItems: entity.TotalItems,
		TotalPages: totalPages,
		PageNumber: entity.PageNumber,
		PageSize:   entity.PageSize,
	}
}
