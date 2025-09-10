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

func (r *ProductRepository) PublicGetProductDetailBySlug(ctx context.Context, slug string, fields []string) (*entity.ProductDetailEntity, error) {
	// Default fields if none provided
	if len(fields) == 0 {
		fields = []string{"id", "name", "slug", "description", "product_description", "usage_instructions",
			"brand_id", "category_id", "price", "original_price", "sku", "barcode", "registration_number",
			"weight", "volume", "thumbnail", "images", "is_active", "is_featured", "stock_quantity",
			"min_stock_level", "created_at", "updated_at", "ingredients"}
	}

	// Build select clause with requested fields
	selectClause := "p." + strings.Join(fields, ", p.")

	query := fmt.Sprintf(`
		SELECT 
			%s,
			b.name as brand_name,
			c.name as category_name
		FROM products p
		LEFT JOIN brands b ON p.brand_id = b.id
		LEFT JOIN categories c ON p.category_id = c.id
		WHERE p.slug = $1 AND p.is_active = true
	`, selectClause)

	var productDetail entity.ProductDetailEntity
	err := r.DB.GetContext(ctx, &productDetail, query, slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:PublicGetProductDetail - Select", err)
		return nil, err
	}
	return &productDetail, nil
}

func (r *ProductRepository) PublicGetProductDetail(ctx context.Context, id uuid.UUID, fields []string) (*entity.ProductDetailEntity, error) {
	// Default fields if none provided
	if len(fields) == 0 {
		fields = []string{"id", "name", "slug", "description", "product_description", "usage_instructions",
			"brand_id", "category_id", "price", "original_price", "sku", "barcode", "registration_number",
			"weight", "volume", "thumbnail", "images", "is_active", "is_featured", "stock_quantity",
			"min_stock_level", "created_at", "updated_at", "ingredients"}
	}

	// Build select clause with requested fields
	selectClause := "p." + strings.Join(fields, ", p.")

	query := fmt.Sprintf(`
		SELECT 
			%s,
			b.name as brand_name,
			c.name as category_name
		FROM products p
		LEFT JOIN brands b ON p.brand_id = b.id
		LEFT JOIN categories c ON p.category_id = c.id
		WHERE p.id = $1 AND p.is_active = true
	`, selectClause)

	var productDetail entity.ProductDetailEntity
	err := r.DB.GetContext(ctx, &productDetail, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:PublicGetProductDetail - Select", err)
		return nil, err
	}
	return &productDetail, nil
}

func (r *ProductRepository) PrivateCreateProduct(ctx context.Context, product *entity.Product) error {
	query := `
		INSERT INTO products (name, slug, description, product_description, usage_instructions, brand_id, category_id, 
		                     price, original_price, sku, barcode, registration_number, weight, volume, thumbnail, images, 
		                     is_active, is_featured, stock_quantity, min_stock_level, ingredients)
		VALUES (:name, :slug, :description, :product_description, :usage_instructions, :brand_id, :category_id, 
		        :price, :original_price, :sku, :barcode, :registration_number, :weight, :volume, :thumbnail, :images, 
		        :is_active, :is_featured, :stock_quantity, :min_stock_level, :ingredients)
	`
	_, err := r.DB.NamedExecContext(ctx, query, product)
	if err != nil {
		logger.Error("ProductRepository:PrivateCreateProduct:Error: ", err)
		return err
	}
	return nil
}

func (r *ProductRepository) PrivateGetProducts(ctx context.Context, params params.QueryParams) (*entity.PaginatedProductEntity, error) {
	// Tính offset cho pagination
	offset := (params.PageNumber - 1) * params.PageSize

	// Base query để lấy products
	baseQuery := `FROM products p`

	// Thêm điều kiện search nếu có
	var whereClause string
	var args []interface{}

	conditions := []string{}
	argIndex := 1

	if params.Search != "" {
		conditions = append(conditions, fmt.Sprintf("(p.name ILIKE $%d OR p.slug ILIKE $%d OR p.sku ILIKE $%d)", argIndex, argIndex, argIndex))
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
		logger.Error("ProductRepository:PrivateGetProducts - Count", err)
		return nil, err
	}

	// Query để lấy data với pagination
	dataQuery := `
		SELECT 
			p.id, 
			p.name, 
			p.slug, 
			p.description, 
			p.product_description,
			p.usage_instructions,
			p.brand_id,
			p.category_id,
			p.price,
			p.original_price,
			p.sku,
			p.barcode,
			p.registration_number,
			p.weight,
			p.volume,
			p.thumbnail,
			p.images,
			p.is_active,
			p.is_featured,
			p.stock_quantity,
			p.min_stock_level,
			p.created_at, 
			p.updated_at,
			p.ingredients
	` + baseQuery + whereClause + `
		ORDER BY p.name ASC, p.created_at DESC
		OFFSET $` + fmt.Sprintf("%d", argIndex+1) + ` ROWS FETCH NEXT $` + fmt.Sprintf("%d", argIndex) + ` ROWS ONLY`

	// Thêm params cho pagination
	args = append(args, params.PageSize, offset)

	var products []entity.Product
	err = r.DB.SelectContext(ctx, &products, dataQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:PrivateGetProducts - Select", err)
		return nil, err
	}

	// Tạo response pagination
	response := &entity.PaginatedProductEntity{
		Items:      products,
		TotalItems: totalItems,
		PageNumber: params.PageNumber,
		PageSize:   params.PageSize,
	}

	return response, nil
}

func (r *ProductRepository) PrivateGetProductById(ctx context.Context, id uuid.UUID) (*entity.Product, error) {
	query := `
		SELECT 
			p.id, 
			p.name, 
			p.slug, 
			p.description, 
			p.product_description,
			p.usage_instructions,
			p.brand_id,
			p.category_id,
			p.price,
			p.original_price,
			p.sku,
			p.barcode,
			p.registration_number,
			p.weight,
			p.volume,
			p.thumbnail,
			p.images,
			p.is_active,
			p.is_featured,
			p.stock_quantity,
			p.min_stock_level,
			p.created_at, 
			p.updated_at,
			p.ingredients
		FROM products p
		WHERE p.id = $1
	`
	var product entity.Product
	err := r.DB.GetContext(ctx, &product, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:PrivateGetProductById - Select", err)
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) PrivateUpdateProduct(ctx context.Context, product *entity.Product, id uuid.UUID) error {
	query := `
		UPDATE products
		SET name = :name, description = :description, product_description = :product_description, usage_instructions = :usage_instructions,
		    brand_id = :brand_id, category_id = :category_id, price = :price, original_price = :original_price,
		    sku = :sku, barcode = :barcode, registration_number = :registration_number, weight = :weight, volume = :volume, thumbnail = :thumbnail,
		    images = :images, is_active = :is_active, is_featured = :is_featured, 
		    stock_quantity = :stock_quantity, min_stock_level = :min_stock_level, ingredients = :ingredients
		WHERE id = :id
	`

	// Set ID vào product struct để sử dụng trong named query
	product.ID = id

	result, err := r.DB.NamedExecContext(ctx, query, product)
	if err != nil {
		logger.Error("ProductRepository:PrivateUpdateProduct", err)
		return err
	}

	// Kiểm tra xem có record nào được update không
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("ProductRepository:PrivateUpdateProduct - RowsAffected", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product with id %s not found", id)
	}

	return nil
}

func (r *ProductRepository) PrivateDeleteProduct(ctx context.Context, id uuid.UUID) error {
	query := `
		DELETE FROM products
		WHERE id = $1
	`
	err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		logger.Error("ProductRepository:PrivateDeleteProduct", err)
		return err
	}
	return nil
}
