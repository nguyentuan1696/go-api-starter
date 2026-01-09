package utils

import (
	"strings"

	"github.com/google/uuid"
)

// TrimSpace removes leading, trailing, and multiple spaces between words
func TrimSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// TrimSpacePointer handles string pointer and returns trimmed string pointer
func TrimSpacePointer(s *string) *string {
	if s == nil {
		return nil
	}
	trimmed := TrimSpace(*s)
	return &trimmed
}

// TrimAllSpaces removes all spaces from a string
func TrimAllSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

// IsEmpty checks if a string is empty after trimming spaces
func IsEmpty(s string) bool {
	return TrimSpace(s) == ""
}

// ToNumber converts string to number, returns 0 if conversion fails
func ToNumber(s string) int {
	// Remove all spaces
	s = TrimAllSpaces(s)

	// Convert string to number
	var result int
	for _, ch := range s {
		// Check if character is digit
		if ch < '0' || ch > '9' {
			return 0
		}
		// Build number
		result = result*10 + int(ch-'0')
	}

	return result
}

// ToNumberWithDefault converts string to number with a default value
func ToNumberWithDefault(s string, defaultValue int) int {
	if IsEmpty(s) {
		return defaultValue
	}
	return ToNumber(s)
}

func ToString(s uuid.UUID) string {
	return s.String()
}

func ToUUID(s string) uuid.UUID {
	return uuid.MustParse(s)
}

// IdentifierType represents the type of identifier
type IdentifierType string

const (
	IdentifierTypeEmail    IdentifierType = "email"
	IdentifierTypePhone    IdentifierType = "phone"
	IdentifierTypeUsername IdentifierType = "username"
	IdentifierTypeUnknown  IdentifierType = "unknown"
)

// DetectIdentifierType determines if the input string is an email, phone number, or username
func DetectIdentifierType(identifier string) IdentifierType {
	// Trim spaces first
	identifier = TrimSpace(identifier)
	
	// Check if empty
	if IsEmpty(identifier) {
		return IdentifierTypeUnknown
	}
	
	// Check if it's a valid email
	if IsValidEmail(identifier) {
		return IdentifierTypeEmail
	}
	
	// Check if it's a valid phone number
	if IsValidPhone(identifier) {
		return IdentifierTypePhone
	}
	
	// If it's neither email nor phone, consider it as username
	// You can add additional username validation rules here if needed
	if len(identifier) >= 3 && len(identifier) <= 50 {
		return IdentifierTypeUsername
	}
	
	return IdentifierTypeUnknown
}

// IsEmail checks if the identifier is an email
func IsEmail(identifier string) bool {
	return DetectIdentifierType(identifier) == IdentifierTypeEmail
}

// IsPhone checks if the identifier is a phone number
func IsPhone(identifier string) bool {
	return DetectIdentifierType(identifier) == IdentifierTypePhone
}

// IsUsername checks if the identifier is a username
func IsUsername(identifier string) bool {
	return DetectIdentifierType(identifier) == IdentifierTypeUsername
}
