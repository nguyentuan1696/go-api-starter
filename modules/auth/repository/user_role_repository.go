package repository

import (
	"context"
	"database/sql"
	"go-api-starter/core/logger"
	"go-api-starter/modules/auth/entity"
)

func (repo *AuthRepository) PrivateAssignRoleToUser(ctx context.Context, req *entity.UserRole) error {
	// Insert user role assignment vào database
	query := `
		INSERT INTO user_roles (user_id, role_id, assigned_by, assigned_at, expires_at, is_active)
		VALUES (:user_id, :role_id, :assigned_by, :assigned_at, :expires_at, :is_active)
	`

	_, err := repo.DB.NamedExecContext(ctx, query, req)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		logger.Error("AuthRepository:PrivateAssignRoleToUser:Error:", err)
		return err
	}

	return nil
}
