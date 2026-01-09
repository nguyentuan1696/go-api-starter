package constants

import "time"

const (
	// Redis key prefixes
	RedisKeyPrefix = "tinup_backend_psyduck:"

	// OTP related keys
	RedisKeyOTPChangePassword = RedisKeyPrefix + "otp_change_password:"
)

const (
	TokenBlacklistKey = "token_blacklist:jti:"
)

const (
	ScopeTokenAccess            = "access"
	ScopeTokenRefresh           = "refresh"
	ScopeTokenResetPassword     = "reset_password"
	ScopeTokenEmailVerification = "email_verification"
)

// Giới hạn login
const (
	MaxLoginAttempts = 5
	BlockDuration    = 15 * time.Minute
)

// Timeout request
const (
	DefaultRequestTimeout = 5 * time.Second
)

const (
	DefaultTimeout  = 10 * time.Second
	LongTimeout     = 30 * time.Second
	ShortTimeout    = 5 * time.Second
	DatabaseTimeout = 10 * time.Second
	CacheTimeout    = 5 * time.Second
)

// Default Values Constants
const (
	DefaultEmptyString = "" // Giá trị chuỗi rỗng mặc định
	DefaultZeroValue   = 0  // Giá trị số 0 mặc định
)

// OrderState constants
const (
	OrderStatePending   = "pending"
	OrderStateConfirmed = "confirmed"
	OrderStateShipped   = "shipped"
	OrderStateDelivered = "delivered"
	OrderStateCancelled = "cancelled"
)

// PaymentStatus constants
const (
	PaymentStatusPending   = "pending"
	PaymentStatusPaid      = "paid"
	PaymentStatusFailed    = "failed"
	PaymentStatusRefunded  = "refunded"
	PaymentStatusCancelled = "cancelled"
)

// Context Key constants
const (
	ContextTokenData = "token_data"
)
