package entity

import (
	"go-api-starter/core/entity"
)

// PaymentMethod represents a payment method entity
// Đại diện cho phương thức thanh toán
type PaymentMethod struct {
	// ID is the unique identifier for the payment method
	// ID là định danh duy nhất cho phương thức thanh toán
	ID int `db:"id"`

	// Name is the display name of the payment method
	// Tên hiển thị của phương thức thanh toán
	Name string `db:"name"`

	// Description provides detailed information about the payment method
	// Mô tả cung cấp thông tin chi tiết về phương thức thanh toán
	Description string `db:"description"`

	// Provider is the payment service provider (e.g., "VNPay", "MoMo", "ZaloPay")
	// Nhà cung cấp dịch vụ thanh toán (ví dụ: "VNPay", "MoMo", "ZaloPay")
	Provider string `db:"provider"`

	// Type indicates the payment type (e.g., "wallet", "bank_transfer", "cod")
	// Loại thanh toán (ví dụ: "wallet", "bank_transfer", "cod")
	Type string `db:"type"`

	// IsActive indicates whether this payment method is currently available
	// Cho biết phương thức thanh toán này có đang khả dụng không
	IsActive bool `db:"is_active"`

	// SortOrder determines the display order in the payment method list
	// Thứ tự sắp xếp xác định thứ tự hiển thị trong danh sách phương thức thanh toán
	SortOrder int `db:"sort_order"`

	entity.BaseEntity
}

type PaginatedPaymentMethodEntity = entity.Pagination[PaymentMethod]
