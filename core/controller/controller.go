package controller

import (
	"net/http"
	"time"
	"go-api-starter/core/errors"

	"github.com/labstack/echo/v4"
)

// Response types
type (
	SuccessResponse struct {
		Status    int       `json:"status"`
		Message   string    `json:"message"`
		Data      any       `json:"data,omitempty"`
		Timestamp time.Time `json:"timestamp"`
	}

	ErrorResponse struct {
		Status    string           `json:"status"`
		Code      errors.ErrorCode `json:"code"`
		Message   string           `json:"message"`
		Details   any              `json:"details,omitempty"`
		Timestamp time.Time        `json:"timestamp"`
	}

	ValidationError struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	ValidationResponse struct {
		Success bool              `json:"success"`
		Message string            `json:"message"`
		Errors  []ValidationError `json:"errors"`
	}
)

// Response handler interface and implementation
type BaseController interface {
	BadRequest(appErrCode errors.ErrorCode, message string, details ...any) *echo.HTTPError
	InternalServerError(appErrCode errors.ErrorCode, message string, details ...any) *echo.HTTPError
	NotFound(appErrCode errors.ErrorCode, message string, details ...any) *echo.HTTPError
	Unauthorized(appErrCode errors.ErrorCode, message string, details ...any) *echo.HTTPError
	Forbidden(appErrCode errors.ErrorCode, message string, details ...any) *echo.HTTPError
	SuccessResponse(c echo.Context, data any, message string) error
	ErrorResponse(c echo.Context, err error) error
}

type responseHandler struct{}

func NewBaseController() BaseController {
	return &responseHandler{}
}

// Success response functions
func NewSuccessResponse(httpStatusCode int, data any, message string) *SuccessResponse {
	return &SuccessResponse{
		Status:    httpStatusCode,
		Message:   message,
		Data:      data,
		Timestamp: time.Now(),
	}
}

// Error response functions
func NewErrorResponse(httpStatusCode int, appErrCode errors.ErrorCode, message string, details ...any) *echo.HTTPError {
	err := &ErrorResponse{
		Status:    "error",
		Code:      appErrCode,
		Message:   message,
		Timestamp: time.Now(),
	}
	if len(details) > 0 {
		err.Details = details[0]
	}
	return echo.NewHTTPError(httpStatusCode, err)
}

// Validation functions
func NewValidationError(field, message string) ValidationError {
	return ValidationError{
		Field:   field,
		Message: message,
	}
}

// HTTP Error handlers
func (h *responseHandler) BadRequest(appErrCode errors.ErrorCode, message string, details ...any) *echo.HTTPError {
	return NewErrorResponse(http.StatusBadRequest, appErrCode, message, details...)
}

func (h *responseHandler) InternalServerError(appErrCode errors.ErrorCode, message string, details ...any) *echo.HTTPError {
	return NewErrorResponse(http.StatusInternalServerError, appErrCode, message, details...)
}

func (h *responseHandler) NotFound(appErrCode errors.ErrorCode, message string, details ...any) *echo.HTTPError {
	return NewErrorResponse(http.StatusNotFound, appErrCode, message, details...)
}

func (h *responseHandler) Unauthorized(appErrCode errors.ErrorCode, message string, details ...any) *echo.HTTPError {
	return NewErrorResponse(http.StatusUnauthorized, appErrCode, message, details...)
}

func (h *responseHandler) Forbidden(appErrCode errors.ErrorCode, message string, details ...any) *echo.HTTPError {
	return NewErrorResponse(http.StatusForbidden, appErrCode, message, details...)
}

func (h *responseHandler) ValidationError(appErrCode errors.ErrorCode, message string, details ...any) *echo.HTTPError {
	return NewErrorResponse(http.StatusBadRequest, appErrCode, message, details...)
}

func (h *responseHandler) SuccessResponse(c echo.Context, data any, message string) error {
	return c.JSON(http.StatusOK, NewSuccessResponse(http.StatusOK, data, message))
}

func (h *responseHandler) ErrorResponse(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, NewErrorResponse(http.StatusInternalServerError, errors.ErrInternalServer, err.Error()))
}
