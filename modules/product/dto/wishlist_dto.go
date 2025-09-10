package dto

import (
	"time"
	"go-api-starter/core/dto"

	"github.com/google/uuid"
)

type WishlistRequest struct {
	CustomerID uuid.UUID `json:"user_id"`
	ProductID  string    `json:"product_id"`
}

type WishlistResponse struct {
	ID         uuid.UUID `json:"id"`
	CustomerID uuid.UUID `json:"user_id"`
	ProductID  string    `json:"product_id"`
	AddedAt    time.Time `json:"added_at"`
}

type PaginatedWishListDTO = dto.Pagination[WishlistResponse]
