package handler

import (
	"go-api-starter/pkg/apperrors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type (
	SuccessResponse struct {
		Message   string    `json:"message"`
		Data      any       `json:"data,omitempty"`
		Meta      any       `json:"meta,omitempty"`
		Timestamp time.Time `json:"timestamp"`
	}

	ErrorResponse struct {
		Code      apperrors.ErrorCode `json:"code"`
		Message   string              `json:"message"`
		Details   any                 `json:"details,omitempty"`
		Timestamp time.Time           `json:"timestamp"`
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

type BaseHandler interface {
	SuccessResponse(c echo.Context, data any, meta any, message string) error
	BadRequest(appErrCode apperrors.ErrorCode, message string, details ...any) *echo.HTTPError
	NotFound(appErrCode apperrors.ErrorCode, message string, details ...any) *echo.HTTPError
	InternalServerError(appErrCode apperrors.ErrorCode, message string, details ...any) *echo.HTTPError
}

type baseHandler struct{}

func NewBaseHandler() BaseHandler {
	return &baseHandler{}
}

func NewSuccessResponse(data any, meta any, message string) *SuccessResponse {
	return &SuccessResponse{
		Data:      data,
		Meta:      meta,
		Message:   message,
		Timestamp: time.Now(),
	}
}

func NewErrorResponse(httpStatusCode int, appErrCode apperrors.ErrorCode, message string, details ...any) *echo.HTTPError {
	err := &ErrorResponse{
		Code:      appErrCode,
		Message:   message,
		Timestamp: time.Now(),
	}
	if len(details) > 0 {
		err.Details = details[0]
	}
	return echo.NewHTTPError(httpStatusCode, err)
}

func (h *baseHandler) SuccessResponse(c echo.Context, data any, meta any, message string) error {
	return c.JSON(http.StatusOK, NewSuccessResponse(data, meta, message))
}

func (h *baseHandler) BadRequest(appErrCode apperrors.ErrorCode, message string, details ...any) *echo.HTTPError {
	return NewErrorResponse(http.StatusBadRequest, appErrCode, message, details...)
}

func (h *baseHandler) NotFound(appErrCode apperrors.ErrorCode, message string, details ...any) *echo.HTTPError {
	return NewErrorResponse(http.StatusNotFound, appErrCode, message, details...)
}

func (h *baseHandler) InternalServerError(appErrCode apperrors.ErrorCode, message string, details ...any) *echo.HTTPError {
	return NewErrorResponse(http.StatusInternalServerError, appErrCode, message, details...)
}
