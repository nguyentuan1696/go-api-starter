package entity

import (
	"go-api-starter/core/entity"
)

// SkinType represents a skin type classification for cosmetic products
// Đại diện cho phân loại loại da dành cho sản phẩm mỹ phẩm
type SkinType struct {

	// Name is the display name of the skin type (e.g., "Oily", "Dry", "Combination")
	// Tên hiển thị của loại da (ví dụ: "Da dầu", "Da khô", "Da hỗn hợp")
	Name string `db:"name"`

	// Slug is the URL-friendly version of the skin type name
	// Slug là phiên bản thân thiện URL của tên loại da
	Slug string `db:"slug"`

	// Description is a brief description of the skin type
	// Mô tả ngắn gọn về loại da
	Description string `db:"description"`

	// Characteristics describes the specific traits and features of this skin type
	// Đặc điểm mô tả các tính chất và đặc trưng cụ thể của loại da này
	Characteristics string `db:"characteristics"`

	// CareTips provides skincare advice and recommendations for this skin type
	// Lời khuyên chăm sóc da và gợi ý dành cho loại da này
	CareTips string `db:"care_tips"`

	// Color is the hex color code used to represent this skin type in UI
	// Mã màu hex được sử dụng để đại diện cho loại da này trong giao diện
	Color string `db:"color"`

	// Icon is the icon identifier or URL for visual representation
	// Định danh biểu tượng hoặc URL để hiển thị trực quan
	Icon string `db:"icon"`

	// IsActive indicates whether this skin type is currently available for selection
	// Cho biết loại da này có đang khả dụng để lựa chọn không
	IsActive bool `db:"is_active"`

	entity.BaseEntity
}

type PaginatedSkinTypeEntity = entity.Pagination[SkinType]
