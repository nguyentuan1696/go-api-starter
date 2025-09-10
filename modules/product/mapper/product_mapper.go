package mapper

import (
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/entity"
)

func ToProductDetailDTO(entity *entity.ProductDetailEntity) *dto.ProductDetailResponse {
	return &dto.ProductDetailResponse{
		ProductResponse: *ToProductDTO(&entity.Product),
		BrandName:       entity.BrandName,
		CategoryName:    entity.CategoryName,
	}
}

func ToProductEntity(dto *dto.ProductRequest) *entity.Product {
	return &entity.Product{
		Name:               dto.Name,
		Slug:               dto.Slug,
		Description:        dto.Description,
		ProductDescription: dto.ProductDescription,
		UsageInstructions:  dto.UsageInstructions,
		BrandID:            dto.BrandID,
		CategoryID:         dto.CategoryID,
		Price:              dto.Price,
		OriginalPrice:      dto.OriginalPrice,
		SKU:                dto.SKU,
		Barcode:            dto.Barcode,
		RegistrationNumber: dto.RegistrationNumber,
		Weight:             dto.Weight,
		Volume:             dto.Volume,
		Thumbnail:          dto.Thumbnail,
		Images:             dto.Images,
		IsActive:           dto.IsActive,
		IsFeatured:         dto.IsFeatured,
		StockQuantity:      dto.StockQuantity,
		MinStockLevel:      dto.MinStockLevel,
		Ingredients:        &dto.Ingredients,
	}
}

func ToProductDTO(entity *entity.Product) *dto.ProductResponse {
	return &dto.ProductResponse{
		ID:                 entity.ID,
		Name:               entity.Name,
		Slug:               entity.Slug,
		Description:        entity.Description,
		ProductDescription: entity.ProductDescription,
		UsageInstructions:  entity.UsageInstructions,
		BrandID:            entity.BrandID,
		CategoryID:         entity.CategoryID,
		Price:              entity.Price,
		OriginalPrice:      entity.OriginalPrice,
		SKU:                entity.SKU,
		Barcode:            entity.Barcode,
		RegistrationNumber: entity.RegistrationNumber,
		Weight:             entity.Weight,
		Volume:             entity.Volume,
		Thumbnail:          entity.Thumbnail,
		Images:             entity.Images,
		IsActive:           entity.IsActive,
		IsFeatured:         entity.IsFeatured,
		StockQuantity:      entity.StockQuantity,
		MinStockLevel:      entity.MinStockLevel,
		CreatedAt:          entity.CreatedAt,
		UpdatedAt:          entity.UpdatedAt,
		Ingredients:        entity.Ingredients,
	}
}

func ToProductPaginationDTO(entity *entity.PaginatedProductEntity) *dto.PaginatedProductDTO {

	productResponses := make([]dto.ProductResponse, len(entity.Items))
	for i, tag := range entity.Items {
		productResponses[i] = *ToProductDTO(&tag)
	}

	// Tính total pages
	totalPages := 0
	if entity.PageSize > 0 {
		totalPages = (entity.TotalItems + entity.PageSize - 1) / entity.PageSize
	}

	return &dto.PaginatedProductDTO{
		Items:      productResponses,
		TotalItems: entity.TotalItems,
		TotalPages: totalPages,
		PageNumber: entity.PageNumber,
		PageSize:   entity.PageSize,
	}
}
