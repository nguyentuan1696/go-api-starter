package mapper

import (
	"go-api-starter/modules/auth/dto"
	"go-api-starter/modules/auth/entity"
)

func ToPermissionEntity(permission *dto.PermissionRequest) *entity.Permission {
	return &entity.Permission{
		Name:        permission.Name,
		Slug:        permission.Slug,
		Description: permission.Description,
		Resource:    permission.Resource,
		Action:      entity.Action(permission.Action),
	}
}

func ToPermissionDTO(permission *entity.Permission) *dto.PermissionResponse {
	return &dto.PermissionResponse{
		ID:          permission.ID,
		Name:        permission.Name,
		Slug:        permission.Slug,
		Resource:    permission.Resource,
		Action:      string(permission.Action),
		Description: permission.Description,
		IsSystem:    permission.IsSystem,
		CreatedAt:   permission.CreatedAt,
		UpdatedAt:   permission.UpdatedAt,
	}
}

func ToPaginatedPermissionDTO(entity *entity.PaginatedPermissionEntity) *dto.PaginatedPermissionDTO {
	permissionResponses := make([]dto.PermissionResponse, len(entity.Items))
	for i, role := range entity.Items {
		permissionResponses[i] = *ToPermissionDTO(&role)
	}

	// Tính total pages
	totalPages := 0
	if entity.PageSize > 0 {
		totalPages = (entity.TotalItems + entity.PageSize - 1) / entity.PageSize
	}

	return &dto.PaginatedPermissionDTO{
		Items:      permissionResponses,
		TotalItems: entity.TotalItems,
		TotalPages: totalPages,
		PageNumber: entity.PageNumber,
		PageSize:   entity.PageSize,
	}
}

func ToPermissionDTOs(permissions *[]entity.Permission) *[]dto.PermissionResponse {
	if permissions == nil {
		return nil
	}

	permissionResponses := make([]dto.PermissionResponse, len(*permissions))
	for i, permission := range *permissions {
		permissionResponses[i] = *ToPermissionDTO(&permission)
	}

	return &permissionResponses
}
