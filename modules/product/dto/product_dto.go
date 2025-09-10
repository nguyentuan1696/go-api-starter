package dto

import (
	"time"
	"go-api-starter/core/dto"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ProductRequest struct {
	Name               string         `json:"name"`
	Slug               string         `json:"-"`
	Description        string         `json:"description"`
	ProductDescription string         `json:"product_description"`
	UsageInstructions  string         `json:"usage_instructions"`
	BrandID            uuid.UUID      `json:"brand_id"`
	CategoryID         uuid.UUID      `json:"category_id"`
	Price              float64        `json:"price"`
	OriginalPrice      float64        `json:"original_price"`
	SKU                *string        `json:"sku"`
	Barcode            *string        `json:"barcode"`
	RegistrationNumber string         `json:"registration_number"`
	Weight             float64        `json:"weight"`
	Volume             string         `json:"volume"`
	Thumbnail          string         `json:"thumbnail"`
	Images             pq.StringArray `json:"images"`
	IsActive           bool           `json:"is_active"`
	IsFeatured         bool           `json:"is_featured"`
	StockQuantity      int            `json:"stock_quantity"`
	MinStockLevel      int            `json:"min_stock_level"`
	Ingredients        string         `json:"ingredients"`
}

type ProductResponse struct {
	ID                 uuid.UUID      `json:"id"`
	Name               string         `json:"name"`
	Slug               string         `json:"slug"`
	Description        string         `json:"description"`
	ProductDescription string         `json:"product_description"`
	UsageInstructions  string         `json:"usage_instructions"`
	BrandID            uuid.UUID      `json:"brand_id"`
	CategoryID         uuid.UUID      `json:"category_id"`
	Price              float64        `json:"price"`
	OriginalPrice      float64        `json:"original_price"`
	SKU                *string        `json:"sku"`
	Barcode            *string        `json:"barcode"`
	RegistrationNumber string         `json:"registration_number"`
	Weight             float64        `json:"weight"`
	Volume             string         `json:"volume"`
	Thumbnail          string         `json:"thumbnail"`
	Images             pq.StringArray `json:"images"`
	IsActive           bool           `json:"is_active"`
	IsFeatured         bool           `json:"is_featured"`
	StockQuantity      int            `json:"stock_quantity"`
	MinStockLevel      int            `json:"min_stock_level"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	Ingredients        *string        `json:"ingredients"`
}

type ProductDetailResponse struct {
	ProductResponse
	BrandName    string `json:"brand_name"`
	CategoryName string `json:"category_name"`
}

type PaginatedProductDTO = dto.Pagination[ProductResponse]
