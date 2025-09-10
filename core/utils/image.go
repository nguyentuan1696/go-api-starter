package utils

import (
	"fmt"
	"mime/multipart"
	"strings"
	"go-api-starter/core/constants"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

func IsValidImageType(contentType string) bool {
	allowedTypes := strings.Split(constants.UploadAllowedTypes, ",")
	
	for _, allowedType := range allowedTypes {
		if strings.TrimSpace(allowedType) == contentType {
			return true
		}
	}
	
	return false
}

// ValidateFileSize kiểm tra kích thước file có vượt quá giới hạn không
func ValidateFileSize(fileSize int64, maxSizeBytes int64) bool {
	return fileSize <= maxSizeBytes
}

// GetMaxFileSizeBytes trả về kích thước tối đa tính bằng bytes từ constants
func GetMaxFileSizeBytes() int64 {
	return int64(constants.UploadMaxSize)
}

func GetExtensionFromContentType(contentType string) string {
	switch contentType {
	case "image/jpeg", "image/jpg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "image/gif":
		return ".gif"
	case "image/webp":
		return ".webp"
	default:
		return ".jpg"
	}
}

func GenerateFileName(originalName string, extension string) string {
	// Remove extension from original name
	originalName = strings.TrimSuffix(originalName, extension)
	originalName = slug.Make(originalName)
	shortUUID := uuid.New().String()
	return fmt.Sprintf("images/%s-%s%s", originalName, shortUUID, extension)
}

func ValidateUploadFile(fileHeader *multipart.FileHeader) error {
	// Validate file không rỗng
	if fileHeader.Size == 0 {
		return fmt.Errorf("empty file not allowed")
	}

	// Kiểm tra loại file
	contentType := fileHeader.Header.Get("Content-Type")
	if contentType == "" {
		return fmt.Errorf("missing content type")
	}

	if !IsValidImageType(contentType) {
		return fmt.Errorf("invalid image type. Allowed: %s", constants.UploadAllowedTypes)
	}

	// Check file size using constants
	maxSizeBytes := GetMaxFileSizeBytes()
	if !ValidateFileSize(fileHeader.Size, maxSizeBytes) {
		maxSizeMB := maxSizeBytes / (1024 * 1024)
		return fmt.Errorf("file size exceeds %dMB limit", maxSizeMB)
	}

	return nil
}
