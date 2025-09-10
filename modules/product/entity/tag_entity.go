package entity

import (
	"time"
	"go-api-starter/core/entity"

	"github.com/google/uuid"
)

// Tag represents a product tag for categorization and filtering
// Đại diện cho thẻ sản phẩm dùng để phân loại và lọc
type Tag struct {
	// ID is the unique identifier for the tag
	// ID là định danh duy nhất cho thẻ
	ID uuid.UUID `db:"id"`

	// Name is the display name of the tag (e.g., "Organic", "Anti-aging", "Sensitive Skin")
	// Tên hiển thị của thẻ (ví dụ: "Hữu cơ", "Chống lão hóa", "Da nhạy cảm")
	Name string `db:"name"`

	// Slug is the URL-friendly version of the tag name
	// Slug là phiên bản thân thiện URL của tên thẻ
	Slug string `db:"slug"`

	// Description is a brief description of what this tag represents
	// Mô tả ngắn gọn về ý nghĩa của thẻ này
	Description string `db:"description"`

	// Color is the hex color code used to display this tag in UI
	// Mã màu hex được sử dụng để hiển thị thẻ này trong giao diện
	Color string `db:"color"`

	// Icon is the icon identifier or URL for visual representation of the tag
	// Định danh biểu tượng hoặc URL để hiển thị trực quan cho thẻ
	Icon string `db:"icon"`

	// IsActive indicates whether this tag is currently available for use
	// Cho biết thẻ này có đang khả dụng để sử dụng không
	IsActive bool `db:"is_active"`

	// CreatedAt is the timestamp when the record was created
	// Thời gian tạo bản ghi
	CreatedAt time.Time `db:"created_at"`

	// UpdatedAt is the timestamp when the record was last updated
	// Thời gian cập nhật bản ghi lần cuối
	UpdatedAt time.Time `db:"updated_at"`
}

type PaginatedTagEntity = entity.Pagination[Tag]
