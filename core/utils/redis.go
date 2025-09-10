package utils

import "go-api-starter/core/constants"

// GenerateRolePermissionsKey tạo Redis key cho permissions của role
func GenerateRolePermissionsKey(roleSlug string) string {
	return constants.RedisKeyPrefix + "role:" + roleSlug + ":permissions"
}

// GenerateUserPermissionsKey tạo Redis key cho permissions của user
func GenerateUserPermissionsKey(userID string) string {
	return constants.RedisKeyPrefix + "user:" + userID + ":permissions"
}

// GenerateUserRolesKey tạo Redis key cho roles của user
func GenerateUserRolesKey(userID string) string {
	return constants.RedisKeyPrefix + "user:" + userID + ":roles"
}