package utils

import (
	"context"
	"fmt"
	"sync"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/redis/go-redis/v9"
)

// Các bộ ký tự có thể sử dụng
const (
	AlphanumericOTP = "0123456789"
	OTPExpiry       = 5 * time.Minute // OTP hết hạn sau 5 phút
	MaxOTPAttempts  = 3               // Tối đa 3 lần thử
)

// OTPType định nghĩa loại OTP
type OTPType string

const (
	OTPTypeForgotPassword OTPType = "forgot_password"
	OTPTypeVerification   OTPType = "verification"
	OTPTypeLogin          OTPType = "login"
)

// OTPData chứa thông tin OTP
type OTPData struct {
	Code       string    `json:"code"`
	Type       OTPType   `json:"type"`
	Identifier string    `json:"identifier"` // phone hoặc email
	Attempts   int       `json:"attempts"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiresAt  time.Time `json:"expires_at"`
}

// OTPService quản lý OTP
type OTPService struct {
	redisClient *redis.Client
}

// NewOTPService tạo OTP service mới
func NewOTPService(redisClient *redis.Client) *OTPService {
	return &OTPService{
		redisClient: redisClient,
	}
}

func GenerateOTP() string {
	// Sử dụng bộ ký tự tùy chỉnh
	id, err := gonanoid.Generate(AlphanumericOTP, 6)
	if err != nil {
		return ""
	}
	return id
}

// GenerateAndStoreOTP tạo và lưu trữ OTP trong Redis
func (s *OTPService) GenerateAndStoreOTP(ctx context.Context, identifier string, otpType OTPType) (*OTPData, error) {
	// Tạo OTP code
	code := GenerateOTP()
	if code == "" {
		return nil, fmt.Errorf("failed to generate OTP")
	}

	// Tạo OTP data
	otpData := &OTPData{
		Code:       code,
		Type:       otpType,
		Identifier: identifier,
		Attempts:   0,
		CreatedAt:  time.Now(),
		ExpiresAt:  time.Now().Add(OTPExpiry),
	}

	// Tạo Redis key
	key := fmt.Sprintf("otp:%s:%s", string(otpType), identifier)

	// Lưu vào Redis với expiry
	err := s.redisClient.Set(ctx, key, otpData, OTPExpiry).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to store OTP: %w", err)
	}

	return otpData, nil
}

// VerifyOTP xác thực OTP
func (s *OTPService) VerifyOTP(ctx context.Context, identifier string, code string, otpType OTPType) (bool, error) {
	// Tạo Redis key
	key := fmt.Sprintf("otp:%s:%s", string(otpType), identifier)

	// Lấy OTP data từ Redis
	var otpData OTPData
	err := s.redisClient.Get(ctx, key).Scan(&otpData)
	if err == redis.Nil {
		return false, fmt.Errorf("OTP not found or expired")
	}
	if err != nil {
		return false, fmt.Errorf("failed to get OTP: %w", err)
	}

	// Kiểm tra số lần thử
	if otpData.Attempts >= MaxOTPAttempts {
		// Xóa OTP nếu đã thử quá nhiều lần
		s.redisClient.Del(ctx, key)
		return false, fmt.Errorf("maximum OTP attempts exceeded")
	}

	// Kiểm tra OTP code
	if otpData.Code != code {
		// Tăng số lần thử
		otpData.Attempts++
		s.redisClient.Set(ctx, key, otpData, time.Until(otpData.ExpiresAt))
		return false, fmt.Errorf("invalid OTP code")
	}

	// OTP hợp lệ, xóa khỏi Redis
	s.redisClient.Del(ctx, key)
	return true, nil
}

// DeleteOTP xóa OTP khỏi Redis
func (s *OTPService) DeleteOTP(ctx context.Context, identifier string, otpType OTPType) error {
	key := fmt.Sprintf("otp:%s:%s", string(otpType), identifier)
	return s.redisClient.Del(ctx, key).Err()
}

// GetOTPInfo lấy thông tin OTP (không bao gồm code)
func (s *OTPService) GetOTPInfo(ctx context.Context, identifier string, otpType OTPType) (*OTPData, error) {
	key := fmt.Sprintf("otp:%s:%s", string(otpType), identifier)

	var otpData OTPData
	err := s.redisClient.Get(ctx, key).Scan(&otpData)
	if err == redis.Nil {
		return nil, fmt.Errorf("OTP not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get OTP info: %w", err)
	}

	// Không trả về code để bảo mật
	otpData.Code = ""
	return &otpData, nil
}

// Global OTP service variables
var (
	globalOTPService *OTPService
	globalOTPOnce    sync.Once
)

// InitOTPService khởi tạo global OTP service một lần
func InitOTPService(redisClient *redis.Client) {
	globalOTPOnce.Do(func() {
		globalOTPService = NewOTPService(redisClient)
	})
}

// GetOTPService trả về global OTP service
func GetOTPService() *OTPService {
	if globalOTPService == nil {
		panic("OTP service not initialized. Call InitOTPService first")
	}
	return globalOTPService
}

// Global wrapper functions để sử dụng trực tiếp

// GenerateAndStoreOTPGlobal tạo và lưu trữ OTP sử dụng global service
func GenerateAndStoreOTPGlobal(ctx context.Context, identifier string, otpType OTPType) (*OTPData, error) {
	return GetOTPService().GenerateAndStoreOTP(ctx, identifier, otpType)
}

// VerifyOTPGlobal xác thực OTP sử dụng global service
func VerifyOTPGlobal(ctx context.Context, identifier string, code string, otpType OTPType) (bool, error) {
	return GetOTPService().VerifyOTP(ctx, identifier, code, otpType)
}

// DeleteOTPGlobal xóa OTP sử dụng global service
func DeleteOTPGlobal(ctx context.Context, identifier string, otpType OTPType) error {
	return GetOTPService().DeleteOTP(ctx, identifier, otpType)
}

// GetOTPInfoGlobal lấy thông tin OTP sử dụng global service
func GetOTPInfoGlobal(ctx context.Context, identifier string, otpType OTPType) (*OTPData, error) {
	return GetOTPService().GetOTPInfo(ctx, identifier, otpType)
}
