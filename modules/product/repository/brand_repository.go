package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"go-api-starter/core/logger"
	"go-api-starter/core/params"
	"go-api-starter/modules/product/entity"

	"github.com/google/uuid"
)

func (r *ProductRepository) PrivateCreateBrand(ctx context.Context, brand *entity.Brand) error {
	query := `
		INSERT INTO brands (name, slug, description, logo, website, country, founded_year, is_active, is_featured, sort_order)
		VALUES (:name, :slug, :description, :logo, :website, :country, :founded_year, :is_active, :is_featured, :sort_order)
	`

	_, err := r.DB.NamedExecContext(ctx, query, brand)
	if err != nil {
		logger.Error("ProductRepository:PrivateCreateBrand:Error: ", err)
		return err
	}

	return nil
}

func (r *ProductRepository) PrivateGetBrandById(ctx context.Context, id uuid.UUID) (*entity.Brand, error) {
	query := `
		SELECT id, name, slug, description, logo, website, country, founded_year, is_active, is_featured, sort_order
		FROM brands
		WHERE id = $1
	`

	var brand entity.Brand
	err := r.DB.GetContext(ctx, &brand, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:GetBrandByID:Error: ", err)
		return nil, err
	}

	return &brand, nil
}

func (r *ProductRepository) PrivateGetBrands(ctx context.Context, params params.QueryParams) (*entity.PaginatedBrandResponse, error) {
	// Tính offset cho pagination
	offset := (params.PageNumber - 1) * params.PageSize

	// Base query để lấy brands
	baseQuery := `FROM brands b`

	// Thêm điều kiện search nếu có
	var whereClause string
	var args []interface{}

	conditions := []string{}
	argIndex := 1

	if params.Search != "" {
		conditions = append(conditions, fmt.Sprintf("b.name ILIKE $%d", argIndex))
		args = append(args, "%"+params.Search+"%")
		argIndex++
	}

	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}

	// Query để đếm tổng số records
	countQuery := "SELECT COUNT(*) " + baseQuery + whereClause

	var totalItems int
	err := r.DB.GetContext(ctx, &totalItems, countQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:PrivateGetBrands - Count", err)
		return nil, err
	}

	// Query để lấy data với pagination
	dataQuery := `
		SELECT 
			b.id, 
			b.name, 
			b.slug, 
			b.description, 
			b.logo, 
			b.website,
			b.country,
			b.founded_year,
			b.is_active,
			b.is_featured,
			b.sort_order, 
			b.created_at, 
			b.updated_at
	` + baseQuery + whereClause + `
		ORDER BY b.sort_order ASC, b.created_at DESC
		LIMIT $` + fmt.Sprintf("%d", argIndex) + ` OFFSET $` + fmt.Sprintf("%d", argIndex+1)

	// Thêm params cho pagination
	args = append(args, params.PageSize, offset)

	var brands []entity.Brand
	err = r.DB.SelectContext(ctx, &brands, dataQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:PrivateGetBrands - Select", err)
		return nil, err
	}

	// Tạo response pagination
	response := &entity.PaginatedBrandResponse{
		Items:      brands,
		TotalItems: totalItems,
		PageNumber: params.PageNumber,
		PageSize:   params.PageSize,
	}

	return response, nil
}

func (r *ProductRepository) PrivateUpdateBrand(ctx context.Context, brand *entity.Brand, id uuid.UUID) error {
	query := `
		UPDATE brands
		SET name = :name, slug = :slug, description = :description, logo = :logo, 
		    website = :website, country = :country, founded_year = :founded_year, 
		    is_active = :is_active, is_featured = :is_featured, sort_order = :sort_order 
		WHERE id = :id
	`

	// Set ID vào brand struct để sử dụng trong named query
	brand.ID = id

	result, err := r.DB.NamedExecContext(ctx, query, brand)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		logger.Error("ProductRepository:PrivateUpdateBrand", err)
		return err
	}

	// Kiểm tra xem có record nào được update không
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("ProductRepository:PrivateUpdateBrand - RowsAffected", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("brand with id %s not found", id)
	}

	return nil
}
func (r *ProductRepository) PrivateDeleteBrand(ctx context.Context, id uuid.UUID) error {
	query := `
		DELETE FROM brands
		WHERE id = $1
	`

	err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		logger.Error("ProductRepository:PrivateDeleteBrand:Error: ", err)
		return err
	}

	return nil
}
