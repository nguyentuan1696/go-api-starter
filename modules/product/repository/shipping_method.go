package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"go-api-starter/core/logger"
	"go-api-starter/core/params"
	"go-api-starter/modules/product/entity"
)

func (r *ProductRepository) PrivateCreateShippingMethod(ctx context.Context, shippingMethod *entity.ShippingMethod) error {
	query := `
		INSERT INTO shipping_methods (name, description, provider, base_cost, cost_per_kg, 
		                             free_shipping_threshold, estimated_days_min, estimated_days_max, 
		                             is_active, sort_order)
		VALUES (:name, :description, :provider, :base_cost, :cost_per_kg, 
		        :free_shipping_threshold, :estimated_days_min, :estimated_days_max, 
		        :is_active, :sort_order)
	`
	_, err := r.DB.NamedExecContext(ctx, query, shippingMethod)
	if err != nil {
		logger.Error("ProductRepository:PrivateCreateShippingMethod:Error: ", err)
		return err
	}
	return nil
}

func (r *ProductRepository) PrivateGetShippingMethods(ctx context.Context, params params.QueryParams) (*entity.PaginatedShippingMethodEntity, error) {
	// Tính offset cho pagination
	offset := (params.PageNumber - 1) * params.PageSize

	// Base query để lấy shipping methods
	baseQuery := `FROM shipping_methods sm`

	// Thêm điều kiện search nếu có
	var whereClause string
	var args []interface{}

	conditions := []string{}
	argIndex := 1

	if params.Search != "" {
		conditions = append(conditions, fmt.Sprintf("(sm.name ILIKE $%d OR sm.provider ILIKE $%d)", argIndex, argIndex))
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
		logger.Error("ProductRepository:PrivateGetShippingMethods - Count", err)
		return nil, err
	}

	// Query để lấy data với pagination
	dataQuery := `
		SELECT 
			sm.id, 
			sm.name, 
			sm.description, 
			sm.provider,
			sm.base_cost,
			sm.cost_per_kg,
			sm.free_shipping_threshold,
			sm.estimated_days_min,
			sm.estimated_days_max,
			sm.is_active,
			sm.sort_order,
			sm.created_at, 
			sm.updated_at
	` + baseQuery + whereClause + `
		ORDER BY sm.sort_order ASC, sm.name ASC, sm.created_at DESC
		OFFSET $` + fmt.Sprintf("%d", argIndex+1) + ` ROWS
		FETCH NEXT $` + fmt.Sprintf("%d", argIndex) + ` ROWS ONLY`

	// Thêm params cho pagination
	args = append(args, params.PageSize, offset)

	var shippingMethods []entity.ShippingMethod
	err = r.DB.SelectContext(ctx, &shippingMethods, dataQuery, args...)
	if err != nil {
		logger.Error("ProductRepository:PrivateGetShippingMethods - Select", err)
		return nil, err
	}

	// Tạo response pagination
	response := &entity.PaginatedShippingMethodEntity{
		Items:      shippingMethods,
		TotalItems: totalItems,
		PageNumber: params.PageNumber,
		PageSize:   params.PageSize,
	}

	return response, nil
}

func (r *ProductRepository) PrivateGetShippingMethodById(ctx context.Context, id int) (*entity.ShippingMethod, error) {
	query := `
		SELECT 
			sm.id, 
			sm.name, 
			sm.description, 
			sm.provider,
			sm.base_cost,
			sm.cost_per_kg,
			sm.free_shipping_threshold,
			sm.estimated_days_min,
			sm.estimated_days_max,
			sm.is_active,
			sm.sort_order,
			sm.created_at, 
			sm.updated_at
		FROM shipping_methods sm
		WHERE sm.id = $1
	`
	var shippingMethod entity.ShippingMethod
	err := r.DB.GetContext(ctx, &shippingMethod, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:PrivateGetShippingMethodById - Select", err)
		return nil, err
	}
	return &shippingMethod, nil
}

func (r *ProductRepository) PrivateUpdateShippingMethod(ctx context.Context, shippingMethod *entity.ShippingMethod, id int) error {
	query := `
		UPDATE shipping_methods
		SET name = :name, description = :description, provider = :provider, 
		    base_cost = :base_cost, cost_per_kg = :cost_per_kg, 
		    free_shipping_threshold = :free_shipping_threshold,
		    estimated_days_min = :estimated_days_min, estimated_days_max = :estimated_days_max,
		    is_active = :is_active, sort_order = :sort_order
		WHERE id = :id
	`

	// Set ID vào shipping method struct để sử dụng trong named query
	shippingMethod.ID = id

	result, err := r.DB.NamedExecContext(ctx, query, shippingMethod)
	if err != nil {
		logger.Error("ProductRepository:PrivateUpdateShippingMethod", err)
		return err
	}

	// Kiểm tra xem có record nào được update không
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("ProductRepository:PrivateUpdateShippingMethod - RowsAffected", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("shipping method with id %v not found", id)
	}

	return nil
}

func (r *ProductRepository) PrivateDeleteShippingMethod(ctx context.Context, id int) error {
	query := `
		DELETE FROM shipping_methods
		WHERE id = :id
	`
	shippingMethod := &entity.ShippingMethod{ID: id}
	result, err := r.DB.NamedExecContext(ctx, query, shippingMethod)
	if err != nil {
		logger.Error("ProductRepository:PrivateDeleteShippingMethod", err)
		return err
	}
	// Kiểm tra xem có record nào được delete không
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("ProductRepository:PrivateDeleteShippingMethod - RowsAffected", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("shipping method with id %v not found", id)
	}

	return nil
}
