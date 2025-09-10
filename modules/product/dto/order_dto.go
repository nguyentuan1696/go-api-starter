package dto

import (
	"encoding/json"
	"time"
	"go-api-starter/core/dto"

	"github.com/google/uuid"
)

type OrderRequest struct {
	OrderNumber            string     `json:"-"`
	CustomerID             *uuid.UUID `json:"customer_id"`
	CustomerEmail          string     `json:"customer_email"`
	CustomerPhone          string     `json:"customer_phone"`
	CustomerName           string     `json:"customer_name"`
	ShippingRecipientName  string     `json:"shipping_recipient_name"`
	ShippingRecipientPhone string     `json:"shipping_recipient_phone"`
	ShippingAddress        string     `json:"shipping_address"`
	ShippingWardName       string     `json:"shipping_ward_name"`
	ShippingDistrictName   string     `json:"shipping_district_name"`
	ShippingProvinceName   string     `json:"shipping_province_name"`
	PaymentMethodID        int        `json:"payment_method_id"`
	PaymentMethodName      string     `json:"payment_method_name"`
	ShippingMethodID       int        `json:"shipping_method_id"`
	ShippingMethodName     string     `json:"shipping_method_name"`
	OrderState             string     `json:"-"`
	PaymentStatus          string     `json:"-"`
	Subtotal               float64    `json:"-"`
	ShippingCost           float64    `json:"-"`
	TaxAmount              float64    `json:"-"`
	DiscountAmount         float64    `json:"-"`
	TotalAmount            float64    `json:"-"`
	CouponID               *uuid.UUID `json:"-"`
	CouponCode             string     `json:"coupon_code"`
	Notes                  string     `json:"notes"`
	AdminNotes             string     `json:"admin_notes"`
	OrderAt                time.Time  `json:"order_at"`
	ConfirmAt              time.Time  `json:"confirm_at"`
	ShippedAt              time.Time  `json:"shipped_at"`
	DeliveredAt            time.Time  `json:"delivered_at"`
	CancelledAt            time.Time  `json:"cancelled_at"`
}

type OrderResponse struct {
	ID                     uuid.UUID  `json:"id"`
	OrderNumber            string     `json:"order_number"`
	CustomerID             *uuid.UUID `json:"customer_id"`
	CustomerEmail          string     `json:"customer_email"`
	CustomerPhone          string     `json:"customer_phone"`
	CustomerName           string     `json:"customer_name"`
	ShippingRecipientName  string     `json:"shipping_recipient_name"`
	ShippingRecipientPhone string     `json:"shipping_recipient_phone"`
	ShippingAddress        string     `json:"shipping_address"`
	ShippingWardName       string     `json:"shipping_ward_name"`
	ShippingDistrictName   string     `json:"shipping_district_name"`
	ShippingProvinceName   string     `json:"shipping_province_name"`
	OrderState             string     `json:"order_state"`
	PaymentStatus          string     `json:"payment_status"`
	Subtotal               float64    `json:"subtotal"`
	ShippingCost           float64    `json:"shipping_cost"`
	TaxAmount              float64    `json:"tax_amount"`
	DiscountAmount         float64    `json:"discount_amount"`
	TotalAmount            float64    `json:"total_amount"`
	ShippingMethodID       int        `json:"shipping_method_id"`
	ShippingMethodName     string     `json:"shipping_method_name"`
	PaymentMethodID        int        `json:"payment_method_id"`
	PaymentMethodName      string     `json:"payment_method_name"`
	CouponID               *uuid.UUID `json:"coupon_id"`
	CouponCode             string     `json:"coupon_code"`
	CouponDiscountAmount   float64    `json:"coupon_discount_amount"`
	Notes                  string     `json:"notes"`
	AdminNotes             string     `json:"admin_notes"`
	OrderAt                time.Time  `json:"order_at"`
	ConfirmAt              time.Time  `json:"confirm_at"`
	ShippedAt              time.Time  `json:"shipped_at"`
	DeliveredAt            time.Time  `json:"delivered_at"`
	CancelledAt            time.Time  `json:"cancelled_at"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at"`
}

type OrderItemRequest struct {
	ProductID     uuid.UUID `json:"product_id"`
	ProductName   string    `json:"product_name"`
	VariantID     string    `json:"variant_id"`
	VariantName   string    `json:"variant_name"`
	Price         float64   `json:"price"`
	OriginalPrice float64   `json:"original_price"`
	Quantity      int       `json:"quantity"`
	Total         float64   `json:"total"`
}

type PlaceOrderRequest struct {
	OrderRequest
	OrderItems []OrderItemRequest `json:"order_items"`
}

type PlaceOrderResponse struct {
	OrderNumber            string  `json:"order_number"`
	ShippingRecipientName  string  `json:"shipping_recipient_name"`
	ShippingRecipientPhone string  `json:"shipping_recipient_phone"`
	ShippingAddress        string  `json:"shipping_address"`
	TotalAmount            float64 `json:"total_amount"`
}

type PaginatedOrderDTO = dto.Pagination[OrderResponse]

type OrderDetailWithItemsDTO struct {
	OrderID                string          `json:"order_id"`
	OrderNumber            string          `json:"order_number"`
	CustomerID             *string         `json:"customer_id"`
	CustomerName           string          `json:"customer_name"`
	CustomerEmail          string          `json:"customer_email"`
	CustomerPhone          string          `json:"customer_phone"`
	ShippingRecipientName  string          `json:"shipping_recipient_name"`
	ShippingRecipientPhone string          `json:"shipping_recipient_phone"`
	ShippingAddress        string          `json:"shipping_address"`
	ShippingWardName       string          `json:"shipping_ward_name"`
	ShippingDistrictName   string          `json:"shipping_district_name"`
	ShippingProvinceName   string          `json:"shipping_province_name"`
	OrderState             string          `json:"order_state"`
	Notes                  string          `json:"notes"`
	AdminNotes             string          `json:"admin_notes"`
	PaymentMethodID        int             `json:"payment_method_id"`
	PaymentMethodName      string          `json:"payment_method_name"`
	ShippingMethodID       int             `json:"shipping_method_id"`
	ShippingMethodName     string          `json:"shipping_method_name"`
	PaymentStatus          string          `json:"payment_status"`
	Subtotal               float64         `json:"subtotal"`
	ShippingCost           float64         `json:"shipping_cost"`
	TaxAmount              float64         `json:"tax_amount"`
	DiscountAmount         float64         `json:"discount_amount"`
	TotalAmount            float64         `json:"total_amount"`
	OrderedAt              *time.Time      `json:"ordered_at"`
	ConfirmedAt            *time.Time      `json:"confirmed_at,omitempty"`
	ShippedAt              *time.Time      `json:"shipped_at,omitempty"`
	DeliveredAt            *time.Time      `json:"delivered_at,omitempty"`
	CancelledAt            *time.Time      `json:"cancelled_at,omitempty"`
	Items                  json.RawMessage `json:"items"`
}
