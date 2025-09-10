package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"go-api-starter/core/logger"
	"go-api-starter/core/params"
	"go-api-starter/modules/auth/entity"

	"github.com/google/uuid"
)

func (r *AuthRepository) PrivateUpdatePasswordUser(ctx context.Context, userID uuid.UUID, password string) error {
	query := `UPDATE users SET password = $1 WHERE id = $2`
	err := r.DB.ExecContext(ctx, query, password, userID)
	if err != nil {
		logger.Error("AuthRepository:PrivateUpdatePasswordUser - Exec", err)
		return err
	}
	return nil
}

func (r *AuthRepository) PrivateGetUsers(ctx context.Context, params params.QueryParams) (*entity.PaginatedUserEntity, error) {
	// Tính offset cho pagination
	offset := (params.PageNumber - 1) * params.PageSize

	// Base query để lấy users
	baseQuery := `FROM users u`

	// Thêm điều kiện search nếu có
	var whereClause string
	var args []interface{}

	conditions := []string{}
	argIndex := 1

	if params.Search != "" {
		conditions = append(conditions, fmt.Sprintf("(u.email ILIKE $%d OR u.phone ILIKE $%d OR u.username ILIKE $%d)", argIndex, argIndex, argIndex))
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
		logger.Error("AuthRepository:PrivateGetUsers - Count", err)
		return nil, err
	}

	// Query để lấy data với pagination
	dataQuery := `
		SELECT 
			u.id, 
			u.email, 
			u.phone, 
			u.username, 
			u.password,
			u.email_verified_at,
			u.phone_verified_at,
			u.locked_until,
			u.is_active,
			u.created_at, 
			u.updated_at
	` + baseQuery + whereClause + `
		ORDER BY u.created_at DESC
		OFFSET $` + fmt.Sprintf("%d", argIndex+1) + ` ROWS FETCH NEXT $` + fmt.Sprintf("%d", argIndex) + ` ROWS ONLY`

	// Thêm params cho pagination
	args = append(args, params.PageSize, offset)

	var users []entity.User
	err = r.DB.SelectContext(ctx, &users, dataQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("AuthRepository:PrivateGetUsers - Select", err)
		return nil, err
	}

	// Tạo response pagination
	response := &entity.PaginatedUserEntity{
		Items:      users,
		TotalItems: totalItems,
		PageNumber: params.PageNumber,
		PageSize:   params.PageSize,
	}

	return response, nil
}

func (r *AuthRepository) PrivateGetUser(ctx context.Context, id uuid.UUID) (*entity.UserDetail, error) {
	query := `
		SELECT
			u.id                     AS id,
			u.email,
			u.phone,
			u.username,
			u.is_active,
			u.created_at,
			up.display_name,
			up.full_name,
			up.avatar,
			up.date_of_birth,
			up.gender,
			string_agg(r.name, ', ') AS roles
		FROM users u
		LEFT JOIN user_profiles up
			ON u.id = up.user_id
		LEFT JOIN user_roles ur
			ON u.id = ur.user_id AND ur.is_active = true
		LEFT JOIN roles r
			ON ur.role_id = r.id AND r.is_active = true
		WHERE u.id = $1
		GROUP BY
			u.id, u.email, u.phone, u.username, u.is_active, u.created_at,
			up.display_name, up.full_name, up.avatar, up.date_of_birth, up.gender;
	`

	var userDetail entity.UserDetail
	err := r.DB.GetContext(ctx, &userDetail, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Info("AuthRepository:PrivateGetUser:UserNotFound:", id)
			return nil, nil // Return nil instead of error for not found case
		}
		logger.Error("AuthRepository:PrivateGetUser:Error:", err)
		return nil, err
	}

	return &userDetail, nil
}

