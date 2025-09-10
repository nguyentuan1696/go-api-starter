package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"go-api-starter/core/logger"
	"go-api-starter/core/params"
	"go-api-starter/modules/auth/entity"

	"github.com/google/uuid"
)

func (r *AuthRepository) PrivateCreatePermission(ctx context.Context, permission *entity.Permission) error {
	query := `
		INSERT INTO permissions (name, slug, resource, action, description)
		VALUES (:name, :slug, :resource, :action, :description)
	`
	result, err := r.DB.NamedExecContext(ctx, query, permission)
	if err != nil {
		logger.Error("AuthRepository:CreatePermission:Error %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("AuthRepository:CreatePermission:Error %v", err)
		return err
	}

	if rowsAffected == 0 {
		logger.Error("AuthRepository:CreatePermission:Error %v", err)
		return err
	}

	return nil
}

func (r *AuthRepository) PrivateGetPermissions(ctx context.Context, params params.QueryParams) (*entity.PaginatedPermissionEntity, error) {
	// Tính offset cho pagination
	offset := (params.PageNumber - 1) * params.PageSize

	// Base query để lấy permissions
	baseQuery := `FROM permissions p`

	// Thêm điều kiện search nếu có
	var whereClause string
	var args []interface{}

	conditions := []string{}
	argIndex := 1

	if params.Search != "" {
		conditions = append(conditions, fmt.Sprintf("(p.name ILIKE $%d OR p.slug ILIKE $%d OR p.resource ILIKE $%d OR p.action ILIKE $%d)", argIndex, argIndex, argIndex, argIndex))
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
		logger.Error("AuthRepository:PrivateGetPermissions - Count", err)
		return nil, err
	}

	// Query để lấy data với pagination
	dataQuery := `
		SELECT 
			p.id, 
			p.name, 
			p.slug, 
			p.resource,
			p.action,
			p.description,
			p.is_system,
			p.created_at, 
			p.updated_at
	` + baseQuery + whereClause + `
		ORDER BY p.name ASC, p.created_at DESC
		OFFSET $` + fmt.Sprintf("%d", argIndex) + ` ROWS FETCH NEXT $` + fmt.Sprintf("%d", argIndex+1) + ` ROWS ONLY`

	// Thêm params cho pagination
	args = append(args, offset, params.PageSize)

	var permissions []entity.Permission
	err = r.DB.SelectContext(ctx, &permissions, dataQuery, args...)
	if err != nil {
		logger.Error("AuthRepository:PrivateGetPermissions - Select", err)
		return nil, err
	}

	// Tạo response pagination
	response := &entity.PaginatedPermissionEntity{
		Items:      permissions,
		TotalItems: totalItems,
		PageNumber: params.PageNumber,
		PageSize:   params.PageSize,
	}

	return response, nil
}

func (r *AuthRepository) PrivateGetPermissionByID(ctx context.Context, id uuid.UUID) (*entity.Permission, error) {
	var permission entity.Permission
	query := `SELECT * FROM permissions WHERE id = $1`
	err := r.DB.GetContext(ctx, &permission, query, id)
	if err != nil {
		logger.Error("AuthRepository:PrivateGetPermissionByID - Select", err)
		return nil, err
	}
	return &permission, nil
}

func (r *AuthRepository) PrivateUpdatePermission(ctx context.Context, id uuid.UUID, permission *entity.Permission) error {
	query := `
		UPDATE permissions
		SET name = :name, slug = :slug, resource = :resource, action = :action, description = :description, is_system = :is_system
		WHERE id = :id
	`

	permission.ID = id

	result, err := r.DB.NamedExecContext(ctx, query, permission)
	if err != nil {
		logger.Error("AuthRepository:PrivateUpdatePermission:Error %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("AuthRepository:PrivateUpdatePermission:Error %v", err)
		return err
	}

	if rowsAffected == 0 {
		return errors.New("permission not found")
	}

	return nil
}

func (r *AuthRepository) PrivateDeletePermission(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM permissions WHERE id = $1`
	err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		logger.Error("AuthRepository:PrivateDeletePermission:Error %v", err)
		return err
	}
	return nil
}
