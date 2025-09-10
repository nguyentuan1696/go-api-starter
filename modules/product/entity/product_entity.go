package entity

import (
	"go-api-starter/core/entity"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// Product represents a cosmetic product entity
// Đại diện cho sản phẩm mỹ phẩm
type Product struct {

	// Name is the display name of the product
	// Tên hiển thị của sản phẩm
	Name string `db:"name"`

	// Slug is the URL-friendly version of the product name
	// Slug là phiên bản thân thiện URL của tên sản phẩm
	Slug string `db:"slug"`

	// Description is a brief description of the product
	// Mô tả ngắn gọn về sản phẩm
	Description string `db:"description"`

	// ProductDescription is a detailed description of the product
	// Mô tả chi tiết về sản phẩm
	ProductDescription string `db:"product_description"`

	// UsageInstructions provides instructions on how to use the product
	// Hướng dẫn sử dụng sản phẩm
	UsageInstructions string `db:"usage_instructions"`

	// BrandID is the foreign key reference to the brand
	// BrandID là khóa ngoại tham chiếu đến thương hiệu
	BrandID uuid.UUID `db:"brand_id"`

	// CategoryID is the foreign key reference to the category
	// CategoryID là khóa ngoại tham chiếu đến danh mục
	CategoryID uuid.UUID `db:"category_id"`

	// Price is the current selling price of the product
	// Giá bán hiện tại của sản phẩm
	Price float64 `db:"price"`

	// OriginalPrice is the original price before any discounts
	// Giá gốc trước khi có bất kỳ giảm giá nào
	OriginalPrice float64 `db:"original_price"`

	// SKU is the Stock Keeping Unit identifier
	// SKU là mã định danh đơn vị lưu kho
	SKU *string `db:"sku"`

	// Barcode is the product's barcode
	// Mã vạch của sản phẩm
	Barcode *string `db:"barcode"`

	// RegistrationNumber is the official registration number for cosmetics
	// Số đăng ký chính thức cho mỹ phẩm
	RegistrationNumber string `db:"registration_number"`

	// Weight is the product weight in grams
	// Trọng lượng sản phẩm tính bằng gram
	Weight float64 `db:"weight"`

	// Volume is the product volume in milliliters
	// Thể tích sản phẩm tính bằng milliliter
	Volume string `db:"volume"`

	// Thumbnail is the main product image URL
	// Hình ảnh chính của sản phẩm
	Thumbnail string `db:"thumbnail"`

	// Images is an array of additional product image URLs
	// Mảng các URL hình ảnh bổ sung của sản phẩm
	Images pq.StringArray `db:"images"`

	// IsActive indicates whether the product is currently available
	// Cho biết sản phẩm có đang khả dụng không
	IsActive bool `db:"is_active"`

	// IsFeatured indicates whether the product is featured
	// Cho biết sản phẩm có được nổi bật không
	IsFeatured bool `db:"is_featured"`

	// StockQuantity is the current stock quantity
	// Số lượng tồn kho hiện tại
	StockQuantity int `db:"stock_quantity"`

	// MinStockLevel is the minimum stock level before reordering
	// Mức tồn kho tối thiểu trước khi đặt hàng lại
	MinStockLevel int `db:"min_stock_level"`

	entity.BaseEntity

	// Ingredients is the list of ingredients in the product
	// Danh sách các thành phần trong sản phẩm
	Ingredients *string `db:"ingredients"`
}

type ProductDetailEntity struct {
	Product
	BrandName    string `db:"brand_name"`
	CategoryName string `db:"category_name"`
}

type PaginatedProductEntity = entity.Pagination[Product]
