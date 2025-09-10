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
	// ID là định danh duy nhất cho bản ghi
	ID uuid.UUID `db:"id"`
	// CreatedAt is the timestamp when the record was created
	// Thời gian tạo bản ghi
	CreatedAt time.Time `db:"created_at"`

	// UpdatedAt is the timestamp when the record was last updated
	// Thời gian cập nhật bản ghi lần cuối
	UpdatedAt time.Time `db:"updated_at"`
}
