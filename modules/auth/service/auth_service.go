package service

import (
	"context"
	"fmt"
	"go-api-starter/core/constants"
	"go-api-starter/core/errors"
	"go-api-starter/core/logger"
	"go-api-starter/core/utils"
	"go-api-starter/modules/auth/dto"
	"go-api-starter/modules/auth/entity"
	"go-api-starter/modules/auth/mapper"
)

func (service *AuthService) SendOTPChangePassword(ctx context.Context, token string) *errors.AppError {

	// Check if token is blacklisted
	blacklisted, err := service.cache.IsTokenBlacklisted(ctx, token)
	if err != nil {
		logger.Error("AuthService:SendOTPChangePassword:IsTokenBlacklisted:Error:", err)
		return errors.NewAppError(errors.ErrInternalServer, "failed to check token blacklist", err)
	}
	if blacklisted {
		return errors.NewAppError(errors.ErrUnauthorized, "token is blacklisted", nil)
	}

	tokenData, err := utils.ValidateAndParseToken(token)
	if err != nil {
		logger.Error("AuthService:SendOTPChangePassword:ValidateAndParseToken:Error:", err)
		return errors.NewAppError(errors.ErrInternalServer, "failed to validate token", err)
	}

	user, errGet := service.GetUserByIdentifier(ctx, string(utils.ToString(tokenData.UserID)))
	if errGet != nil || user == nil {
		logger.Error("AuthService:SendOTPChangePassword:GetUserByIdentifier:Error:", errGet)
		return errors.NewAppError(errors.ErrInternalServer, "failed to get user", errGet)
	}

	// Kiểm tra trạng thái xác minh email và phone
	isEmailVerified := user.EmailVerifiedAt != nil
	isPhoneVerified := user.PhoneVerifiedAt != nil

	// Kiểm tra xem có ít nhất một kênh đã được xác minh
	if !isEmailVerified && !isPhoneVerified {
		return errors.NewAppError(errors.ErrUnauthorized, "no verified contact method available", nil)
	}

	// Generate OTP
	otpCode := utils.GenerateOTP()

	// Save OTP to cache
	key := constants.RedisKeyOTPChangePassword + utils.ToString(user.ID)
	errCache := service.cache.SetOTP(ctx, key, otpCode)
	if errCache != nil {
		logger.Error("AuthService:SendOTPChangePassword:SetOTP:Error:", errCache)
		return errors.NewAppError(errors.ErrInternalServer, "failed to save OTP", errCache)
	}

	// Ưu tiên gửi email nếu đã xác minh, nếu không thì gửi SMS
	if isEmailVerified {
		// Tạo template data cho OTP
		data := utils.TemplateData{
			OTPCode: otpCode,
		}

		// Gửi email với OTP template
		errSend := utils.SendTemplateEmailFromTemplatesDir(
			[]string{*user.Email},
			"Your OTP Code",
			"otp_email.html",
			data,
		)
		if errSend != nil {
			logger.Error("AuthService:SendOTPChangePassword:SendTemplateEmailFromTemplatesDir:Error:", errSend)
			return errors.NewAppError(errors.ErrInternalServer, "failed to send OTP email", errSend)
		}
	} else if isPhoneVerified {
		// Gửi SMS OTP
		// TODO: Implement SMS sending functionality
		// errSend := utils.SendSMS(user.Phone, fmt.Sprintf("Your OTP code is: %s", otpCode))
		// if errSend != nil {
		//     logger.Error("AuthService:SendOTPChangePassword:SendSMS:Error:", errSend)
		//     return errors.NewAppError(errors.ErrInternalServer, "failed to send OTP SMS", errSend)
		// }
		logger.Info("SMS OTP sending not implemented yet. OTP code:", otpCode)
	}

	return nil
}

