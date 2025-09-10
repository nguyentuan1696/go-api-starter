package repository

import (
	"context"
	"database/sql"
	"go-api-starter/core/logger"
	"go-api-starter/modules/auth/entity"
)

func (repo *AuthRepository) PrivateAssignPermissionToUser(ctx context.Context, req *entity.UserPermission) error {
	query := `
		INSERT INTO user_permissions (user_id, permission_id, granted_by, granted_at)
		VALUES ($1, $2, $3, NOW())
	`
	err := repo.DB.ExecContext(ctx, query, req.UserID, req.PermissionID, req.GrantedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		logger.Error("AuthRepository:PrivateAssignPermissionToUser:Error:", err)
		return err
	}
	return nil
}
