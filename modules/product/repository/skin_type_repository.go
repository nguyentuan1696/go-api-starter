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

func (r *ProductRepository) PrivateCreateSkinType(ctx context.Context, skinType *entity.SkinType) error {
	query := `
		INSERT INTO skin_types (name, slug, description, characteristics, care_tips, color, icon, is_active)
		VALUES (:name, :slug, :description, :characteristics, :care_tips, :color, :icon, :is_active)
	`
	_, err := r.DB.NamedExecContext(ctx, query, skinType)
	if err != nil {
		logger.Error("ProductRepository:PrivateCreateSkinType:Error: ", err)
		return err
	}
	return nil
}

func (r *ProductRepository) PrivateGetSkinTypes(ctx context.Context, params params.QueryParams) (*entity.PaginatedSkinTypeEntity, error) {
	// Tính offset cho pagination
	offset := (params.PageNumber - 1) * params.PageSize

	// Base query để lấy skin types
	baseQuery := `FROM skin_types st`

	// Thêm điều kiện search nếu có
	var whereClause string
	var args []interface{}

	conditions := []string{}
	argIndex := 1

	if params.Search != "" {
		conditions = append(conditions, fmt.Sprintf("(st.name ILIKE $%d OR st.slug ILIKE $%d)", argIndex, argIndex))
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
		logger.Error("ProductRepository:PrivateGetSkinTypes - Count", err)
		return nil, err
	}

	// Query để lấy data với pagination
	dataQuery := `
		SELECT 
			st.id, 
			st.name, 
			st.slug, 
			st.description, 
			st.characteristics,
			st.care_tips,
			st.color,
			st.icon,
			st.is_active,
			st.created_at, 
			st.updated_at
	` + baseQuery + whereClause + `
		ORDER BY st.name ASC, st.created_at DESC
		OFFSET $` + fmt.Sprintf("%d", argIndex+1) + ` ROWS FETCH NEXT $` + fmt.Sprintf("%d", argIndex) + ` ROWS ONLY`

	// Thêm params cho pagination
	args = append(args, params.PageSize, offset)

	var skinTypes []entity.SkinType
	err = r.DB.SelectContext(ctx, &skinTypes, dataQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:PrivateGetSkinTypes - Select", err)
		return nil, err
	}

	// Tạo response pagination
	response := &entity.PaginatedSkinTypeEntity{
		Items:      skinTypes,
		TotalItems: totalItems,
		PageNumber: params.PageNumber,
		PageSize:   params.PageSize,
	}

	return response, nil
}

func (r *ProductRepository) PrivateGetSkinTypeById(ctx context.Context, id uuid.UUID) (*entity.SkinType, error) {
	query := `
		SELECT 
			st.id, 
			st.name, 
			st.slug, 
			st.description, 
			st.characteristics,
			st.care_tips,
			st.color,
			st.icon,
			st.is_active,
			st.created_at, 
			st.updated_at
		FROM skin_types st
		WHERE st.id = $1
	`
	var skinType entity.SkinType
	err := r.DB.GetContext(ctx, &skinType, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:PrivateGetSkinTypeById - Select", err)
		return nil, err
	}
	return &skinType, nil
}

func (r *ProductRepository) PrivateUpdateSkinType(ctx context.Context, skinType *entity.SkinType, id uuid.UUID) error {
	query := `
		UPDATE skin_types
		SET name = :name, slug = :slug, description = :description, characteristics = :characteristics,
		    care_tips = :care_tips, color = :color, icon = :icon, is_active = :is_active
		WHERE id = :id
	`

	// Set ID vào skinType struct để sử dụng trong named query
	skinType.ID = id

	result, err := r.DB.NamedExecContext(ctx, query, skinType)
	if err != nil {
		logger.Error("ProductRepository:PrivateUpdateSkinType", err)
		return err
	}

	// Kiểm tra xem có record nào được update không
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("ProductRepository:PrivateUpdateSkinType - RowsAffected", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("skin type with id %s not found", id)
	}

	return nil
}

func (r *ProductRepository) PrivateDeleteSkinType(ctx context.Context, id uuid.UUID) error {
	query := `
		DELETE FROM skin_types
		WHERE id = $1
	`
	err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		logger.Error("ProductRepository:PrivateDeleteSkinType", err)
		return err
	}
	return nil
}
