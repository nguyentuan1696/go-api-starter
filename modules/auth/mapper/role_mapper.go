package mapper

import (
	"go-api-starter/modules/auth/dto"
	"go-api-starter/modules/auth/entity"
)

func ToRoleEntity(role *dto.RoleRequest) *entity.Role {
	return &entity.Role{
		Name:        role.Name,
		Slug:        role.Slug,
		Description: role.Description,
	}
}

func ToRoleDTO(role *entity.Role) *dto.RoleResponse {
	return &dto.RoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		Slug:        role.Slug,
		Description: role.Description,
		IsSystem:    role.IsSystem,
		IsActive:    role.IsActive,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}
}

func ToPaginatedRoleDTO(entity *entity.PaginatedRoleEntity) *dto.PaginatedRoleDTO {

	roleResponses := make([]dto.RoleResponse, len(entity.Items))
	for i, role := range entity.Items {
		roleResponses[i] = *ToRoleDTO(&role)
	}

	// Tính total pages
	totalPages := 0
	if entity.PageSize > 0 {
		totalPages = (entity.TotalItems + entity.PageSize - 1) / entity.PageSize
	}

	return &dto.PaginatedRoleDTO{
		Items:      roleResponses,
		TotalItems: entity.TotalItems,
		TotalPages: totalPages,
		PageNumber: entity.PageNumber,
		PageSize:   entity.PageSize,
	}
}
