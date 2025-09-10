package repository

import (
	"context"
	"go-api-starter/core/database"
	"go-api-starter/core/params"
	"go-api-starter/modules/auth/entity"

	"github.com/google/uuid"
)

// AuthRepository handles all authentication and authorization related database operations
type AuthRepository struct {
	DB database.Database
}

// NewAuthRepository creates a new instance of AuthRepository
func NewAuthRepository(db database.Database) *AuthRepository {
	return &AuthRepository{DB: db}
}

// AuthRepositoryInterface defines the contract for authentication repository operations
type AuthRepositoryInterface interface {
	// ========================================
	// Public User Operations
	// ========================================
	GetUserByIdentifier(ctx context.Context, identifier string) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	GetUserPermissions(ctx context.Context, userID uuid.UUID) ([]entity.Permission, error)

	// ========================================
	// Private User Management Operations
	// ========================================
	PrivateGetUsers(ctx context.Context, params params.QueryParams) (*entity.PaginatedUserEntity, error)
	PrivateGetUser(ctx context.Context, id uuid.UUID) (*entity.UserDetail, error)
	PrivateUpdateUser(ctx context.Context, user *entity.User, userId uuid.UUID) error
	PrivateUpdatePasswordUser(ctx context.Context, userID uuid.UUID, password string) error

	// ========================================
	// Role Management Operations
	// ========================================
	PrivateCreateRole(ctx context.Context, role *entity.Role) error
	PrivateGetRoles(ctx context.Context, params params.QueryParams) (*entity.PaginatedRoleEntity, error)
	PrivateGetRoleByID(ctx context.Context, id uuid.UUID) (*entity.Role, error)
	PrivateUpdateRole(ctx context.Context, id uuid.UUID, role *entity.Role) error
	PrivateDeleteRole(ctx context.Context, id uuid.UUID) error

	// ========================================
	// Permission Management Operations
	// ========================================
	PrivateCreatePermission(ctx context.Context, permission *entity.Permission) error
	PrivateGetPermissions(ctx context.Context, params params.QueryParams) (*entity.PaginatedPermissionEntity, error)
	PrivateGetPermissionByID(ctx context.Context, id uuid.UUID) (*entity.Permission, error)
	PrivateUpdatePermission(ctx context.Context, id uuid.UUID, permission *entity.Permission) error
	PrivateDeletePermission(ctx context.Context, id uuid.UUID) error

	// ========================================
	// Role & Permission Assignment Operations
	// ========================================
	PrivateAssignRoleToUser(ctx context.Context, req *entity.UserRole) error
	PrivateAssignPermissionToRole(ctx context.Context, roleID uuid.UUID, permissionIDs []uuid.UUID, grantedBy uuid.UUID) error
	PrivateGetPermissionsByUserID(ctx context.Context, userId uuid.UUID) (*[]entity.Permission, error)
	PrivateAssignPermissionToUser(ctx context.Context, req *entity.UserPermission) error
}
