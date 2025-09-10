package service

import (
	"context"
	"go-api-starter/core/constants"
	"go-api-starter/core/logger"
	"go-api-starter/core/params"
	"go-api-starter/modules/auth/dto"
	"go-api-starter/modules/auth/mapper"

	"github.com/google/uuid"
)

func (service *AuthService) PrivateCreatePermission(ctx context.Context, permission *dto.PermissionRequest) error {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	err := service.repo.PrivateCreatePermission(ctx, mapper.ToPermissionEntity(permission))
	if err != nil {
		logger.Error("AuthService:PrivateCreatePermission error: %v", err)
		return err
	}

	return nil
}

func (service *AuthService) PrivateGetPermissions(ctx context.Context, params params.QueryParams) (*dto.PaginatedPermissionDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	permissions, err := service.repo.PrivateGetPermissions(ctx, params)
	if err != nil {
		logger.Error("AuthService:PrivateGetPermissions error: %v", err)
		return nil, err
	}

	return mapper.ToPaginatedPermissionDTO(permissions), nil
}

func (service *AuthService) PrivateGetPermissionByID(ctx context.Context, id uuid.UUID) (*dto.PermissionResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	permission, err := service.repo.PrivateGetPermissionByID(ctx, id)
	if err != nil {
		logger.Error("AuthService:PrivateGetPermissionByID error: %v", err)
		return nil, err
	}

	return mapper.ToPermissionDTO(permission), nil
}

func (service *AuthService) PrivateUpdatePermission(ctx context.Context, id uuid.UUID, permission *dto.PermissionRequest) error {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	err := service.repo.PrivateUpdatePermission(ctx, id, mapper.ToPermissionEntity(permission))
	if err != nil {
		logger.Error("AuthService:PrivateUpdatePermission error: %v", err)
		return err
	}

	return nil
}

func (service *AuthService) PrivateDeletePermission(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	err := service.repo.PrivateDeletePermission(ctx, id)
	if err != nil {
		logger.Error("AuthService:PrivateDeletePermission error: %v", err)
		return err
	}

	return nil
}
