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

// Order methods

func (r *ProductRepository) GetOrderDetailWithItems(ctx context.Context, orderID uuid.UUID) (*entity.OrderDetailWithItems, error) {
	var order entity.OrderDetailWithItems

	query := `
    SELECT 
        o.id AS order_id,
        o.order_number,
        o.customer_id,
        o.customer_name,
        o.customer_email,
        o.customer_phone,
        o.shipping_recipient_name,
        o.shipping_recipient_phone,
        o.shipping_address,
        o.shipping_ward_name,
        o.shipping_district_name,
        o.shipping_province_name,
		o.order_state,
		o.shipping_method_id,
		o.shipping_method_name,
		pm.id AS payment_method_id,
		pm.name AS payment_method_name,
        o.order_state,
        o.payment_status,
		o.notes,
		o.admin_notes,
        o.subtotal,
        o.shipping_cost,
        o.tax_amount,
        o.discount_amount,
        o.total_amount,
        o.ordered_at,
        o.confirmed_at,
        o.shipped_at,
        o.delivered_at,
        o.cancelled_at,
        COALESCE(
            json_agg(
                jsonb_build_object(
                    'order_item_id', oi.id,
                    'product_id', oi.product_id,
                    'product_name', oi.product_name,
                    'product_sku', oi.product_sku,
                    'product_thumbnail', p.thumbnail,
                    'unit_price', oi.unit_price,
                    'quantity', oi.quantity,
                    'total_price', oi.total_price
                )
            ) FILTER (WHERE oi.id IS NOT NULL),
            '[]'
        ) AS items
    FROM orders o
    LEFT JOIN order_items oi ON oi.order_id = o.id
    LEFT JOIN products p ON p.id = oi.product_id
	LEFT JOIN payment_methods pm ON pm.id = o.payment_method_id
    WHERE o.id = $1
    GROUP BY o.id, pm.id
    `

	err := r.DB.GetContext(ctx, &order, query, orderID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:PublicGetOrderDetailWithItems:Error:", err)
		return nil, err
	}

	return &order, nil
}

func (r *ProductRepository) PublicCreateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error) {
	query := `
        INSERT INTO orders (
            order_number, customer_id, customer_email, customer_phone, customer_name,
            shipping_recipient_name, shipping_recipient_phone, 
            shipping_address, shipping_district_name, shipping_ward_name, shipping_province_name,
            order_state, payment_status, subtotal, shipping_cost, tax_amount, discount_amount, total_amount,
            shipping_method_id, shipping_method_name, payment_method_id, payment_method_name,
            coupon_id, coupon_code, coupon_discount_amount, notes, admin_notes
        ) VALUES (
            :order_number, :customer_id, :customer_email, :customer_phone, :customer_name,
            :shipping_recipient_name, :shipping_recipient_phone, :shipping_address,
            :shipping_district_name, :shipping_ward_name, :shipping_province_name,
            :order_state, :payment_status, :subtotal, :shipping_cost, :tax_amount, :discount_amount, :total_amount,
            :shipping_method_id, :shipping_method_name, :payment_method_id, :payment_method_name,
            :coupon_id, :coupon_code, :coupon_discount_amount, :notes, :admin_notes
        )
        RETURNING *
    `
	rows, err := r.DB.NamedQueryContext(ctx, query, order)
	if err != nil {
		logger.Error("ProductRepository:PublicCreateOrder:Error:", err)
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var inserted entity.Order
		if err := rows.StructScan(&inserted); err != nil {
			logger.Error("ProductRepository:PublicCreateOrder:StructScan", err)
			return nil, err
		}
		return &inserted, nil
	}
	return nil, fmt.Errorf("insert order failed")
}

func (r *ProductRepository) PrivateCreateOrderItems(ctx context.Context, orderItems []*entity.OrderItem) error {
	query := `
		INSERT INTO order_items (
			order_id, product_id, product_name, product_sku, unit_price, quantity, total_price
		) VALUES (
			:order_id, :product_id, :name, :sku, :unit_price, :quantity, :total_price
		)
	`
	_, err := r.DB.NamedExecContext(ctx, query, orderItems)
	return err
}

