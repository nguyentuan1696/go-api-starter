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

func (r *ProductRepository) PublicCreateWishlist(ctx context.Context, wishlist *entity.Wishlist) error {
	query := `
		INSERT INTO wishlists (user_id, product_id, added_at)
		VALUES (:user_id, :product_id, :added_at)
	`
	_, err := r.DB.NamedExecContext(ctx, query, wishlist)
	if err != nil {
		logger.Error("ProductRepository:PrivateCreateWishlist:Error: ", err)
		return err
	}
	return nil
}

func (r *ProductRepository) PublicGetWishlists(ctx context.Context, params params.QueryParams) (*entity.PaginatedWishlistEntity, error) {
	// Tính offset cho pagination
	offset := (params.PageNumber - 1) * params.PageSize

	// Base query để lấy wishlists
	baseQuery := `FROM wishlists w`

	// Thêm điều kiện search nếu có
	var whereClause string
	var args []interface{}

	conditions := []string{}
	argIndex := 1

	if params.Search != "" {
		conditions = append(conditions, fmt.Sprintf("(w.product_id ILIKE $%d)", argIndex))
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
		logger.Error("ProductRepository:PrivateGetWishlists - Count", err)
		return nil, err
	}

	// Query để lấy data với pagination
	dataQuery := `
		SELECT 
			w.id, 
			w.user_id, 
			w.product_id, 
			w.added_at
	` + baseQuery + whereClause + `
		ORDER BY w.added_at DESC
		OFFSET $` + fmt.Sprintf("%d", argIndex+1) + ` ROWS
		FETCH NEXT $` + fmt.Sprintf("%d", argIndex) + ` ROWS ONLY`

	// Thêm params cho pagination
	args = append(args, params.PageSize, offset)

	var wishlists []entity.Wishlist
	err = r.DB.SelectContext(ctx, &wishlists, dataQuery, args...)
	if err != nil {
		logger.Error("ProductRepository:PrivateGetWishlists - Select", err)
		return nil, err
	}

	// Tạo response pagination
	response := &entity.PaginatedWishlistEntity{
		Items:      wishlists,
		TotalItems: totalItems,
		PageNumber: params.PageNumber,
		PageSize:   params.PageSize,
	}

	return response, nil
}

func (r *ProductRepository) PublicGetWishlistById(ctx context.Context, id string) (*entity.Wishlist, error) {
	query := `
		SELECT 
			w.id, 
			w.user_id, 
			w.product_id, 
			w.added_at
		FROM wishlists w
		WHERE w.id = $1
	`
	var wishlist entity.Wishlist
	err := r.DB.GetContext(ctx, &wishlist, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:PrivateGetWishlistById - Select", err)
		return nil, err
	}
	return &wishlist, nil
}

func (r *ProductRepository) PublicGetWishlistByCustomerAndProduct(ctx context.Context, customerID string, productID string) (*entity.Wishlist, error) {
	query := `
		SELECT 
			w.id, 
			w.user_id, 
			w.product_id, 
			w.added_at
		FROM wishlists w
		WHERE w.user_id = $1 AND w.product_id = $2
	`
	var wishlist entity.Wishlist
	err := r.DB.GetContext(ctx, &wishlist, query, customerID, productID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("ProductRepository:PrivateGetWishlistByCustomerAndProduct - Select", err)
		return nil, err
	}
	return &wishlist, nil
}

func (r *ProductRepository) PublicDeleteWishlist(ctx context.Context, id string) error {
	query := `
		DELETE FROM wishlists
		WHERE id = :id
	`
	wishlistID, err := uuid.Parse(id)
	if err != nil {
		logger.Error("ProductRepository:PrivateDeleteWishlist - Parse UUID", err)
		return err
	}

	wishlist := &entity.Wishlist{ID: wishlistID}
	result, err := r.DB.NamedExecContext(ctx, query, wishlist)
	if err != nil {
		logger.Error("ProductRepository:PrivateDeleteWishlist", err)
		return err
	}

	// Kiểm tra xem có record nào được delete không
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("ProductRepository:PrivateDeleteWishlist - RowsAffected", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("wishlist with id %s not found", id)
	}

	return nil
}
