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
	DefaultAccessTokenExpiry        = 24 * time.Hour
	DefaultRefreshTokenExpiry       = 7 * 24 * time.Hour
	DefaultResetPasswordTokenExpiry = 5 * time.Minute
	DefaultOTPExpiration            = 5 * time.Minute
	DefaultZeroExpiration           = 0
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
	DefaultPageNumber = 1
	DefaultPageSize   = 10
)

const (
	DefaultTimeout  = 10 * time.Second
	LongTimeout     = 30 * time.Second
	ShortTimeout    = 5 * time.Second
	DatabaseTimeout = 10 * time.Second
	CacheTimeout    = 5 * time.Second
)

// Database Configuration Constants
const (
	DatabaseSSLMode                = "disable"
	DatabaseMaxOpenConns           = 25
	DatabaseMaxIdleConns           = 5
	DatabaseConnMaxLifetime        = 300 // seconds
	DatabaseConnectTimeout         = 10  // seconds
	DatabaseStatementTimeout       = 30  // seconds
	DatabaseIdleInTxSessionTimeout = 60  // seconds
)

// Server Configuration Constants
const (
	ServerReadTimeout  = 30  // seconds
	ServerWriteTimeout = 30  // seconds
	ServerIdleTimeout  = 120 // seconds
)

// Redis Configuration Constants
const (
	RedisPoolSize     = 10
	RedisMinIdleConns = 5
	RedisDialTimeout  = 5 // seconds
	RedisReadTimeout  = 3 // seconds
	RedisWriteTimeout = 3 // seconds
)

// Rate Limiting Constants
const (
	RateLimitRequests = 100
	RateLimitWindow   = 60 // seconds
)

// File Upload Constants
const (
	UploadMaxSize      = 2 * 1024 * 1024 // 2MB in bytes
	UploadAllowedTypes = "image/jpeg,image/png,image/gif,image/webp"
	UploadPath         = "uploads"
)

// Logger Configuration Constants
const (
	LogRetentionDays = 15        // Số ngày giữ lại log files
	LogFileName      = "app.log" // Tên file log mặc định
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
