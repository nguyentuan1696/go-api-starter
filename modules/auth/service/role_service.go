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

func (service *AuthService) PrivateCreateRole(ctx context.Context, role *dto.RoleRequest) error {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	err := service.repo.PrivateCreateRole(ctx, mapper.ToRoleEntity(role))
	if err != nil {
		logger.Error("AuthService:PrivateCreateRole error: %v", err)
		return err
	}

	return nil
}

func (service *AuthService) PrivateGetRoles(ctx context.Context, params params.QueryParams) (*dto.PaginatedRoleDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	roles, err := service.repo.PrivateGetRoles(ctx, params)
	if err != nil {
		logger.Error("AuthService:PrivateGetRoles error: %v", err)
		return nil, err
	}

	return mapper.ToPaginatedRoleDTO(roles), nil
}

func (service *AuthService) PrivateGetRoleByID(ctx context.Context, id uuid.UUID) (*dto.RoleResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	role, err := service.repo.PrivateGetRoleByID(ctx, id)
	if err != nil {
		logger.Error("AuthService:PrivateGetRoleByID error: %v", err)
		return nil, err
	}

	return mapper.ToRoleDTO(role), nil
}

func (service *AuthService) PrivateUpdateRole(ctx context.Context, id uuid.UUID, role *dto.RoleRequest) error {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	err := service.repo.PrivateUpdateRole(ctx, id, mapper.ToRoleEntity(role))
	if err != nil {
		logger.Error("AuthService:PrivateUpdateRole error: %v", err)
		return err
	}

	return nil
}

func (service *AuthService) PrivateDeleteRole(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultRequestTimeout)
	defer cancel()

	err := service.repo.PrivateDeleteRole(ctx, id)
	if err != nil {
		logger.Error("AuthService:PrivateDeleteRole error: %v", err)
		return err
	}

	return nil
}
