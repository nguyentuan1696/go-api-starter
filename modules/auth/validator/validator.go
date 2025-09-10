package validator

import (
	"strings"
	"go-api-starter/core/utils"
	"go-api-starter/core/validation"
	"go-api-starter/modules/auth/dto"
)

func ValidateVerifyOTPRequest(req *dto.VerifyOTPRequest) *validation.ValidationResult {
	if req == nil {
		return nil
	}

	result := validation.NewValidationResult()

	if utils.IsEmpty(utils.ToString(req.UserID)) {
		result.AddError("user_id", "User ID is required")
	}

	if utils.IsEmpty(req.OTP) {
		result.AddError("otp", "OTP is required")
	}

	return result
}

func ValidateAssignPermissionToUserRequest(req *dto.UserPermissionRequest) *validation.ValidationResult {
	if req == nil {
		return nil
	}

	result := validation.NewValidationResult()

	if utils.IsEmpty(utils.ToString(req.UserID)) {
		result.AddError("user_id", "User ID is required")
	}

	if utils.IsEmpty(utils.ToString(req.PermissionID)) {
		result.AddError("permission_id", "Permission ID is required")
	}

	return result
}

func ValidateAssignPermissionToRoleRequest(req *dto.RolePermissionRequest) *validation.ValidationResult {
	if req == nil {
		return nil
	}

	result := validation.NewValidationResult()

	if utils.IsEmpty(utils.ToString(req.RoleID)) {
		result.AddError("role_id", "Role ID is required")
	}

	if len(req.PermissionID) == 0 {
		result.AddError("permission_id", "Permission ID is required")
	}

	return result
}

func ValidateAssignRoleToUserRequest(req *dto.UserRoleRequest) *validation.ValidationResult {
	if req == nil {
		return nil
	}

	result := validation.NewValidationResult()

	if utils.IsEmpty(utils.ToString(req.UserID)) {
		result.AddError("user_id", "User ID is required")
	}

	if utils.IsEmpty(utils.ToString(req.RoleID)) {
		result.AddError("role_id", "Role ID is required")
	}

	return result
}

func ValidateForgotPasswordRequest(req *dto.ForgotPasswordRequest) *validation.ValidationResult {
	if req == nil {
		return nil
	}

	result := ValidateIdentifier(req.Identifier)

	return result
}

func ValidateResetPasswordRequest(req *dto.ResetPasswordRequest) *validation.ValidationResult {
	if req == nil {
		return nil
	}

	result := validation.NewValidationResult()

	if utils.IsEmpty(req.NewPassword) {
		result.AddError("password", "Password is required")
	}

	if utils.IsEmpty(req.NewPassword) {
		result.AddError("new_password", "New password is required")
	}

	if req.NewPassword != req.ConfirmPassword {
		result.AddError("confirm_password", "Confirm password not match")
	}

	return result
}

func ValidateChangePasswordRequest(req *dto.ChangePasswordRequest) *validation.ValidationResult {
	if req == nil {
		return nil
	}

	result := validation.NewValidationResult()

	if utils.IsEmpty(req.Password) {
		result.AddError("password", "Password is required")
	}

	if utils.IsEmpty(req.NewPassword) {
		result.AddError("new_password", "New password is required")
	}

	if req.Password == req.NewPassword {
		result.AddError("new_password", "New password required not the same as old password")
	}

	if req.NewPassword != req.ConfirmPassword {
		result.AddError("confirm_password", "Confirm password not match")
	}

	return result

}

func ValidateLoginRequest(req *dto.LoginRequest) *validation.ValidationResult {
	if req == nil {
		return nil
	}

	result := ValidateIdentifier(req.Identifier)

	return result

}

func ValidateRegisterRequest(req *dto.RegisterRequest) *validation.ValidationResult {
	if req == nil {
		return nil
	}

	result := validation.NewValidationResult()

	if utils.IsEmpty(req.Phone) {
		result.AddError("phone", "Phone is required")
	}

	if !utils.IsValidPhone(req.Phone) {
		result.AddError("phone", "Phone is invalid")
	}

	if err := utils.ValidateStrongPassword(req.Password); err != nil {
		result.AddError("password", err.Error())
	}

	return result
}

func ValidateRoleRequest(req *dto.RoleRequest) *validation.ValidationResult {
	if req == nil {
		return nil
	}
	result := validation.NewValidationResult()

	if utils.IsEmpty(req.Name) {
		result.AddError("name", "Name is required")
	}

	return result
}

func ValidatePermissionRequest(req *dto.PermissionRequest) *validation.ValidationResult {
	if req == nil {
		return nil
	}
	result := validation.NewValidationResult()

	if utils.IsEmpty(req.Name) {
		result.AddError("name", "Name is required")
	}

	if utils.IsEmpty(req.Resource) {
		result.AddError("resource", "Resource is required")
	}

	if utils.IsEmpty(req.Action) {
		result.AddError("action", "Action is required")
	}

	return result
}

// ValidateIdentifier validates a string as either phone number or email
// Returns validation result with appropriate error messages
func ValidateIdentifier(identifier string) *validation.ValidationResult {
	result := validation.NewValidationResult()

	if utils.IsEmpty(identifier) {
		result.AddError("identifier", "Identifier is required")
		return result
	}

	// Check if identifier contains @ symbol to determine if it's an email
	if strings.Contains(identifier, "@") {
		// Validate as email
		if !utils.IsValidEmail(identifier) && !utils.IsValidEmailDomain(identifier) {
			result.AddError("identifier", "Invalid email format")
		}
	} else {
		// Validate as phone number
		if !utils.IsValidPhone(identifier) {
			result.AddError("identifier", "Invalid phone number format")
		}
	}

	return result
}
