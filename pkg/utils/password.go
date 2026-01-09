package utils

import (
	"errors"
	"regexp"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

var (
	reLower   = regexp.MustCompile(`\p{Ll}`)       // chữ thường (Unicode)
	reUpper   = regexp.MustCompile(`\p{Lu}`)       // chữ hoa (Unicode)
	reDigit   = regexp.MustCompile(`\p{Nd}`)       // chữ số (Unicode)
	reSpecial = regexp.MustCompile(`[\p{P}\p{S}]`) // ký tự đặc biệt: Punctuation + Symbol
	reSpace   = regexp.MustCompile(`\s`)           // khoảng trắng
)

func ValidateStrongPassword(pw string) error {
	n := utf8.RuneCountInString(pw)
	if n < 8 || n > 32 {
		return errors.New("password phải dài từ 8-32 ký tự")
	}
	if reSpace.MatchString(pw) {
		return errors.New("password không được chứa khoảng trắng")
	}
	if !reLower.MatchString(pw) {
		return errors.New("cần ít nhất 1 chữ thường")
	}
	if !reUpper.MatchString(pw) {
		return errors.New("cần ít nhất 1 chữ hoa")
	}
	if !reDigit.MatchString(pw) {
		return errors.New("cần ít nhất 1 chữ số")
	}
	if !reSpecial.MatchString(pw) {
		return errors.New("cần ít nhất 1 ký tự đặc biệt")
	}
	return nil
}

// HashPassword takes a plain text password and returns a bcrypt hashed version
// The cost parameter determines how computationally expensive the hash will be
// Default cost is 10, but you can adjust based on your security requirements
func HashPassword(password string) (string, error) {
	// Generate a bcrypt hash with default cost (10)
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

// ComparePassword compares a hashed password with a plain text password
// Returns nil if they match, error otherwise
func ComparePassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
