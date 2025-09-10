package entity

import (
	"go-api-starter/core/entity"
)

// Brand represents a product brand entity
// Đại diện cho thương hiệu sản phẩm
type Brand struct {

	// Name is the display name of the brand
	// Tên hiển thị của thương hiệu
	Name string `db:"name"`

	// Slug is the URL-friendly version of the brand name
	// Slug là phiên bản thân thiện URL của tên thương hiệu
	Slug string `db:"slug"`

	// Description provides detailed information about the brand
	// Mô tả cung cấp thông tin chi tiết về thương hiệu
	Description string `db:"description"`

	// Logo is the brand's logo image URL
	// Logo là URL hình ảnh logo của thương hiệu
	Logo string `db:"logo"`

	// Website is the brand's official website URL
	// Website là URL trang web chính thức của thương hiệu
	Website string `db:"website"`

	// Country is the country where the brand originates from
	// Quốc gia nơi thương hiệu có nguồn gốc
	Country string `db:"country"`

	// FoundedYear is the year when the brand was founded
	// Năm thành lập thương hiệu
	FoundedYear int `db:"founded_year"`

	// IsActive indicates whether this brand is currently active
	// Cho biết thương hiệu này có đang hoạt động không
	IsActive bool `db:"is_active"`

	// IsFeatured indicates whether this brand is featured
	// Cho biết thương hiệu này có được nổi bật không
	IsFeatured bool `db:"is_featured"`

	// SortOrder determines the display order in lists
	// Thứ tự sắp xếp xác định thứ tự hiển thị trong danh sách
	SortOrder int `db:"sort_order"`

	entity.BaseEntity
}

type PaginatedBrandResponse = entity.Pagination[Brand]
