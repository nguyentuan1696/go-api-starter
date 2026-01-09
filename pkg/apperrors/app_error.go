package apperrors

import "fmt"

type AppError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Err     error     `json:"-"` // Internal error (not exposed to client)
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%v", e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NewAppError(code ErrorCode, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
