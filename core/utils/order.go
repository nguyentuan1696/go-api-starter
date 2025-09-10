package utils

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// Các bộ ký tự có thể sử dụng
const (
	// Chỉ số và chữ hoa (dễ đọc, tránh nhầm lẫn)
	AlphanumericUppercase = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func GenerateOrderNumber() string {
	// Sử dụng bộ ký tự tùy chỉnh
	id, _ := gonanoid.Generate(AlphanumericUppercase, 10)
	return "TINUP" + id
}
