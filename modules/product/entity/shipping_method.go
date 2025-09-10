package entity

import (
	"go-api-starter/core/entity"
)

// ShippingMethod represents a shipping method entity
// Đại diện cho phương thức vận chuyển
type ShippingMethod struct {
	// ID is the auto-incrementing primary key
	// ID tự tăng làm khóa chính
	ID int `db:"id"`

	// Name is the display name of the shipping method
	// Tên hiển thị của phương thức vận chuyển
	Name string `db:"name"`

	// Description provides detailed information about the shipping method
	// Mô tả chi tiết về phương thức vận chuyển
	Description string `db:"description"`

	// Provider is the shipping company (e.g., "Giao hàng nhanh", "Giao hàng tiết kiệm")
	// Nhà cung cấp vận chuyển (ví dụ: "Giao hàng nhanh", "Giao hàng tiết kiệm")
	Provider string `db:"provider"`

	// BaseCost is the base shipping cost in VND
	// Chi phí vận chuyển cơ bản tính bằng VND
	BaseCost float64 `db:"base_cost"`

	// CostPerKg is the additional cost per kilogram
	// Chi phí bổ sung theo từng kilogram
	CostPerKg float64 `db:"cost_per_kg"`

	// FreeShippingThreshold is the minimum order amount for free shipping
	// Ngưỡng đơn hàng tối thiểu để được miễn phí vận chuyển
	FreeShippingThreshold float64 `db:"free_shipping_threshold"`

	// EstimatedDaysMin is the minimum estimated delivery days
	// Số ngày giao hàng ước tính tối thiểu
	EstimatedDaysMin int `db:"estimated_days_min"`

	// EstimatedDaysMax is the maximum estimated delivery days
	// Số ngày giao hàng ước tính tối đa
	EstimatedDaysMax int `db:"estimated_days_max"`

	// IsActive indicates whether this shipping method is currently available
	// Cho biết phương thức vận chuyển này có đang khả dụng không
	IsActive bool `db:"is_active"`

	// SortOrder determines the display order in the list
	// Thứ tự hiển thị trong danh sách
	SortOrder int `db:"sort_order"`

	entity.BaseEntity
}

type PaginatedShippingMethodEntity = entity.Pagination[ShippingMethod]
