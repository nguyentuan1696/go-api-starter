package utils

import "regexp"

func IsValidPhone(phone string) bool {
	// Vietnamese phone number regex pattern
	// Matches: +84, 0084, or 0 followed by valid Vietnamese mobile prefixes
	// Valid prefixes: 2, 3, 5, 7, 8, 9 followed by 1-2 digits, then 7 more digits
	pattern := `^(?:\+84|0084|0)[235789][0-9]{1,2}[0-9]{7}$`

	regex := regexp.MustCompile(pattern)
	return regex.MatchString(phone)
}
