package mapper

import (
	"go-api-starter/modules/auth/dto"
	"go-api-starter/modules/auth/entity"
)

func ToUserDTO(user *entity.User) *dto.UserResponse {
	if user == nil {
		return nil
	}
	return &dto.UserResponse{
		ID:              user.ID,
		Phone:           user.Phone,
		Username:        user.Username,
		Email:           user.Email,
		EmailVerifiedAt: user.EmailVerifiedAt,
		PhoneVerifiedAt: user.PhoneVerifiedAt,
		Password:        user.Password,
		LockedUntil:     user.LockedUntil,
		IsActive:        user.IsActive,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
	}
}

func ToUserPaginationDTO(entity *entity.PaginatedUserEntity) *dto.PaginatedUserDTO {

	if entity == nil {
		return nil
	}

	// Convert từng user entity sang user response
	userResponses := make([]dto.UserResponse, len(entity.Items))
	for i, user := range entity.Items {
		userResponses[i] = *ToUserDTO(&user)
	}

	// Tính total pages
	totalPages := 0
	if entity.PageSize > 0 {
		totalPages = (entity.TotalItems + entity.PageSize - 1) / entity.PageSize
	}

	return &dto.PaginatedUserDTO{
		Items:      userResponses,
		TotalItems: entity.TotalItems,
		TotalPages: totalPages,
		PageNumber: entity.PageNumber,
		PageSize:   entity.PageSize,
	}
}

func ToUserDetailDTO(user *entity.UserDetail) *dto.UserDetailDTO {
	if user == nil {
		return nil
	}
	return &dto.UserDetailDTO{
		ID:          user.ID,
		Email:       user.Email,
		Phone:       user.Phone,
		Username:    user.Username,
		IsActive:    user.IsActive,
		CreatedAt:   user.CreatedAt,
		DisplayName: user.DisplayName,
		FullName:    user.FullName,
		Avatar:      user.Avatar,
		DateOfBirth: user.DateOfBirth,
		Gender:      user.Gender,
		Roles:       user.Roles,
	}
}
