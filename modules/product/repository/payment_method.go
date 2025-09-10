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

func (r *ProductRepository) PrivateCreatePaymentMethod(ctx context.Context, paymentMethod *entity.PaymentMethod) error {
	query := `
		INSERT INTO payment_methods (name, description, provider, type, is_active, sort_order)
		VALUES (:name, :description, :provider, :type, :is_active, :sort_order)
	`
	_, err := r.DB.NamedExecContext(ctx, query, paymentMethod)
	if err != nil {
		logger.Error("ProductRepository:PrivateCreatePaymentMethod:Error: ", err)
		return err
	}
	return nil
}

func (r *ProductRepository) PrivateGetPaymentMethods(ctx context.Context, params params.QueryParams) (*entity.PaginatedPaymentMethodEntity, error) {
	// Tính offset cho pagination
	offset := (params.PageNumber - 1) * params.PageSize

	// Base query để lấy payment methods
	baseQuery := `FROM payment_methods pm`

	// Thêm điều kiện search nếu có
	var whereClause string
	var args []interface{}

	conditions := []string{}
	argIndex := 1

	if params.Search != "" {
		conditions = append(conditions, fmt.Sprintf("(pm.name ILIKE $%d OR pm.provider ILIKE $%d OR pm.type ILIKE $%d)", argIndex, argIndex, argIndex))
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
		logger.Error("ProductRepository:PrivateGetPaymentMethods - Count", err)
		return nil, err
	}

	// Query để lấy data với pagination
	dataQuery := `
		SELECT 
			pm.id, 
			pm.name, 
			pm.description, 
			pm.provider,
			pm.type,
			pm.is_active,
			pm.sort_order,
			pm.created_at, 
			pm.updated_at
	` + baseQuery + whereClause + `
		ORDER BY pm.sort_order ASC, pm.name ASC, pm.created_at DESC
		LIMIT $` + fmt.Sprintf("%d", argIndex) + ` OFFSET $` + fmt.Sprintf("%d", argIndex+1)

	// Thêm params cho pagination
	args = append(args, params.PageSize, offset)

	var paymentMethods []entity.PaymentMethod
	err = r.DB.SelectContext(ctx, &paymentMethods, dataQuery, args...)
	if err != nil {
		logger.Error("ProductRepository:PrivateGetPaymentMethods - Select", err)
		return nil, err
	}

	// Tạo response pagination
	response := &entity.PaginatedPaymentMethodEntity{
		Items:      paymentMethods,
		TotalItems: totalItems,
		PageNumber: params.PageNumber,
		PageSize:   params.PageSize,
	}

	return response, nil
}

func (r *ProductRepository) PrivateGetPaymentMethodById(ctx context.Context, id int) (*entity.PaymentMethod, error) {
	query := `
		SELECT 
			pm.id, 
			pm.name, 
			pm.description, 
			pm.provider,
			pm.type,
			pm.is_active,
			pm.sort_order,
			pm.created_at, 
			pm.updated_at
		FROM payment_methods pm
		WHERE pm.id = $1
	`
	var paymentMethod entity.PaymentMethod
	err := r.DB.GetContext(ctx, &paymentMethod, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:PrivateGetPaymentMethodById - Select", err)
		return nil, err
	}
	return &paymentMethod, nil
}

func (r *ProductRepository) PrivateUpdatePaymentMethod(ctx context.Context, paymentMethod *entity.PaymentMethod, id int) error {
	query := `
		UPDATE payment_methods
		SET name = :name, description = :description, provider = :provider, type = :type, 
		    is_active = :is_active, sort_order = :sort_order 
		WHERE id = :id
	`

	// Set ID vào paymentMethod struct để sử dụng trong named query
	paymentMethod.ID = id

	result, err := r.DB.NamedExecContext(ctx, query, paymentMethod)
	if err != nil {
		logger.Error("ProductRepository:PrivateUpdatePaymentMethod", err)
		return err
	}

	// Kiểm tra xem có record nào được update không
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("ProductRepository:PrivateUpdatePaymentMethod - RowsAffected", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("payment method with id %v not found", id)
	}

	return nil
}

func (r *ProductRepository) PrivateDeletePaymentMethod(ctx context.Context, id int) error {
	query := `
		DELETE FROM payment_methods
		WHERE id = :id
	`
	paymentMethod := &entity.PaymentMethod{ID: id}
	result, err := r.DB.NamedExecContext(ctx, query, paymentMethod)
	if err != nil {
		logger.Error("ProductRepository:PrivateDeletePaymentMethod", err)
		return err
	}
	// Kiểm tra xem có record nào được delete không
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("ProductRepository:PrivateDeletePaymentMethod - RowsAffected", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("payment method with id %v not found", id)
	}

	return nil
}
