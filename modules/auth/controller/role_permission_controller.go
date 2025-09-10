package controller

import (
	"go-api-starter/core/errors"
	"go-api-starter/modules/auth/dto"
	"go-api-starter/modules/auth/validator"

	"github.com/labstack/echo/v4"
)

func (controller *AuthController) PrivateAssignPermissionToUser(c echo.Context) error {
	ctx := c.Request().Context()

	requestData := new(dto.UserPermissionRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateAssignPermissionToUserRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	err := controller.AuthService.PrivateAssignPermissionToUser(ctx, requestData)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "Assign permission to user failed", nil)
	}

	return controller.SuccessResponse(c, nil, "Assign permission to user success")
}

func (controller *AuthController) PrivateAssignPermissionToRole(c echo.Context) error {
	ctx := c.Request().Context()

	requestData := new(dto.RolePermissionRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateAssignPermissionToRoleRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	err := controller.AuthService.PrivateAssignPermissionToRole(ctx, requestData)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "Assign permission to role failed", nil)
	}

	return controller.SuccessResponse(c, nil, "Assign permission to role success")
}

func (controller *AuthController) AssignPermissionToUser(c echo.Context) error {
	return controller.SuccessResponse(c, nil, "Assign permission to user success")
}
