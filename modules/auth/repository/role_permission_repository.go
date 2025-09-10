package repository

import (
	"context"
	"database/sql"
	"go-api-starter/core/logger"
	"go-api-starter/modules/auth/entity"

	"github.com/google/uuid"
)

func (repo *AuthRepository) PrivateGetPermissionsByUserID(ctx context.Context, userID uuid.UUID) (*[]entity.Permission, error) {
	query := `
		SELECT DISTINCT p.resource, p.action
		FROM user_roles ur
		JOIN role_permissions rp ON ur.role_id = rp.role_id
		JOIN permissions p ON rp.permission_id = p.id
		WHERE ur.user_id = $1
		  AND ur.is_active = true
		  AND (ur.expires_at IS NULL OR ur.expires_at > NOW())
		ORDER BY p.resource, p.action
	`

	var permissions []entity.Permission
	err := repo.DB.SelectContext(ctx, &permissions, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("AuthRepository:PrivateGetPermissionsByUserID:NoRows:", err)
			return nil, nil
		}
		logger.Error("AuthRepository:PrivateGetPermissionsByUserID:Error:", err)
		return nil, err
	}

	return &permissions, nil
}

func (repo *AuthRepository) PrivateAssignPermissionToRole(ctx context.Context, roleID uuid.UUID, permissionIDs []uuid.UUID, grantedBy uuid.UUID) error {
	// Nếu không có permission nào để gán, return nil
	if len(permissionIDs) == 0 {
		return nil
	}

	// Xóa tất cả permissions hiện tại của role trước khi gán mới
	deleteQuery := `DELETE FROM role_permissions WHERE role_id = $1`
	err := repo.DB.ExecContext(ctx, deleteQuery, roleID)
	if err != nil {
		logger.Error("AuthRepository:PrivateAssignPermissionToRole:Delete:Error:", err)
		return err
	}

	// Chuẩn bị data để batch insert với struct entity.RolePermission
	rolePermissions := make([]entity.RolePermission, len(permissionIDs))

	for i, permissionID := range permissionIDs {
		rolePermissions[i] = entity.RolePermission{
			RoleID:       roleID,
			PermissionID: permissionID,
			GrantedBy:    &grantedBy,
			// GrantedAt sẽ được set bởi NOW() trong SQL
		}
	}

	// Batch insert với NamedExecContext và NOW() SQL function
	insertQuery := `
		INSERT INTO role_permissions (role_id, permission_id, granted_by, granted_at)
		VALUES (:role_id, :permission_id, :granted_by, NOW())
	`

	for _, rolePermission := range rolePermissions {
		_, err := repo.DB.NamedExecContext(ctx, insertQuery, rolePermission)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			logger.Error("AuthRepository:PrivateAssignPermissionToRole:Insert:Error:", err)
			return err
		}
	}

	return nil
}
