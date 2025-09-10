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

func (r *ProductRepository) PrivateCreateCategory(ctx context.Context, category *entity.Category) error {
	query := `
		INSERT INTO categories (name, slug, description, thumbnail, parent_id, sort_order, is_active)
		VALUES (:name, :slug, :description, :thumbnail, :parent_id, :sort_order, :is_active)
	`
	_, err := r.DB.NamedExecContext(ctx, query, category)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		logger.Error("ProductRepository:PrivateCreateCategory", err)
		return err
	}
	return nil
}

func (r *ProductRepository) PrivateUpdateCategory(ctx context.Context, category *entity.Category, id uuid.UUID) error {
	query := `
		UPDATE categories
		SET name = $1, slug = $2, description = $3, thumbnail = $4, parent_id = $5, 
		    sort_order = $6, is_active = $7, updated_at = now()
		WHERE id = $8
	`

	result, err := r.DB.SQLx().ExecContext(ctx, query,
		category.Name,
		category.Slug,
		category.Description,
		category.Thumbnail,
		category.ParentID,
		category.SortOrder,
		category.IsActive,
		id,
	)

	if err != nil {
		logger.Error("ProductRepository:PrivateUpdateCategory", err)
		return err
	}

	// Kiểm tra xem có record nào được update không
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("ProductRepository:PrivateUpdateCategory - RowsAffected", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("category with id %s not found", id)
	}

	return nil
}

func (r *ProductRepository) PrivateDeleteCategory(ctx context.Context, id uuid.UUID) error {
	query := `
		DELETE FROM categories
		WHERE id = :id
	`
	_, err := r.DB.NamedExecContext(ctx, query, map[string]any{"id": id})
	if err != nil {
		logger.Error("ProductRepository:PrivateDeleteCategory", err)
		return err
	}
	return nil
}

func (r *ProductRepository) PrivateGetCategoryById(ctx context.Context, id uuid.UUID) (*entity.Category, error) {
	var category entity.Category
	query := `
		SELECT 
			c.id, 
			c.name, 
			c.slug, 
			c.description, 
			c.thumbnail, 
			c.parent_id,
			p.name as parent_name,
			c.sort_order, 
			c.is_active, 
			c.created_at, 
			c.updated_at
		FROM categories c
		LEFT JOIN categories p ON c.parent_id = p.id
		WHERE c.id = $1
	`
	err := r.DB.GetContext(ctx, &category, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:PrivateGetCategoryById", err)
		return nil, err
	}
	return &category, nil
}

func (r *ProductRepository) PrivateGetCategories(ctx context.Context, params params.QueryParams) (*entity.PaginatedCategoryResponse, error) {
	// Tính offset cho pagination
	offset := (params.PageNumber - 1) * params.PageSize

	// Base query để lấy categories với LEFT JOIN
	baseQuery := `FROM categories c LEFT JOIN categories p ON c.parent_id = p.id`

	// Thêm điều kiện search nếu có
	var whereClause string
	var args []interface{}

	conditions := []string{}
	argIndex := 1

	if params.Search != "" {
		conditions = append(conditions, fmt.Sprintf("c.name ILIKE $%d", argIndex))
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
		logger.Error("ProductRepository:PrivateGetCategories - Count", err)
		return nil, err
	}

	// Query để lấy data với pagination
	dataQuery := `
		SELECT 
			c.id, 
			c.name, 
			c.slug, 
			c.description, 
			c.thumbnail, 
			c.parent_id,
			p.name as parent_name,
			c.sort_order, 
			c.is_active, 
			c.created_at, 
			c.updated_at
	` + baseQuery + whereClause + `
		ORDER BY c.sort_order ASC, c.created_at DESC
		LIMIT $` + fmt.Sprintf("%d", argIndex) + ` OFFSET $` + fmt.Sprintf("%d", argIndex+1)

	// Thêm params cho pagination
	args = append(args, params.PageSize, offset)

	var categories []entity.Category
	err = r.DB.SelectContext(ctx, &categories, dataQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:PrivateGetCategories - Select", err)
		return nil, err
	}

	// Tạo response pagination
	response := &entity.PaginatedCategoryResponse{
		Items:      categories,
		TotalItems: totalItems,
		PageNumber: params.PageNumber,
		PageSize:   params.PageSize,
	}

	return response, nil
}