func (service *AuthService) ChangePassword(ctx context.Context, token string, requestData *dto.ChangePasswordRequest) *errors.AppError {

	// Check if token is blacklisted
	blacklisted, err := service.cache.IsTokenBlacklisted(ctx, token)
	if err != nil {
		logger.Error("AuthService:ChangePassword:IsTokenBlacklisted:Error:", err)
		return errors.NewAppError(errors.ErrInternalServer, "failed to check token blacklist", err)
	}
	if blacklisted {
		return errors.NewAppError(errors.ErrUnauthorized, "token is blacklisted", nil)
	}

	parseToken, err := utils.ValidateAndParseToken(token)
	if err != nil {
		logger.Error("AuthService:ChangePassword:ValidateAndParseToken", err)
		return errors.NewAppError(errors.ErrUnauthorized, "invalid token", nil)
	}

	// Check if user exists
	user, errGet := service.GetUserByIdentifier(ctx, utils.ToString(parseToken.UserID))
	if errGet != nil {
		logger.Error("AuthService:ChangePassword:GetUserByIdentifier:Error:", errGet)
		return errors.NewAppError(errors.ErrNotFound, "user not found", errGet)
	}

	// Check if password match
	if !utils.ComparePassword(user.Password, requestData.Password) {
		logger.Error("AuthService:ChangePassword:ComparePassword:Error:", err)
		return errors.NewAppError(errors.ErrUnauthorized, "user has invalid password", nil)
	}

	// Check OTP
	key := constants.RedisKeyOTPChangePassword + utils.ToString(parseToken.UserID)
	otp, err := service.cache.GetOTP(ctx, key)
	if err != nil {
		logger.Error("AuthService:ChangePassword:GetOTP:Error:", err)
		return errors.NewAppError(errors.ErrInternalServer, "failed to get OTP", err)
	}
	if otp != requestData.OTP {
		return errors.NewAppError(errors.ErrUnauthorized, "invalid OTP", nil)
	}

	// Update password
	hashedPassword, err := utils.HashPassword(requestData.NewPassword)
	if err != nil {
		logger.Error("AuthService:ChangePassword:HashPassword:Error:", err)
		return errors.NewAppError(errors.ErrInternalServer, "failed to hash password", err)
	}

	errUpdate := service.repo.PrivateUpdatePasswordUser(ctx, user.ID, hashedPassword)
	if errUpdate != nil {
		logger.Error("AuthService:ChangePassword:UpdateUser:Error:", errUpdate)
		return errors.NewAppError(errors.ErrInternalServer, "failed to change password", errUpdate)
	}
	// Invalid token
	errAdd := service.cache.AddToTokenBlacklist(ctx, token)
	if errAdd != nil {
		logger.Error("AuthService:ChangePassword:AddToBlacklist:Error:", errAdd)
		return errors.NewAppError(errors.ErrInternalServer, "failed to add token to blacklist", errAdd)
	}

	return nil
}

func (service *AuthService) ForgotPassword(ctx context.Context, identifier string) (*dto.ForgotPasswordResponse, *errors.AppError) {

	// Check if user exists
	user, err := service.GetUserByIdentifier(ctx, identifier)
	if err != nil || user == nil {
		logger.Error("AuthService:ForgotPassword:GetUserByIdentifier:Error:", err)
		return nil, errors.NewAppError(errors.ErrNotFound, "user not found", nil)
	}

	identifierType := utils.DetectIdentifierType(identifier)
	if identifierType == utils.IdentifierTypeUnknown {
		return nil, errors.NewAppError(errors.ErrUnauthorized, "invalid identifier", nil)
	}

	if identifierType == utils.IdentifierTypeEmail {
		// Check if user has verified their email
		if user.EmailVerifiedAt == nil {
			return nil, errors.NewAppError(errors.ErrUnauthorized, "user not verified email", nil)
		}

		otpCode := utils.GenerateOTP()

		// Tạo template data cho OTP
		data := utils.TemplateData{
			OTPCode: otpCode,
		}

		// Save OTP to cache
		errCache := service.cache.SetOTP(ctx, utils.ToString(user.ID), otpCode)
		if errCache != nil {
			logger.Error("AuthService:ForgotPassword:SetOTP:Error:", errCache)
			return nil, errors.NewAppError(errors.ErrInternalServer, "failed to save OTP", errCache)
		}

		// Gửi email với OTP template
		errSend := utils.SendTemplateEmailFromTemplatesDir(
			[]string{*user.Email},
			"Your OTP Code",
			"otp_email.html",
			data,
		)
		if errSend != nil {
			logger.Error("AuthService:ForgotPassword:SendTemplateEmailFromTemplatesDir:Error:", errSend)
			return nil, errors.NewAppError(errors.ErrInternalServer, "failed to send OTP email", errSend)
		}

	}

	if identifierType == utils.IdentifierTypePhone {
		if user.PhoneVerifiedAt == nil {
			return nil, errors.NewAppError(errors.ErrUnauthorized, "user not verified phone", nil)
		}
		// TODO: Implement SMS OTP sending
	}

	return &dto.ForgotPasswordResponse{
		UserId: user.ID,
	}, nil
}