func (r *ProductRepository) PrivateGetOrders(ctx context.Context, params params.QueryParams) (*entity.PaginatedOrderEntity, error) {
	// Tính offset cho pagination
	offset := (params.PageNumber - 1) * params.PageSize

	// Base query để lấy orders
	baseQuery := `FROM orders o`

	// Thêm điều kiện search nếu có
	var whereClause string
	var args []interface{}

	conditions := []string{}
	argIndex := 1

	if params.Search != "" {
		conditions = append(conditions, fmt.Sprintf("(o.order_number ILIKE $%d OR o.customer_phone ILIKE $%d OR o.customer_name ILIKE $%d)", argIndex, argIndex, argIndex))
		args = append(args, "%"+params.Search+"%")
		argIndex++
	}

	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}

	// Query để đếm tổng số records
	countQuery := "SELECT COUNT(*) " + baseQuery + whereClause

	var totalItems int
	if err := r.DB.GetContext(ctx, &totalItems, countQuery, args...); err != nil {
		logger.Error("ProductRepository:GetOrders - Count", err)
		return nil, err
	}

	// Query để lấy data với pagination
	dataQuery := `
		SELECT o.*
	` + baseQuery + whereClause + `
		ORDER BY o.created_at DESC
		LIMIT $` + fmt.Sprintf("%d", argIndex) + ` OFFSET $` + fmt.Sprintf("%d", argIndex+1)

	// Thêm params cho pagination
	args = append(args, params.PageSize, offset)

	var orders []entity.Order
	if err := r.DB.SelectContext(ctx, &orders, dataQuery, args...); err != nil {
		logger.Error("ProductRepository:GetOrders - Select", err)
		return nil, err
	}

	// Tạo response pagination
	response := &entity.PaginatedOrderEntity{
		Items:      orders,
		TotalItems: totalItems,
		PageNumber: params.PageNumber,
		PageSize:   params.PageSize,
	}

	return response, nil
}

func (r *ProductRepository) PrivateGetOrderById(ctx context.Context, id uuid.UUID) (*entity.Order, error) {
	query := `
		SELECT 
			o.*
		FROM orders o
		WHERE o.id = $1
	`
	var order entity.Order
	err := r.DB.GetContext(ctx, &order, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:PrivateGetOrderById - Select", err)
		return nil, err
	}
	return &order, nil
}

func (r *ProductRepository) PrivateGetOrderWithItems(ctx context.Context, id uuid.UUID) (*entity.OrderWithItems, error) {
	// Lấy thông tin order
	order, err := r.PrivateGetOrderById(ctx, id)
	if err != nil {
		return nil, err
	}

	// Lấy danh sách order items
	itemsQuery := `
		SELECT 
			oi.id, oi.order_id, oi.product_id, oi.product_name, oi.product_sku, oi.product_thumbnail,
			oi.unit_price, oi.quantity, oi.total_price, oi.created_at, oi.updated_at
		FROM order_items oi
		WHERE oi.order_id = $1
		ORDER BY oi.created_at ASC
	`

	var items []entity.OrderItem
	err = r.DB.SelectContext(ctx, &items, itemsQuery, id)
	if err != nil {
		logger.Error("ProductRepository:PrivateGetOrderWithItems - Select Items", err)
		return nil, err
	}

	// Tạo response
	response := &entity.OrderWithItems{
		Order: *order,
		Items: items,
	}

	return response, nil
}

func (r *ProductRepository) PrivateUpdateOrder(ctx context.Context, order *entity.Order, id uuid.UUID) error {
	query := `
		UPDATE orders
		SET customer_email = :customer_email, customer_phone = :customer_phone, customer_name = :customer_name,
		    shipping_recipient_name = :shipping_recipient_name, shipping_recipient_phone = :shipping_recipient_phone,
		    shipping_address = :shipping_address, shipping_ward_name = :shipping_ward_name, 
		    shipping_district_name = :shipping_district_name, shipping_province_name = :shipping_province_name,
		    order_state = :order_state, payment_status = :payment_status,
		    subtotal = :subtotal, shipping_cost = :shipping_cost, tax_amount = :tax_amount,
		    discount_amount = :discount_amount, total_amount = :total_amount,
		    shipping_method_id = :shipping_method_id, shipping_method_name = :shipping_method_name,
		    payment_method_id = :payment_method_id, payment_method_name = :payment_method_name,
		    coupon_id = :coupon_id, coupon_code = :coupon_code, coupon_discount_amount = :coupon_discount_amount,
		    notes = :notes, admin_notes = :admin_notes, confirmed_at = :confirmed_at, shipped_at = :shipped_at,
		    delivered_at = :delivered_at, cancelled_at = :cancelled_at
		WHERE id = :id
	`

	order.ID = id

	result, err := r.DB.NamedExecContext(ctx, query, order)
	if err != nil {
		logger.Error("ProductRepository:PrivateUpdateOrder", err)
		return err
	}

	// Kiểm tra xem có record nào được update không
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("ProductRepository:PrivateUpdateOrder - RowsAffected", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("order with id %s not found", id)
	}

	return nil
}

func (r *ProductRepository) PrivateDeleteOrder(ctx context.Context, id uuid.UUID) error {
	query := `
		DELETE FROM orders
		WHERE id = :id
	`

	order := &entity.Order{ID: id}
	result, err := r.DB.NamedExecContext(ctx, query, order)
	if err != nil {
		logger.Error("ProductRepository:PrivateDeleteOrder", err)
		return err
	}

	// Kiểm tra xem có record nào được delete không
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("ProductRepository:PrivateDeleteOrder - RowsAffected", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("order with id %s not found", id)
	}

	return nil
}

