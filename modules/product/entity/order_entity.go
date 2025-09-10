package entity

import (
	"encoding/json"
	"time"
	"go-api-starter/core/entity"

	"github.com/google/uuid"
)

type Order struct {
	ID          uuid.UUID  `db:"id"`
	OrderNumber string     `db:"order_number"`
	CustomerID  *uuid.UUID `db:"customer_id"`

	// Thông tin khách hàng (snapshot)
	CustomerEmail string `db:"customer_email"`
	CustomerPhone string `db:"customer_phone"`
	CustomerName  string `db:"customer_name"`

	// Địa chỉ giao hàng (snapshot)
	ShippingRecipientName  string `db:"shipping_recipient_name"`
	ShippingRecipientPhone string `db:"shipping_recipient_phone"`
	ShippingAddress        string `db:"shipping_address"`
	ShippingWardName       string `db:"shipping_ward_name"`
	ShippingDistrictName   string `db:"shipping_district_name"`
	ShippingProvinceName   string `db:"shipping_province_name"`

	// Trạng thái đơn hàng
	OrderState    string `db:"order_state"`
	PaymentStatus string `db:"payment_status"`

	// Tính toán giá
	Subtotal       float64 `db:"subtotal"`
	ShippingCost   float64 `db:"shipping_cost"`
	TaxAmount      float64 `db:"tax_amount"`
	DiscountAmount float64 `db:"discount_amount"`
	TotalAmount    float64 `db:"total_amount"`

	// Thông tin vận chuyển và thanh toán
	ShippingMethodID   int    `db:"shipping_method_id"`
	ShippingMethodName string `db:"shipping_method_name"`
	PaymentMethodID    int    `db:"payment_method_id"`
	PaymentMethodName  string `db:"payment_method_name"`

	// Mã giảm giá
	CouponID             *uuid.UUID `db:"coupon_id"`
	CouponCode           string     `db:"coupon_code"`
	CouponDiscountAmount float64    `db:"coupon_discount_amount"`

	// Ghi chú
	Notes      string `db:"notes"`
	AdminNotes string `db:"admin_notes"`

	// Thời gian quan trọng
	OrderedAt   time.Time  `db:"ordered_at"`
	ConfirmedAt *time.Time `db:"confirmed_at"`
	ShippedAt   *time.Time `db:"shipped_at"`
	DeliveredAt *time.Time `db:"delivered_at"`
	CancelledAt *time.Time `db:"cancelled_at"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	ID        uuid.UUID `db:"id"`
	OrderID   uuid.UUID `db:"order_id"`
	ProductID uuid.UUID `db:"product_id"`

	// Thông tin sản phẩm (snapshot)
	Name      string  `db:"name"`
	SKU       *string `db:"sku"`
	Thumbnail *string `db:"thumbnail"`

	// Giá và số lượng
	UnitPrice  float64 `db:"unit_price"`
	Quantity   int     `db:"quantity"`
	TotalPrice float64 `db:"total_price"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// PaginatedOrderEntity for paginated responses
type PaginatedOrderEntity = entity.Pagination[Order]

// OrderWithItems includes order items
type OrderWithItems struct {
	Order
	Items []OrderItem
}

// OrderSummary for list views
type OrderSummary struct {
	ID          uuid.UUID `db:"id"`
	OrderNumber string    `db:"order_number"`
	CustomerID  *string   `db:"customer_id"`

	CustomerEmail string `db:"customer_email"`
	CustomerName  string `db:"customer_name"`

	OrderState    string  `db:"order_state"`
	PaymentStatus string  `db:"payment_status"`
	TotalAmount   float64 `db:"total_amount"`

	OrderedAt time.Time `db:"ordered_at"`
	CreatedAt time.Time `db:"created_at"`
}

// PaginatedOrderItemEntity for paginated order items responses
type PaginatedOrderItemEntity = entity.Pagination[OrderItem]

// OrderItemSummary for order item list views
type OrderItemSummary struct {
	ID        uuid.UUID `db:"id"`
	OrderID   uuid.UUID `db:"order_id"`
	ProductID string    `db:"product_id"`

	ProductName string  `db:"product_name"`
	UnitPrice   float64 `db:"unit_price"`
	Quantity    int     `db:"quantity"`
	TotalPrice  float64 `db:"total_price"`

	CreatedAt time.Time `db:"created_at"`
}

type OrderDetailWithItems struct {
	OrderID                string          `db:"order_id"`
	OrderNumber            string          `db:"order_number"`
	CustomerID             *string         `db:"customer_id"`
	CustomerName           string          `db:"customer_name"`
	CustomerEmail          string          `db:"customer_email"`
	CustomerPhone          string          `db:"customer_phone"`
	ShippingRecipientName  string          `db:"shipping_recipient_name"`
	ShippingRecipientPhone string          `db:"shipping_recipient_phone"`
	ShippingAddress        string          `db:"shipping_address"`
	ShippingWardName       string          `db:"shipping_ward_name"`
	ShippingDistrictName   string          `db:"shipping_district_name"`
	ShippingProvinceName   string          `db:"shipping_province_name"`
	PaymentMethodID        int             `db:"payment_method_id"`
	PaymentMethodName      string          `db:"payment_method_name"`
	ShippingMethodID       int             `db:"shipping_method_id"`
	ShippingMethodName     string          `db:"shipping_method_name"`
	OrderState             string          `db:"order_state"`
	PaymentStatus          string          `db:"payment_status"`
	Notes                  string          `db:"notes"`
	AdminNotes             string          `db:"admin_notes"`
	Subtotal               float64         `db:"subtotal"`
	ShippingCost           float64         `db:"shipping_cost"`
	TaxAmount              float64         `db:"tax_amount"`
	DiscountAmount         float64         `db:"discount_amount"`
	TotalAmount            float64         `db:"total_amount"`
	OrderedAt              *time.Time      `db:"ordered_at"`
	ConfirmedAt            *time.Time      `db:"confirmed_at"`
	ShippedAt              *time.Time      `db:"shipped_at"`
	DeliveredAt            *time.Time      `db:"delivered_at"`
	CancelledAt            *time.Time      `db:"cancelled_at"`
	Items                  json.RawMessage `db:"items"`
}
