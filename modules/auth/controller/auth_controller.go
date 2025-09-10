package controller

import (
	"go-api-starter/core/errors"
	"go-api-starter/core/logger"
	"go-api-starter/core/utils"
	"go-api-starter/modules/auth/dto"
	"go-api-starter/modules/auth/validator"

	"github.com/labstack/echo/v4"
)

func (controller *AuthController) VerifyOTP(c echo.Context) error {
	ctx := c.Request().Context()

	requestData := new(dto.VerifyOTPRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateVerifyOTPRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	otpResponse, err := controller.AuthService.VerifyOTP(ctx, requestData)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "Invalid request data", validationResult)
	}

	return controller.SuccessResponse(c, otpResponse, "Verify OTP success")
}

func (controller *AuthController) ResetPassword(c echo.Context) error {
	ctx := c.Request().Context()

	requestData := new(dto.ResetPasswordRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateResetPasswordRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	errReset := controller.AuthService.ResetPassword(ctx, requestData)
	if errReset != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "Invalid request data", validationResult)
	}

	return controller.SuccessResponse(c, nil, "Update password success")
}

func (controller *AuthController) Register(c echo.Context) error {
	ctx := c.Request().Context()

	requestData := new(dto.RegisterRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateRegisterRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	registerResponse, err := controller.AuthService.Register(ctx, requestData)
	if err != nil {
		// Handle different error types appropriately
		if err.Code == errors.ErrAlreadyExists {
			return controller.BadRequest(err.Code, err.Message, nil)
		}
		return controller.InternalServerError(err.Code, err.Message, err)
	}

	return controller.SuccessResponse(c, registerResponse, "Register success")
}

func (controller *AuthController) Login(c echo.Context) error {
	ctx := c.Request().Context()

	requestData := new(dto.LoginRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateLoginRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	loginResponse, err := controller.AuthService.Login(ctx, requestData)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "Invalid request data", err)
	}

	return controller.SuccessResponse(c, loginResponse, "Login success")
}

func (controller *AuthController) Logout(c echo.Context) error {
	ctx := c.Request().Context()

	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	errLogout := controller.AuthService.Logout(ctx, token)
	if errLogout != nil {
		logger.Error("AuthController:Logout:Error:", errLogout)
		return controller.InternalServerError(errors.ErrInternalServer, "Logout failed", nil)
	}

	return controller.SuccessResponse(c, nil, "Logout success")
}

func (controller *AuthController) ForgotPassword(c echo.Context) error {

	ctx := c.Request().Context()

	requestData := new(dto.ForgotPasswordRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateForgotPasswordRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	result, err := controller.AuthService.ForgotPassword(ctx, requestData.Identifier)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "Forgot password failed", err)
	}

	return controller.SuccessResponse(c, result, "Forgot password success")
}

func (controller *AuthController) SendOTPChangePassword(c echo.Context) error {
	ctx := c.Request().Context()

	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid token", nil)
	}

	errSend := controller.AuthService.SendOTPChangePassword(ctx, token)
	if errSend != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "Get OTP change password failed", err)
	}

	return controller.SuccessResponse(c, nil, "Get OTP change password success")
}

func (controller *AuthController) ChangePassword(c echo.Context) error {
	ctx := c.Request().Context()

	requestData := new(dto.ChangePasswordRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateChangePasswordRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid token", nil)
	}

	errUpdate := controller.AuthService.ChangePassword(ctx, token, requestData)
	if errUpdate != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "Change password failed", errUpdate)
	}

	return controller.SuccessResponse(c, nil, "Change password success")
}

func (controller *AuthController) RefreshToken(c echo.Context) error {
	ctx := c.Request().Context()

	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid token", nil)
	}

	refreshTokenResponse, errRefresh := controller.AuthService.RefreshToken(ctx, token)
	if errRefresh != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "Refresh token failed", nil)
	}

	return controller.SuccessResponse(c, refreshTokenResponse, "Refresh token success")
}
