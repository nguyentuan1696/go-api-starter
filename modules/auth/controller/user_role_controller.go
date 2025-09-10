package controller

import (
	"go-api-starter/core/errors"
	"go-api-starter/modules/auth/dto"
	"go-api-starter/modules/auth/validator"

	"github.com/labstack/echo/v4"
)

func (controller *AuthController) PrivateAssignRoleToUser(c echo.Context) error {
	ctx := c.Request().Context()

	requestData := new(dto.UserRoleRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateAssignRoleToUserRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	err := controller.AuthService.PrivateAssignRoleToUser(ctx, requestData)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "Assign role to user failed", nil)
	}

	return controller.SuccessResponse(c, nil, "Assign role to user success")
}
