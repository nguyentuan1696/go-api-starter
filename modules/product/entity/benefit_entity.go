package entity

import (
	"go-api-starter/core/entity"
)

// Benefit represents a product benefit entity
// Đại diện cho lợi ích sản phẩm
type Benefit struct {

	// Name is the display name of the benefit
	// Tên hiển thị của lợi ích
	Name string `db:"name"`

	// Slug is the URL-friendly version of the name
	// Slug là phiên bản thân thiện URL của tên
	Slug string `db:"slug"`

	// Description provides detailed information about the benefit
	// Mô tả cung cấp thông tin chi tiết về lợi ích
	Description string `db:"description"`

	// Category groups benefits into different types
	// Danh mục nhóm các lợi ích thành các loại khác nhau
	Category string `db:"category"`

	// TargetArea specifies which body area this benefit targets
	// Vùng mục tiêu chỉ định vùng cơ thể mà lợi ích này nhắm đến
	TargetArea string `db:"target_area"`

	// EffectivenessLevel indicates how effective this benefit is
	// Mức độ hiệu quả cho biết lợi ích này hiệu quả như thế nào
	EffectivenessLevel string `db:"effectiveness_level"`

	// TimeToSeeResults indicates how long it takes to see results
	// Thời gian để thấy kết quả cho biết mất bao lâu để thấy hiệu quả
	TimeToSeeResults string `db:"time_to_see_results"`

	// Color is the theme color associated with this benefit
	// Màu sắc chủ đề được liên kết với lợi ích này
	Color string `db:"color"`

	// Icon is the icon identifier for this benefit
	// Biểu tượng là định danh biểu tượng cho lợi ích này
	Icon string `db:"icon"`

	// IsActive indicates whether this benefit is currently active
	// Cho biết lợi ích này có đang hoạt động không
	IsActive bool `db:"is_active"`

	// IsFeatured indicates whether this benefit is featured
	// Cho biết lợi ích này có được nổi bật không
	IsFeatured bool `db:"is_featured"`

	// SortOrder determines the display order in lists
	// Thứ tự sắp xếp xác định thứ tự hiển thị trong danh sách
	SortOrder int `db:"sort_order"`

	entity.BaseEntity
}

type PaginatedBenefitEntity = entity.Pagination[Benefit]
