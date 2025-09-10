package service

import (
	"context"
	"go-api-starter/core/constants"
	"go-api-starter/core/errors"
	"go-api-starter/core/logger"
	"go-api-starter/core/params"
	"go-api-starter/modules/auth/dto"
	"go-api-starter/modules/auth/mapper"

	"github.com/google/uuid"
)

func (service *AuthService) PrivateGetUser(ctx context.Context, userID uuid.UUID) (*dto.UserDetailDTO, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultTimeout)
	defer cancel()

	user, err := service.repo.PrivateGetUser(ctx, userID)
	if err != nil {
		logger.Error("AuthService:PrivateGetUser:Error:", err)
		return nil, errors.NewAppError(errors.ErrInternalServer, "failed to get user", err)
	}

	if user == nil {
		logger.Info("AuthService:PrivateGetUser:UserNotFound:", userID)
		return nil, errors.NewAppError(errors.ErrNotFound, "user not found", nil)
	}

	return mapper.ToUserDetailDTO(user), nil
}

func (service *AuthService) PrivateGetUsers(ctx context.Context, params params.QueryParams) (*dto.PaginatedUserDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultTimeout)
	defer cancel()

	users, err := service.repo.PrivateGetUsers(ctx, params)
	if err != nil {
		logger.Error("AuthService:PrivateGetUsers:Error:", err)
		return nil, err
	}

	return mapper.ToUserPaginationDTO(users), nil
}