func (r *AuthRepository) PrivateUpdateUser(ctx context.Context, user *entity.User, userId uuid.UUID) error {
	query := `
		UPDATE users 
		SET 
			email = :email,
			phone = :phone,
			username = :username,
			password = :password,
			email_verified_at = :email_verified_at,
			phone_verified_at = :phone_verified_at,
			locked_until = :locked_until,
			is_active = :is_active,
			updated_at = :updated_at
		WHERE id = :id
	`
	result, err := r.DB.NamedExecContext(ctx, query, user)
	if err != nil {
		logger.Error("AuthRepository:PrivateUpdateUser:Error:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("AuthRepository:PrivateUpdateUser:RowsAffected:Error:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

func (r *AuthRepository) UpdateUserProfile(ctx context.Context, userProfile *entity.UserProfile) error {
	query := `
		UPDATE user_profiles 
		SET 
			first_name = :first_name,
			last_name = :last_name,
			display_name = :display_name,
			avatar = :avatar,
			date_of_birth = :date_of_birth,
			gender = :gender,
			updated_at = :updated_at
		WHERE user_id = :user_id
	`
	result, err := r.DB.NamedExecContext(ctx, query, userProfile)
	if err != nil {
		logger.Error("AuthRepository:UpdateUserProfile:Error:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("AuthRepository:UpdateUserProfile:RowsAffected:Error:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user profile not found")
	}

	return nil
}

func (r *AuthRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	query := `
		UPDATE users 
		SET 
			email = :email,
			phone = :phone,
			username = :username,
			password = :password,
			email_verified_at = :email_verified_at,
			phone_verified_at = :phone_verified_at,
			locked_until = :locked_until,
			is_active = :is_active,
			updated_at = :updated_at
		WHERE id = :id
	`
	result, err := r.DB.NamedExecContext(ctx, query, user)
	if err != nil {
		logger.Error("AuthRepository:UpdateUser:Error:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("AuthRepository:UpdateUser:RowsAffected:Error:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

func (r *AuthRepository) GetUserByIdentifier(ctx context.Context, identifier string) (*entity.User, error) {
	var user entity.User
	query := `SELECT * FROM users WHERE phone = $1 OR email = $1 OR id::text = $1 OR username = $1`
	err := r.DB.GetContext(ctx, &user, query, identifier)
	if err != nil {
		if err == sql.ErrNoRows {
			// User not found is not an error, return nil
			return nil, nil
		}
		logger.Error("AuthRepository:GetUserByIdentifier:Error:", err)
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	query := `
		INSERT INTO users (email, phone, username, password)
		VALUES (:email, :phone, :username, :password)
		RETURNING *
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, user)
	if err != nil {
		logger.Error("AuthRepository:CreateUser:Error:", err)
		return nil, nil
	}
	defer rows.Close()

	if rows.Next() {
		var inserted entity.User
		if err := rows.StructScan(&inserted); err != nil {
			logger.Error("AuthRepository:CreateUser:StructScan", err)
			return nil, err
		}
		return &inserted, nil
	}
	return nil, fmt.Errorf("insert user failed")
}

func (r *AuthRepository) GetUserPermissions(ctx context.Context, userID uuid.UUID) ([]entity.Permission, error) {
	// Query để lấy tất cả permissions của user từ 2 nguồn:
	// 1. Permissions trực tiếp được gán cho user (user_permissions)
	// 2. Permissions từ các roles mà user được gán (user_roles -> role_permissions)
	query := `
		SELECT DISTINCT p.id, p.name, p.slug, p.resource, p.action, p.description, p.is_system, p.created_at, p.updated_at
		FROM permissions p
		WHERE p.id IN (
			-- Permissions từ roles của user
			SELECT rp.permission_id
			FROM role_permissions rp
			INNER JOIN user_roles ur ON rp.role_id = ur.role_id
			INNER JOIN roles r ON ur.role_id = r.id
			WHERE ur.user_id = $1 
			  AND ur.is_active = true 
			  AND r.is_active = true
			  AND (ur.expires_at IS NULL OR ur.expires_at > NOW())
			
			UNION
			
			-- Permissions trực tiếp được gán cho user
			SELECT up.permission_id
			FROM user_permissions up
			WHERE up.user_id = $1 
			  AND up.granted = true 
			  AND (up.expires_at IS NULL OR up.expires_at > NOW())
		)
		ORDER BY p.resource, p.action
	`

	var permissions []entity.Permission
	err := r.DB.SelectContext(ctx, &permissions, query, userID)
	if err != nil {
		logger.Error("AuthRepository:GetUserPermissions:Error:", err)
		return nil, err
	}

	return permissions, nil
}
