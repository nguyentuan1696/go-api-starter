package apperrors

type ErrorCode int

const (
	// Authentication & Authorization errors (1000-1099)
	ErrUnauthorized ErrorCode = 1000 + iota
	ErrForbidden

	// Validation errors (2000-2099)
	ErrInvalidInput ErrorCode = 2000 + iota

	// Resource errors (3000-3099)
	ErrNotFound ErrorCode = 3000 + iota

	// Business logic errors (4000-4099)
	ErrBusinessRule ErrorCode = 4000 + iota

	// System errors (5000-5099)
	ErrInternalServer ErrorCode = 5000 + iota
)
