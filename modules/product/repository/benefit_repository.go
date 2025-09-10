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

func (r *ProductRepository) PrivateCreateBenefit(ctx context.Context, benefit *entity.Benefit) error {
	query := `
		INSERT INTO benefits (name, slug, description, category, target_area, 
		                     effectiveness_level, time_to_see_results, color, icon, is_active, 
		                     is_featured, sort_order)
		VALUES (:name, :slug, :description, :category, :target_area, 
		        :effectiveness_level, :time_to_see_results, :color, :icon, :is_active, 
		        :is_featured, :sort_order)
	`
	_, err := r.DB.NamedExecContext(ctx, query, benefit)
	if err != nil {
		logger.Error("ProductRepository:PrivateCreateBenefit:Error: ", err)
		return err
	}
	return nil
}

func (r *ProductRepository) PrivateGetBenefits(ctx context.Context, params params.QueryParams) (*entity.PaginatedBenefitEntity, error) {
	// Tính offset cho pagination
	offset := (params.PageNumber - 1) * params.PageSize

	// Base query để lấy benefits
	baseQuery := `FROM benefits b`

	// Thêm điều kiện search nếu có
	var whereClause string
	var args []interface{}

	conditions := []string{}
	argIndex := 1

	if params.Search != "" {
		conditions = append(conditions, fmt.Sprintf("(b.name ILIKE $%d OR b.slug ILIKE $%d)", argIndex, argIndex))
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
		logger.Error("ProductRepository:PrivateGetBenefits - Count", err)
		return nil, err
	}

	// Query để lấy data với pagination
	dataQuery := `
		SELECT 
			b.id, 
			b.name, 
			b.slug, 
			b.description, 
			b.category,
			b.target_area,
			b.effectiveness_level,
			b.time_to_see_results,
			b.color,
			b.icon,
			b.is_active,
			b.is_featured,
			b.sort_order,
			b.created_at, 
			b.updated_at
	` + baseQuery + whereClause + `
		ORDER BY b.sort_order ASC, b.name ASC, b.created_at DESC
		OFFSET $` + fmt.Sprintf("%d", argIndex+1) + ` ROWS FETCH NEXT $` + fmt.Sprintf("%d", argIndex) + ` ROWS ONLY`

	// Thêm params cho pagination
	args = append(args, params.PageSize, offset)

	var benefits []entity.Benefit
	err = r.DB.SelectContext(ctx, &benefits, dataQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:PrivateGetBenefits - Select", err)
		return nil, err
	}

	// Tạo response pagination
	response := &entity.PaginatedBenefitEntity{
		Items:      benefits,
		TotalItems: totalItems,
		PageNumber: params.PageNumber,
		PageSize:   params.PageSize,
	}

	return response, nil
}

func (r *ProductRepository) PrivateGetBenefitById(ctx context.Context, id uuid.UUID) (*entity.Benefit, error) {
	query := `
		SELECT 
			b.id, 
			b.name, 
			b.slug, 
			b.description, 
			b.category,
			b.target_area,
			b.effectiveness_level,
			b.time_to_see_results,
			b.color,
			b.icon,
			b.is_active,
			b.is_featured,
			b.sort_order,
			b.created_at, 
			b.updated_at
		FROM benefits b
		WHERE b.id = $1
	`
	var benefit entity.Benefit
	err := r.DB.GetContext(ctx, &benefit, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:PrivateGetBenefitById - Select", err)
		return nil, err
	}
	return &benefit, nil
}

func (r *ProductRepository) PrivateUpdateBenefit(ctx context.Context, benefit *entity.Benefit, id uuid.UUID) error {
	query := `
		UPDATE benefits
		SET name = :name, slug = :slug, description = :description,
		    category = :category, target_area = :target_area, effectiveness_level = :effectiveness_level,
		    time_to_see_results = :time_to_see_results, color = :color, icon = :icon, 
		    is_active = :is_active, is_featured = :is_featured, sort_order = :sort_order
		WHERE id = :id
	`

	// Set ID vào benefit struct để sử dụng trong named query
	benefit.ID = id

	result, err := r.DB.NamedExecContext(ctx, query, benefit)
	if err != nil {
		logger.Error("ProductRepository:PrivateUpdateBenefit", err)
		return err
	}

	// Kiểm tra xem có record nào được update không
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("ProductRepository:PrivateUpdateBenefit - RowsAffected", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("benefit with id %s not found", id)
	}

	return nil
}

func (r *ProductRepository) PrivateDeleteBenefit(ctx context.Context, id uuid.UUID) error {
	query := `
		DELETE FROM benefits
		WHERE id = $1
	`
	err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		logger.Error("ProductRepository:PrivateDeleteBenefit", err)
		return err
	}
	return nil
}
