package errors

type ErrorCode int

const (
	// Authentication & Authorization errors (1000-1999)
	ErrInvalidCredentials ErrorCode = 1000 + iota
	ErrTokenExpired
	ErrUnauthorized
	ErrForbidden
	ErrMissingAuthorizationHeader
	ErrInvalidTokenFormat

	// Input validation errors (2000-2999)
	ErrInvalidInput ErrorCode = 2000 + iota
	ErrInvalidEmail
	ErrInvalidPassword
	ErrInvalidFormat
	ErrInvalidRequestData

	// Resource errors (3000-3999)
	ErrNotFound ErrorCode = 3000 + iota
	ErrAlreadyExists
	ErrResourceLocked
	ErrResourceExpired
	ErrCreateFailed
	ErrGetFailed
	ErrInvalidSlug
	ErrDuplicateSlug
	ErrInvalidParent
	ErrUpdateFailed
	ErrDeleteFailed

	// Database errors (4000-4999)
	ErrDatabase ErrorCode = 4000 + iota
	ErrDatabaseTimeout
	ErrUniqueViolation
	ErrForeignKey

	// Business logic errors (5000-5999)
	ErrBusinessRule ErrorCode = 5000 + iota
	ErrInvalidState
	ErrLimitExceeded
	ErrOperationFailed

	// System errors (6000-6999)
	ErrInternalServer ErrorCode = 6000 + iota
	ErrConfiguration
	ErrThirdParty
	ErrNetwork

	// Storage errors (7000-7999)
	ErrUploadToR2 ErrorCode = 7000 + iota
	ErrSaveStorageDb
)
