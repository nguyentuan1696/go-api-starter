package service

import (
	"context"

	"go-api-starter/core/constants"
	"go-api-starter/core/errors"
	"go-api-starter/core/logger"
	"go-api-starter/modules/auth/dto"
	"go-api-starter/modules/auth/mapper"
)

func (service *AuthService) PrivateAssignRoleToUser(ctx context.Context, req *dto.UserRoleRequest) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultTimeout)
	defer cancel()

	err := service.repo.PrivateAssignRoleToUser(ctx, mapper.ToUserRoleEntity(req))
	if err != nil {
		logger.Error("AuthService:PrivateAssignRoleToUser error: %v", err)
		return errors.NewAppError(errors.ErrInternalServer, "Assign role to user failed", err)
	}

	return nil
}
