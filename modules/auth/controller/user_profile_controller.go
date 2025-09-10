package controller

import "github.com/labstack/echo/v4"

func (controller *AuthController) UpdateUserProfile(c echo.Context) error {
	return controller.SuccessResponse(c, nil, "Update user profile success")
}