func (service *AuthService) VerifyOTP(ctx context.Context, requestData *dto.VerifyOTPRequest) (*dto.VerifyOTPResponse, *errors.AppError) {
	// Get OTP from cache
	otp, err := service.cache.GetOTP(ctx, utils.ToString(requestData.UserID))
	if err != nil || otp == "" {
		logger.Error("AuthService:VerifyOTP:GetOTP:Error:", err)
		return nil, errors.NewAppError(errors.ErrInternalServer, "failed to get OTP from cache", err)
	}

	// Compare OTP
	if otp != requestData.OTP {
		return nil, errors.NewAppError(errors.ErrUnauthorized, "invalid OTP", nil)
	}

	resetPasswordToken, err := utils.GenerateToken(requestData.UserID, nil, nil, constants.ScopeTokenResetPassword)
	if err != nil {
		logger.Error("AuthService:VerifyOTP:GenerateToken:Error:", err)
		return nil, errors.NewAppError(errors.ErrInternalServer, "failed to generate token", err)
	}

	return &dto.VerifyOTPResponse{
		Token: resetPasswordToken,
	}, nil

}

func (service *AuthService) ResetPassword(ctx context.Context, requestData *dto.ResetPasswordRequest) *errors.AppError {

	// Check if token is blacklisted
	blacklisted, err := service.cache.IsTokenBlacklisted(ctx, requestData.Token)
	if err != nil {
		logger.Error("AuthService:ResetPassword:IsTokenBlacklisted:Error:", err)
		return errors.NewAppError(errors.ErrInternalServer, "failed to check token blacklist", err)
	}
	if blacklisted {
		return errors.NewAppError(errors.ErrUnauthorized, "token is blacklisted", nil)
	}

	tokenData, err := utils.ValidateAndParseToken(requestData.Token)
	if err != nil {
		logger.Error("AuthService:ResetPassword:ValidateAndParseToken:Error:", err)
		return errors.NewAppError(errors.ErrUnauthorized, "invalid token", nil)
	}

	if tokenData.Scope != constants.ScopeTokenResetPassword {
		return errors.NewAppError(errors.ErrUnauthorized, "invalid token", nil)
	}

	hashPassword, err := utils.HashPassword(requestData.NewPassword)
	if err != nil {
		logger.Error("AuthService:ResetPassword:HashPassword:Error:", err)
		return errors.NewAppError(errors.ErrInternalServer, "failed to hash password", err)
	}

	errUpdate := service.repo.PrivateUpdatePasswordUser(ctx, tokenData.UserID, hashPassword)
	if errUpdate != nil {
		logger.Error("AuthService:ResetPassword:UpdateUser:Error:", errUpdate)
		return errors.NewAppError(errors.ErrInternalServer, "failed to update user", errUpdate)
	}

	// Add token to blacklist
	errBlacklist := service.cache.AddToTokenBlacklist(ctx, requestData.Token)
	if errBlacklist != nil {
		logger.Error("AuthService:ResetPassword:AddToBlacklist:Error:", errBlacklist)
		return errors.NewAppError(errors.ErrInternalServer, "failed to add token to blacklist", errBlacklist)
	}

	return nil
}

func (service *AuthService) Logout(ctx context.Context, token string) *errors.AppError {
	// Add token to blacklist
	err := service.cache.AddToTokenBlacklist(ctx, token)
	if err != nil {
		logger.Error("AuthService:Logout:AddToBlacklist:Error:", err)
		return errors.NewAppError(errors.ErrInternalServer, "failed to add token to blacklist", err)
	}
	return nil
}

