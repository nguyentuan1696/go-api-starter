package controller

import (
	"go-api-starter/core/errors"
	"go-api-starter/core/params"
	"go-api-starter/core/utils"
	"go-api-starter/modules/auth/dto"
	"go-api-starter/modules/auth/validator"

	"github.com/labstack/echo/v4"
)

func (controller *AuthController) PrivateCreatePermission(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(dto.PermissionRequest)
	if err := c.Bind(req); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidatePermissionRequest(req)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", validationResult)
	}

	if err := controller.AuthService.PrivateCreatePermission(ctx, req); err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "create permission failed", err)
	}

	return controller.SuccessResponse(c, nil, "create permission success")
}

func (controller *AuthController) PrivateGetPermissions(c echo.Context) error {
	ctx := c.Request().Context()

	params := params.NewQueryParams(c)

	permissions, err := controller.AuthService.PrivateGetPermissions(ctx, *params)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get permissions failed", err)
	}

	return controller.SuccessResponse(c, permissions, "get permissions success")
}

func (controller *AuthController) PrivateGetPermissionByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))
	permission, err := controller.AuthService.PrivateGetPermissionByID(ctx, id)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get permission failed", err)
	}

	return controller.SuccessResponse(c, permission, "get permission success")
}

func (controller *AuthController) PrivateUpdatePermission(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))
	req := new(dto.PermissionRequest)
	if err := c.Bind(req); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidatePermissionRequest(req)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", validationResult)
	}

	if err := controller.AuthService.PrivateUpdatePermission(ctx, id, req); err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "update permission failed", err)
	}

	return controller.SuccessResponse(c, nil, "update permission success")
}

func (controller *AuthController) PrivateDeletePermission(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))
	if err := controller.AuthService.PrivateDeletePermission(ctx, id); err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "delete permission failed", err)
	}

	return controller.SuccessResponse(c, nil, "delete permission success")
}