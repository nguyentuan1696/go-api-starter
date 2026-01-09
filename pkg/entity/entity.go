package entity

import (
	"time"

	"github.com/google/uuid"
)

type Pagination[T any] struct {
	Items      []T `json:"items"`
	TotalItems int `json:"total_items"`
	PageNumber int `json:"page_number"`
	PageSize   int `json:"page_size"`
}

type BaseEntity struct {

	// ID is the unique identifier for the record
	ID uuid.UUID `db:"id"`
	// CreatedAt is the timestamp when the record was created
	CreatedAt time.Time `db:"created_at"`

	// UpdatedAt is the timestamp when the record was last updated
	UpdatedAt time.Time `db:"updated_at"`
}
