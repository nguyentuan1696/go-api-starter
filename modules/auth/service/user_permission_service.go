package service

import (
	"context"
	"go-api-starter/core/constants"
	"go-api-starter/core/logger"
	"go-api-starter/modules/auth/dto"
	"go-api-starter/modules/auth/entity"
	"go-api-starter/modules/auth/mapper"

	"github.com/google/uuid"
)

func (service *AuthService) PrivateAssignPermissionToUser(ctx context.Context, req *dto.UserPermissionRequest) error {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	err := service.repo.PrivateAssignPermissionToUser(ctx, mapper.ToUserPermissionEntity(req))
	if err != nil {
		logger.Error("AuthService:PrivateAssignPermissionToUser error: %v", err)
		return err
	}

	return nil
}

func (service *AuthService) PrivateGetUserPermissions(ctx context.Context, userID uuid.UUID) ([]entity.Permission, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultTimeout)
	defer cancel()

	result, err := service.repo.GetUserPermissions(ctx, userID)
	if err != nil {
		logger.Error("AuthService:GetUserPermissions:Error:", err)
		return nil, err
	}

	return result, nil
}
