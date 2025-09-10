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

func (r *ProductRepository) PrivateCreateTag(ctx context.Context, tag *entity.Tag) error {
	query := `
		INSERT INTO tags (name, slug, description, color, icon, is_active)
		VALUES (:name, :slug, :description, :color, :icon, :is_active)
	`
	_, err := r.DB.NamedExecContext(ctx, query, tag)
	if err != nil {
		logger.Error("ProductRepository:PrivateCreateTag:Error: ", err)
		return err
	}
	return nil
}

func (r *ProductRepository) PrivateGetTags(ctx context.Context, params params.QueryParams) (*entity.PaginatedTagEntity, error) {
	// Tính offset cho pagination
	offset := (params.PageNumber - 1) * params.PageSize

	// Base query để lấy tags
	baseQuery := `FROM tags t`

	// Thêm điều kiện search nếu có
	var whereClause string
	var args []interface{}

	conditions := []string{}
	argIndex := 1

	if params.Search != "" {
		conditions = append(conditions, fmt.Sprintf("(t.name ILIKE $%d OR t.slug ILIKE $%d)", argIndex, argIndex))
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
		logger.Error("ProductRepository:PrivateGetTags - Count", err)
		return nil, err
	}

	// Query để lấy data với pagination
	dataQuery := `
		SELECT 
			t.id, 
			t.name, 
			t.slug, 
			t.description, 
			t.color,
			t.icon,
			t.is_active,
			t.created_at, 
			t.updated_at
	` + baseQuery + whereClause + `
		ORDER BY t.name ASC, t.created_at DESC
		LIMIT $` + fmt.Sprintf("%d", argIndex) + ` OFFSET $` + fmt.Sprintf("%d", argIndex+1)

	// Thêm params cho pagination
	args = append(args, params.PageSize, offset)

	var tags []entity.Tag
	err = r.DB.SelectContext(ctx, &tags, dataQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:PrivateGetTags - Select", err)
		return nil, err
	}

	// Tạo response pagination
	response := &entity.PaginatedTagEntity{
		Items:      tags,
		TotalItems: totalItems,
		PageNumber: params.PageNumber,
		PageSize:   params.PageSize,
	}

	return response, nil
}

func (r *ProductRepository) PrivateGetTagById(ctx context.Context, id uuid.UUID) (*entity.Tag, error) {
	query := `
		SELECT 
			t.id, 
			t.name, 
			t.slug, 
			t.description, 
			t.color,
			t.icon,
			t.is_active,
			t.created_at, 
			t.updated_at
		FROM tags t
		WHERE t.id = $1
	`
	var tag entity.Tag
	err := r.DB.GetContext(ctx, &tag, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:PrivateGetTagById - Select", err)
		return nil, err
	}
	return &tag, nil
}

func (r *ProductRepository) PrivateUpdateTag(ctx context.Context, tag *entity.Tag, id uuid.UUID) error {
	query := `
		UPDATE tags
		SET name = :name, slug = :slug, description = :description, color = :color, 
		    icon = :icon, is_active = :is_active 
		WHERE id = :id
	`

	// Set ID vào brand struct để sử dụng trong named query
	tag.ID = id

	result, err := r.DB.NamedExecContext(ctx, query, tag)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		logger.Error("ProductRepository:PrivateUpdateTag", err)
		return err
	}

	// Kiểm tra xem có record nào được update không
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("ProductRepository:PrivateUpdateIngredient - RowsAffected", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("ingredient with id %s not found", id)
	}

	return nil
}

func (r *ProductRepository) PrivateDeleteTag(ctx context.Context, id uuid.UUID) error {
	query := `
        DELETE FROM tags
        WHERE id = :id
    `
	tag := &entity.Tag{ID: id}
	result, err := r.DB.NamedExecContext(ctx, query, tag)
	if err != nil {
		logger.Error("ProductRepository:PrivateDeleteTag", err)
		return err
	}
	// Kiểm tra xem có record nào được delete không
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("ProductRepository:PrivateDeleteTag - RowsAffected", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("tag with id %s not found", id)
	}

	return nil
}