// OrderItem methods
func (r *ProductRepository) PrivateCreateOrderItem(ctx context.Context, orderItem *entity.OrderItem) error {
	query := `
		INSERT INTO order_items (
			id, order_id, product_id, product_name, product_sku, product_thumbnail,
			unit_price, quantity, total_price
		) VALUES (
			:id, :order_id, :product_id, :product_name, :product_sku, :product_thumbnail,
			:unit_price, :quantity, :total_price
		)
	`
	_, err := r.DB.NamedExecContext(ctx, query, orderItem)
	if err != nil {
		logger.Error("ProductRepository:PrivateCreateOrderItem:Error: ", err)
		return err
	}
	return nil
}

func (r *ProductRepository) PrivateGetOrderItems(ctx context.Context, orderID uuid.UUID, params params.QueryParams) (*entity.PaginatedOrderItemEntity, error) {
	// Tính offset cho pagination
	offset := (params.PageNumber - 1) * params.PageSize

	// Base query để lấy order items
	baseQuery := `FROM order_items oi WHERE oi.order_id = $1`

	// Thêm điều kiện search nếu có
	var whereClause string
	var args []interface{}
	args = append(args, orderID) // orderID là tham số đầu tiên

	argIndex := 2 // Bắt đầu từ $2 vì $1 đã dùng cho orderID

	if params.Search != "" {
		whereClause = fmt.Sprintf(" AND (oi.product_name ILIKE $%d OR oi.product_sku ILIKE $%d)", argIndex, argIndex)
		args = append(args, "%"+params.Search+"%")
		argIndex++
	}

	// Query để đếm tổng số records
	countQuery := "SELECT COUNT(*) " + baseQuery + whereClause

	var totalItems int
	err := r.DB.GetContext(ctx, &totalItems, countQuery, args...)
	if err != nil {
		logger.Error("ProductRepository:PrivateGetOrderItems - Count", err)
		return nil, err
	}

	// Query để lấy data với pagination
	dataQuery := `
		SELECT 
			oi.id, oi.order_id, oi.product_id, oi.product_name, oi.product_sku, oi.product_thumbnail,
			oi.unit_price, oi.quantity, oi.total_price, oi.created_at, oi.updated_at
	` + baseQuery + whereClause + `
		ORDER BY oi.created_at ASC
		OFFSET $` + fmt.Sprintf("%d", argIndex+1) + ` ROWS
		FETCH NEXT $` + fmt.Sprintf("%d", argIndex) + ` ROWS ONLY`

	// Thêm params cho pagination
	args = append(args, params.PageSize, offset)

	var orderItems []entity.OrderItem
	err = r.DB.SelectContext(ctx, &orderItems, dataQuery, args...)
	if err != nil {
		logger.Error("ProductRepository:PrivateGetOrderItems - Select", err)
		return nil, err
	}

	// Tạo response pagination
	response := &entity.PaginatedOrderItemEntity{
		Items:      orderItems,
		TotalItems: totalItems,
		PageNumber: params.PageNumber,
		PageSize:   params.PageSize,
	}

	return response, nil
}

func (r *ProductRepository) PrivateGetOrderItemById(ctx context.Context, id uuid.UUID) (*entity.OrderItem, error) {
	query := `
		SELECT 
			oi.id, oi.order_id, oi.product_id, oi.product_name, oi.product_sku, oi.product_thumbnail,
			oi.unit_price, oi.quantity, oi.total_price, oi.created_at, oi.updated_at
		FROM order_items oi
		WHERE oi.id = $1
	`
	var orderItem entity.OrderItem
	err := r.DB.GetContext(ctx, &orderItem, query, id)
	if err != nil {
		logger.Error("ProductRepository:PrivateGetOrderItemById - Select", err)
		return nil, err
	}
	return &orderItem, nil
}

func (r *ProductRepository) PrivateUpdateOrderItem(ctx context.Context, orderItem *entity.OrderItem, id uuid.UUID) error {
	query := `
		UPDATE order_items
		SET product_name = :product_name, product_sku = :product_sku, product_thumbnail = :product_thumbnail,
		    unit_price = :unit_price, quantity = :quantity, total_price = :total_price
		WHERE id = :id
	`

	orderItem.ID = id

	result, err := r.DB.NamedExecContext(ctx, query, orderItem)
	if err != nil {
		logger.Error("ProductRepository:PrivateUpdateOrderItem", err)
		return err
	}

	// Kiểm tra xem có record nào được update không
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("ProductRepository:PrivateUpdateOrderItem - RowsAffected", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("order item with id %s not found", id)
	}

	return nil
}

func (r *ProductRepository) PrivateDeleteOrderItem(ctx context.Context, id uuid.UUID) error {
	query := `
		DELETE FROM order_items
		WHERE id = :id
	`

	orderItem := &entity.OrderItem{ID: id}
	result, err := r.DB.NamedExecContext(ctx, query, orderItem)
	if err != nil {
		logger.Error("ProductRepository:PrivateDeleteOrderItem", err)
		return err
	}

	// Kiểm tra xem có record nào được delete không
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("ProductRepository:PrivateDeleteOrderItem - RowsAffected", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("order item with id %s not found", id)
	}

	return nil
}
