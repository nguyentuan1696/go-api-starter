package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"go-api-starter/core/constants"
	"go-api-starter/core/logger"
	"go-api-starter/core/params"
	"go-api-starter/modules/auth/entity"

	"github.com/google/uuid"
)

func (r *AuthRepository) PrivateCreateRole(ctx context.Context, role *entity.Role) error {
	query := `
		INSERT INTO roles (name, slug, description)
		VALUES (:name, :slug, :description)
	`
	result, err := r.DB.NamedExecContext(ctx, query, role)
	if err != nil {
		logger.Error("AuthRepository:CreateRole:Error %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("AuthRepository:CreateRole:Error %v", err)
		return err
	}

	if rowsAffected == 0 {
		logger.Error("AuthRepository:CreateRole:Error %v", err)
		return err
	}

	return nil
}

func (r *AuthRepository) PrivateGetRoles(ctx context.Context, params params.QueryParams) (*entity.PaginatedRoleEntity, error) {
	// Tính offset cho pagination
	offset := (params.PageNumber - 1) * params.PageSize

	// Base query để lấy roles
	baseQuery := `FROM roles r`

	// Thêm điều kiện search nếu có
	var whereClause string
	var args []interface{}

	conditions := []string{}
	argIndex := 1

	if params.Search != "" {
		conditions = append(conditions, fmt.Sprintf("(r.name ILIKE $%d OR r.slug ILIKE $%d)", argIndex, argIndex))
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
		logger.Error("AuthRepository:PrivateGetRoles - Count", err)
		return nil, err
	}

	// Query để lấy data với pagination
	dataQuery := `
		SELECT 
			r.id, 
			r.name, 
			r.slug, 
			r.description, 
			r.is_system,
			r.is_active,
			r.created_at, 
			r.updated_at
	` + baseQuery + whereClause + `
		ORDER BY r.name ASC, r.created_at DESC
		OFFSET $` + fmt.Sprintf("%d", argIndex) + ` ROWS FETCH NEXT $` + fmt.Sprintf("%d", argIndex+1) + ` ROWS ONLY`

	// Thêm params cho pagination
	args = append(args, offset, params.PageSize)

	var roles []entity.Role
	err = r.DB.SelectContext(ctx, &roles, dataQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("AuthRepository:PrivateGetRoles - Select", err)
		return nil, err
	}

	// Tạo response pagination
	response := &entity.PaginatedRoleEntity{
		Items:      roles,
		TotalItems: totalItems,
		PageNumber: params.PageNumber,
		PageSize:   params.PageSize,
	}

	return response, nil
}

func (r *AuthRepository) PrivateGetRoleByID(ctx context.Context, id uuid.UUID) (*entity.Role, error) {
	var role entity.Role
	query := `SELECT * FROM roles WHERE id = $1`
	err := r.DB.GetContext(ctx, &role, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("AuthRepository:PrivateGetRoleByID - Select", err)
		return nil, err
	}
	return &role, nil
}

func (r *AuthRepository) PrivateUpdateRole(ctx context.Context, id uuid.UUID, role *entity.Role) error {
	query := `
		UPDATE roles
		SET name = :name, slug = :slug, description = :description, is_system = :is_system, is_active = :is_active
		WHERE id = :id
	`

	role.ID = id

	result, err := r.DB.NamedExecContext(ctx, query, role)
	if err != nil {
		logger.Error("AuthRepository:PrivateUpdateRole:Error %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("AuthRepository:PrivateUpdateRole:Error %v", err)
		return err
	}

	if rowsAffected == constants.DefaultZeroValue {
		return errors.New("role not found")
	}

	return nil
}

func (r *AuthRepository) PrivateDeleteRole(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM roles WHERE id = $1`
	err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		logger.Error("AuthRepository:PrivateDeleteRole:Error %v", err)
		return err
	}
	return nil
}