// Login authenticates a user with their identifier (phone/email) and password
// It implements login attempt blocking to prevent brute force attacks
func (service *AuthService) Login(ctx context.Context, requestData *dto.LoginRequest) (*dto.LoginResponse, *errors.AppError) {
	// Create unique key for tracking login attempts per user
	loginKey := fmt.Sprintf("login:%s", requestData.Identifier)

	// Check if user is currently blocked due to too many failed login attempts
	loginCount, err := service.cache.IsLoginBlocked(ctx, loginKey)
	if err != nil {
		logger.Error("AuthService:Login:IsLoginBlocked:Error:", err)
		return nil, errors.NewAppError(errors.ErrInternalServer, "failed to get login attempt", err)
	}

	// If user is blocked, refresh the block duration and return error
	if loginCount {
		errExpire := service.cache.Expire(ctx, loginKey, constants.BlockDuration)
		if errExpire != nil {
			logger.Error("AuthService:Login:Expire:Error:", errExpire)
			return nil, errors.NewAppError(errors.ErrInternalServer, "failed to expire login attempt", err)
		}
		return nil, errors.NewAppError(errors.ErrUnauthorized, "user is locked in 15 minite", nil)
	}

	// Retrieve user using identifier (phone or email)
	user, errGet := service.GetUserByIdentifier(ctx, requestData.Identifier)
	if errGet != nil || user == nil {
		return nil, errors.NewAppError(errors.ErrNotFound, "user not found", nil)
	}

	// Check if identifier is email or phone
	identifierType := utils.DetectIdentifierType(requestData.Identifier)
	if identifierType == utils.IdentifierTypeUnknown {
		return nil, errors.NewAppError(errors.ErrUnauthorized, "invalid identifier", nil)
	}

	if identifierType == utils.IdentifierTypeEmail {
		// Check if user has verified their email
		if user.EmailVerifiedAt == nil {
			errIncrement := service.cache.IncrementLoginAttempt(ctx, loginKey)
			if errIncrement != nil {
				logger.Error("AuthService:Login:IncrementLoginAttempt:Error:", errIncrement)
				return nil, errors.NewAppError(errors.ErrInternalServer, "failed to increment login attempt", err)
			}
			return nil, errors.NewAppError(errors.ErrUnauthorized, "user not verified email", nil)
		}
	}

	if identifierType == utils.IdentifierTypePhone {
		// Check if user has verified their phone
		if user.PhoneVerifiedAt == nil {
			errIncrement := service.cache.IncrementLoginAttempt(ctx, loginKey)
			if errIncrement != nil {
				logger.Error("AuthService:Login:IncrementLoginAttempt:Error:", errIncrement)
				return nil, errors.NewAppError(errors.ErrInternalServer, "failed to increment login attempt", err)
			}
			return nil, errors.NewAppError(errors.ErrUnauthorized, "user not verified phone", nil)
		}
	}

	// Check if user account is active
	if !user.IsActive {
		errIncrement := service.cache.IncrementLoginAttempt(ctx, loginKey)
		if errIncrement != nil {
			logger.Error("AuthService:Login:IncrementLoginAttempt:Error:", errIncrement)
			return nil, errors.NewAppError(errors.ErrInternalServer, "failed to increment login attempt", err)
		}
		return nil, errors.NewAppError(errors.ErrUnauthorized, "user not active", nil)
	}

	// Verify password - if incorrect, increment failed login attempts
	if !utils.ComparePassword(user.Password, requestData.Password) {
		//Increment failed login attempt counter
		errIncrement := service.cache.IncrementLoginAttempt(ctx, loginKey)
		if errIncrement != nil {
			logger.Error("AuthService:Login:IncrementLoginAttempt:Error:", errIncrement)
			return nil, errors.NewAppError(errors.ErrInternalServer, "failed to increment login attempt", err)
		}
		logger.Error("AuthService:Login:IncrementLoginAttempt:Error:", errIncrement)
		return nil, errors.NewAppError(errors.ErrUnauthorized, "incorrect password", nil)
	}

	// Generate JWT access token for API authentication
	accessToken, err := utils.GenerateToken(user.ID, user.Email, user.Username, constants.ScopeTokenAccess)
	if err != nil {
		return nil, errors.NewAppError(errors.ErrInternalServer, "failed to generate access token", err)
	}

	// Generate JWT refresh token for token renewal
	refreshToken, err := utils.GenerateToken(user.ID, user.Email, user.Username, constants.ScopeTokenRefresh)
	if err != nil {
		return nil, errors.NewAppError(errors.ErrInternalServer, "failed to generate refresh token", err)
	}

	// Clear any existing login attempts for this user
	errExpire := service.cache.Del(ctx, loginKey)
	if errExpire != nil {
		logger.Error("AuthService:Login:Expire:Error:", errExpire)
		return nil, errors.NewAppError(errors.ErrInternalServer, "failed to expire login attempt", err)
	}

	// Return successful login response with both tokens
	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (service *AuthService) GetUserByIdentifier(ctx context.Context, identifier string) (*dto.UserResponse, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultTimeout)
	defer cancel()

	// TODO: Implement cache user info by identifier
	result, err := service.repo.GetUserByIdentifier(ctx, identifier)
	if err != nil {
		return nil, errors.NewAppError(errors.ErrInternalServer, "failed to get user by identifier", err)
	}

	return mapper.ToUserDTO(result), nil
}

