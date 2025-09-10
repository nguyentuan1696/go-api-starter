package mapper

import (
	"go-api-starter/modules/auth/dto"
	"go-api-starter/modules/auth/entity"
)

func ToUserRoleEntity(req *dto.UserRoleRequest) *entity.UserRole {
	if req == nil {
		return nil
	}

	return &entity.UserRole{
		UserID:     req.UserID,
		RoleID:     req.RoleID,
		AssignedBy: req.AssignedBy,
		ExpiresAt:  req.ExpiresAt,
		IsActive:   req.IsActive,
	}
}
