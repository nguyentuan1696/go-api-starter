package controller

import (
	"go-api-starter/core/errors"
	"go-api-starter/core/logger"
	"go-api-starter/core/params"
	"go-api-starter/core/utils"

	"github.com/labstack/echo/v4"
)

func (controller *AuthController) PrivateGetUsers(c echo.Context) error {

	ctx := c.Request().Context()

	params := params.NewQueryParams(c)

	users, err := controller.AuthService.PrivateGetUsers(ctx, *params)
	if err != nil {
		logger.Error("AuthController:PrivateGetUsers:Error:", err)
		return controller.InternalServerError(errors.ErrInternalServer, "Get users failed", nil)
	}

	return controller.SuccessResponse(c, users, "Get users success")
}

func (controller *AuthController) PrivateGetUser(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))
	user, errGet := controller.AuthService.PrivateGetUser(ctx, id)
	if errGet != nil {
		logger.Error("AuthController:PrivateGetUser:Error:", errGet)
		// Check if it's a not found error
		if errGet.Code == errors.ErrNotFound {
			return controller.NotFound(errors.ErrNotFound, "user not found", errGet)
		}
		return controller.InternalServerError(errors.ErrInternalServer, "failed to get user", errGet)
	}

	return controller.SuccessResponse(c, user, "get user success")
}

func (controller *AuthController) PrivateUpdateUser(c echo.Context) error {
	return nil
}
