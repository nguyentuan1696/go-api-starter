package entity

import (
	"time"
	"go-api-starter/core/entity"

	"github.com/google/uuid"
)

type Wishlist struct {
	ID         uuid.UUID `db:"id"`
	CustomerID uuid.UUID `db:"user_id"`
	ProductID  string    `db:"product_id"`
	AddedAt    time.Time `db:"added_at"`
}

type PaginatedWishlistEntity = entity.Pagination[Wishlist]
