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

func (r *ProductRepository) PrivateCreateIngredient(ctx context.Context, ingredient *entity.Ingredient) error {
	query := `
		INSERT INTO ingredients (name, slug, inci_name, description, origin, function, cas_number, ewg_score, is_restricted, is_banned)
		VALUES (:name, :slug, :inci_name, :description, :origin, :function, :cas_number, :ewg_score, :is_restricted, :is_banned)
	`
	_, err := r.DB.NamedExecContext(ctx, query, ingredient)
	if err != nil {
		logger.Error("ProductRepository:PrivateCreateIngredient:Error: ", err)
		return err
	}
	return nil
}

func (r *ProductRepository) PrivateGetIngredients(ctx context.Context, params params.QueryParams) (*entity.PaginatedIngredientEntity, error) {
	// Tính offset cho pagination
	offset := (params.PageNumber - 1) * params.PageSize

	// Base query để lấy ingredients
	baseQuery := `FROM ingredients i`

	// Thêm điều kiện search nếu có
	var whereClause string
	var args []interface{}

	conditions := []string{}
	argIndex := 1

	if params.Search != "" {
		conditions = append(conditions, fmt.Sprintf("(i.name ILIKE $%d OR i.inci_name ILIKE $%d)", argIndex, argIndex))
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
		logger.Error("ProductRepository:PrivateGetIngredients - Count", err)
		return nil, err
	}

	// Query để lấy data với pagination
	dataQuery := `
		SELECT 
			i.id, 
			i.name, 
			i.slug, 
			i.inci_name,
			i.description, 
			i.origin,
			i.function,
			i.cas_number,
			i.ewg_score,
			i.is_restricted,
			i.is_banned,
			i.created_at, 
			i.updated_at
	` + baseQuery + whereClause + `
		ORDER BY i.name ASC, i.created_at DESC
		LIMIT $` + fmt.Sprintf("%d", argIndex) + ` OFFSET $` + fmt.Sprintf("%d", argIndex+1)

	// Thêm params cho pagination
	args = append(args, params.PageSize, offset)

	var ingredients []entity.Ingredient
	err = r.DB.SelectContext(ctx, &ingredients, dataQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:PrivateGetIngredients - Select", err)
		return nil, err
	}

	// Tạo response pagination
	response := &entity.PaginatedIngredientEntity{
		Items:      ingredients,
		TotalItems: totalItems,
		PageNumber: params.PageNumber,
		PageSize:   params.PageSize,
	}

	return response, nil
}

func (r *ProductRepository) PrivateGetIngredientById(ctx context.Context, id uuid.UUID) (*entity.Ingredient, error) {
	query := `
		SELECT 
			i.id, 
			i.name, 
			i.slug, 
			i.inci_name,
			i.description, 
			i.origin,
			i.function,
			i.cas_number,
			i.ewg_score,
			i.is_restricted,
			i.is_banned,
			i.created_at, 
			i.updated_at
		FROM ingredients i
		WHERE i.id = $1
	`
	var ingredient entity.Ingredient
	err := r.DB.GetContext(ctx, &ingredient, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:PrivateGetIngredientById - Select", err)
		return nil, err
	}
	return &ingredient, nil
}

func (r *ProductRepository) PrivateUpdateIngredient(ctx context.Context, ingredient *entity.Ingredient, id uuid.UUID) error {
	query := `
		UPDATE ingredients
		SET name = :name, slug = :slug, inci_name = :inci_name, description = :description, origin = :origin, 
		    function = :function, cas_number = :cas_number, ewg_score = :ewg_score, is_restricted = :is_restricted, 
		    is_banned = :is_banned 
		WHERE id = :id
	`

	// Set ID vào brand struct để sử dụng trong named query
	ingredient.ID = id

	result, err := r.DB.NamedExecContext(ctx, query, ingredient)
	if err != nil {
		logger.Error("ProductRepository:PrivateUpdateIngredient", err)
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

func (r *ProductRepository) PrivateDeleteIngredient(ctx context.Context, id uuid.UUID) error {
	query := `
		DELETE FROM ingredients
		WHERE id = $1
	`

	err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		logger.Error("ProductRepository:PrivateDeleteIngredient:Error: ", err)
		return err
	}

	return nil
}
