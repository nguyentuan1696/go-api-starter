package service

import (
	"context"
	"encoding/json"
	"go-api-starter/core/constants"
	"go-api-starter/core/logger"
	"go-api-starter/core/utils"
	"go-api-starter/modules/auth/dto"
	"go-api-starter/modules/auth/mapper"

	"github.com/google/uuid"
)

func (service *AuthService) PrivateGetPermissionsByUserIDFromCache(ctx context.Context, userID uuid.UUID) (*[]dto.PermissionResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	key := utils.GenerateUserPermissionsKey(userID.String())

	val, err := service.cache.Get(ctx, key).Result()
	if err != nil {
		logger.Error("AuthService:PrivateGetPermissionsByUserIDFromCache error: %v", err)
		return nil, err
	}

	permissionResponses := &[]dto.PermissionResponse{}
	errUnmarshal := json.Unmarshal([]byte(val), permissionResponses)
	if errUnmarshal != nil {
		logger.Error("AuthService:PrivateGetPermissionsByUserIDFromCache error: %v", errUnmarshal)
		return nil, errUnmarshal
	}

	return permissionResponses, nil
}

func (service *AuthService) PrivateGetPermissionsByUserID(ctx context.Context, userID uuid.UUID) (*[]dto.PermissionResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	permissions, err := service.repo.PrivateGetPermissionsByUserID(ctx, userID)
	if err != nil {
		logger.Error("AuthService:PrivateGetPermissionsByUserID error: %v", err)
		return nil, err
	}

	rolePermissions := mapper.ToPermissionDTOs(permissions)

	// Cache the rolePermissions
	key := utils.GenerateUserPermissionsKey(userID.String())
	rolePermissionsJSON, err := json.Marshal(rolePermissions)
	if err != nil {
		logger.Error("AuthService:PrivateGetPermissionsByUserID error: %v", err)
		return nil, err
	}
	errSet := service.cache.Set(ctx, key, rolePermissionsJSON, constants.DefaultZeroExpiration)
	if errSet != nil {
		logger.Error("AuthService:PrivateGetPermissionsByRole error: %v", errSet)
		return nil, errSet
	}

	return rolePermissions, nil
}

func (service *AuthService) PrivateAssignPermissionToRole(ctx context.Context, req *dto.RolePermissionRequest) error {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	err := service.repo.PrivateAssignPermissionToRole(ctx, req.RoleID, req.PermissionID, req.GrantedBy)
	if err != nil {
		logger.Error("AuthService:PrivateAssignPermissionToRole error: %v", err)
		return err
	}

	return nil
}
