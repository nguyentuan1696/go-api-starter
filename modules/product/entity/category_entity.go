package entity

import (
	"go-api-starter/core/entity"
)

// Category represents a product category entity
// Đại diện cho danh mục sản phẩm
type Category struct {

	// Name is the display name of the category
	// Tên hiển thị của danh mục
	Name string `db:"name"`

	// Slug is the URL-friendly version of the category name
	// Slug là phiên bản thân thiện URL của tên danh mục
	Slug string `db:"slug"`

	// Description provides detailed information about the category
	// Mô tả cung cấp thông tin chi tiết về danh mục
	Description string `db:"description"`

	// Thumbnail is the category's thumbnail image URL
	// Thumbnail là URL hình ảnh thu nhỏ của danh mục
	Thumbnail string `db:"thumbnail"`

	// ParentID is the ID of the parent category (for hierarchical structure)
	// ParentID là ID của danh mục cha (cho cấu trúc phân cấp)
	ParentID *string `db:"parent_id"`

	// ParentName is the name of the parent category
	// ParentName là tên của danh mục cha
	ParentName *string `db:"parent_name"`

	// SortOrder determines the display order in lists
	// Thứ tự sắp xếp xác định thứ tự hiển thị trong danh sách
	SortOrder int `db:"sort_order"`

	// IsActive indicates whether this category is currently active
	// Cho biết danh mục này có đang hoạt động không
	IsActive bool `db:"is_active"`

	entity.BaseEntity
}

type PaginatedCategoryResponse = entity.Pagination[Category]