func (service *AuthService) Register(ctx context.Context, requestData *dto.RegisterRequest) (*dto.RegisterResponse, *errors.AppError) {
	// Check if user already exists
	existingUser, _ := service.GetUserByIdentifier(ctx, requestData.Phone)
	if existingUser != nil {
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "user with phone already exists", nil)
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(requestData.Password)
	if err != nil {
		return nil, errors.NewAppError(errors.ErrInternalServer, "failed to hash password", err)
	}

	// Create user entity
	userEntity := &entity.User{
		Phone:    requestData.Phone,
		Password: hashedPassword,
	}

	// Save user to database
	createdUser, err := service.repo.CreateUser(ctx, userEntity)
	if err != nil {
		return nil, errors.NewAppError(errors.ErrInternalServer, "failed to create user", err)
	}

	// Generate JWT tokens
	accessToken, err := utils.GenerateToken(createdUser.ID, createdUser.Email, createdUser.Username, constants.ScopeTokenAccess)
	if err != nil {
		return nil, errors.NewAppError(errors.ErrInternalServer, "failed to generate access token", err)
	}

	refreshToken, err := utils.GenerateToken(createdUser.ID, createdUser.Email, createdUser.Username, constants.ScopeTokenRefresh)
	if err != nil {
		return nil, errors.NewAppError(errors.ErrInternalServer, "failed to generate refresh token", err)
	}

	return &dto.RegisterResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (service *AuthService) RefreshToken(ctx context.Context, token string) (*dto.RefreshTokenResponse, *errors.AppError) {
	// TODO: Check if token is blacklisted
	isBlacklisted, errCheck := service.cache.IsTokenBlacklisted(ctx, token)
	if errCheck != nil {
		return nil, errors.NewAppError(errors.ErrUnauthorized, "failed to check token", nil)
	}
	if isBlacklisted {
		return nil, errors.NewAppError(errors.ErrUnauthorized, "token is blacklisted", nil)
	}

	user, err := utils.ValidateAndParseToken(token)
	if err != nil {
		return nil, errors.NewAppError(errors.ErrUnauthorized, "failed to parse token", nil)
	}

	result, err := service.repo.GetUserByIdentifier(ctx, user.ID)
	if err != nil {
		return nil, errors.NewAppError(errors.ErrInternalServer, "failed to get user by identifier", err)
	}

	// Generate new tokens
	accessToken, err := utils.GenerateToken(result.ID, result.Email, nil, constants.ScopeTokenAccess)
	if err != nil {
		return nil, errors.NewAppError(errors.ErrInternalServer, "failed to generate access token", err)
	}

	refreshToken, err := utils.GenerateToken(result.ID, result.Email, nil, constants.ScopeTokenRefresh)
	if err != nil {
		return nil, errors.NewAppError(errors.ErrInternalServer, "failed to generate refresh token", err)
	}

	// Add Refresh Token to Blacklist
	errAdd := service.cache.AddToTokenBlacklist(ctx, refreshToken)
	if errAdd != nil {
		return nil, errors.NewAppError(errors.ErrInternalServer, "failed to add refresh token to blacklist", errAdd)
	}

	// Return new tokens
	return &dto.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
