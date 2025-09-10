package service

import (
	"context"
	"go-api-starter/core/cache"
	"go-api-starter/core/errors"
	"go-api-starter/core/params"
	"go-api-starter/modules/auth/dto"
	"go-api-starter/modules/auth/entity"
	"go-api-starter/modules/auth/repository"

	"github.com/google/uuid"
)

type AuthService struct {
	repo  repository.AuthRepositoryInterface
	cache cache.Cache
}

func NewAuthService(repo repository.AuthRepositoryInterface, cache cache.Cache) AuthServiceInterface {
	return &AuthService{repo: repo, cache: cache}
}

type AuthServiceInterface interface {
	Register(ctx context.Context, requestData *dto.RegisterRequest) (*dto.RegisterResponse, *errors.AppError)
	Login(ctx context.Context, requestData *dto.LoginRequest) (*dto.LoginResponse, *errors.AppError)
	Logout(ctx context.Context, token string) *errors.AppError
	ChangePassword(ctx context.Context, token string, requestData *dto.ChangePasswordRequest) *errors.AppError
	ForgotPassword(ctx context.Context, identifier string) (*dto.ForgotPasswordResponse, *errors.AppError)
	VerifyOTP(ctx context.Context, requestData *dto.VerifyOTPRequest) (*dto.VerifyOTPResponse, *errors.AppError)
	ResetPassword(ctx context.Context, requestData *dto.ResetPasswordRequest) *errors.AppError
	SendOTPChangePassword(ctx context.Context, token string) *errors.AppError

	// Role methods
	PrivateCreateRole(ctx context.Context, role *dto.RoleRequest) error
	PrivateGetRoles(ctx context.Context, params params.QueryParams) (*dto.PaginatedRoleDTO, error)
	PrivateGetRoleByID(ctx context.Context, id uuid.UUID) (*dto.RoleResponse, error)
	PrivateUpdateRole(ctx context.Context, id uuid.UUID, role *dto.RoleRequest) error
	PrivateDeleteRole(ctx context.Context, id uuid.UUID) error

	// Permission methods
	PrivateCreatePermission(ctx context.Context, permission *dto.PermissionRequest) error
	PrivateGetPermissions(ctx context.Context, params params.QueryParams) (*dto.PaginatedPermissionDTO, error)
	PrivateGetPermissionByID(ctx context.Context, id uuid.UUID) (*dto.PermissionResponse, error)
	PrivateUpdatePermission(ctx context.Context, id uuid.UUID, permission *dto.PermissionRequest) error
	PrivateDeletePermission(ctx context.Context, id uuid.UUID) error

	RefreshToken(ctx context.Context, token string) (*dto.RefreshTokenResponse, *errors.AppError)
	GetUserByIdentifier(ctx context.Context, identifier string) (*dto.UserResponse, *errors.AppError)

	PrivateAssignRoleToUser(ctx context.Context, req *dto.UserRoleRequest) *errors.AppError
	PrivateAssignPermissionToRole(ctx context.Context, req *dto.RolePermissionRequest) error
	PrivateAssignPermissionToUser(ctx context.Context, req *dto.UserPermissionRequest) error

	PrivateGetUserPermissions(ctx context.Context, userID uuid.UUID) ([]entity.Permission, error)
	PrivateGetUsers(ctx context.Context, params params.QueryParams) (*dto.PaginatedUserDTO, error)
	PrivateGetUser(ctx context.Context, userID uuid.UUID) (*dto.UserDetailDTO, *errors.AppError)
	PrivateGetPermissionsByUserID(ctx context.Context, userID uuid.UUID) (*[]dto.PermissionResponse, error)
	PrivateGetPermissionsByUserIDFromCache(ctx context.Context, userID uuid.UUID) (*[]dto.PermissionResponse, error)
}
