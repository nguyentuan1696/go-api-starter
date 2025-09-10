package controller

import (
	"go-api-starter/core/errors"
	"go-api-starter/core/params"
	"go-api-starter/core/utils"
	"go-api-starter/modules/auth/dto"
	"go-api-starter/modules/auth/validator"

	"github.com/labstack/echo/v4"
)

func (controller *AuthController) PrivateCreateRole(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(dto.RoleRequest)
	if err := c.Bind(req); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateRoleRequest(req)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", validationResult)
	}

	if err := controller.AuthService.PrivateCreateRole(ctx, req); err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "create role failed", err)
	}

	return controller.SuccessResponse(c, nil, "create role success")
}

func (controller *AuthController) PrivateGetRoles(c echo.Context) error {
	ctx := c.Request().Context()

	params := params.NewQueryParams(c)

	roles, err := controller.AuthService.PrivateGetRoles(ctx, *params)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get roles failed", err)
	}

	return controller.SuccessResponse(c, roles, "get roles success")
}

func (controller *AuthController) PrivateGetRoleByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))
	role, err := controller.AuthService.PrivateGetRoleByID(ctx, id)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get role failed", err)
	}

	return controller.SuccessResponse(c, role, "get role success")
}

func (controller *AuthController) PrivateUpdateRole(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))
	req := new(dto.RoleRequest)
	if err := c.Bind(req); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateRoleRequest(req)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", validationResult)
	}

	if err := controller.AuthService.PrivateUpdateRole(ctx, id, req); err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "update role failed", err)
	}

	return controller.SuccessResponse(c, nil, "update role success")
}

func (controller *AuthController) PrivateDeleteRole(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))
	if err := controller.AuthService.PrivateDeleteRole(ctx, id); err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "delete role failed", err)
	}

	return controller.SuccessResponse(c, nil, "delete role success")
}
