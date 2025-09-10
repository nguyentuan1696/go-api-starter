package entity

import (
	"go-api-starter/core/entity"
)

// Ingredient represents a cosmetic ingredient entity
// Đại diện cho thành phần mỹ phẩm
type Ingredient struct {

	// Name is the common display name of the ingredient
	// Tên hiển thị thông thường của thành phần
	Name string `db:"name"`

	// Slug is the URL-friendly version of the ingredient name
	// Slug là phiên bản thân thiện URL của tên thành phần
	Slug string `db:"slug"`

	// InciName is the International Nomenclature of Cosmetic Ingredients name
	// InciName là tên theo Danh pháp Quốc tế về Thành phần Mỹ phẩm
	InciName string `db:"inci_name"`

	// Description provides detailed information about the ingredient
	// Mô tả cung cấp thông tin chi tiết về thành phần
	Description string `db:"description"`

	// Origin indicates where the ingredient comes from (natural, synthetic, etc.)
	// Nguồn gốc cho biết thành phần đến từ đâu (tự nhiên, tổng hợp, v.v.)
	Origin string `db:"origin"`

	// Function describes what the ingredient does in cosmetic products
	// Chức năng mô tả thành phần làm gì trong sản phẩm mỹ phẩm
	Function string `db:"function"`

	// CasNumber is the Chemical Abstracts Service registry number
	// CasNumber là số đăng ký của Dịch vụ Tóm tắt Hóa học
	CasNumber string `db:"cas_number"`

	// EwgScore is the Environmental Working Group safety score (1-10)
	// EwgScore là điểm an toàn của Nhóm Làm việc Môi trường (1-10)
	EwgScore int `db:"ewg_score"`

	// IsRestricted indicates if the ingredient has usage restrictions
	// Cho biết thành phần có bị hạn chế sử dụng không
	IsRestricted bool `db:"is_restricted"`

	// IsBanned indicates if the ingredient is banned in certain regions
	// Cho biết thành phần có bị cấm ở một số khu vực không
	IsBanned bool `db:"is_banned"`

	entity.BaseEntity
}

type PaginatedIngredientEntity = entity.Pagination[Ingredient]
