package entity

import (
	"go-api-starter/core/entity"
)

type Action string

const (
	ActionCreate Action = "create"
	ActionRead   Action = "read"
	ActionUpdate Action = "update"
	ActionDelete Action = "delete"
)

type Permission struct {
	Name        string  `db:"name"`
	Slug        string  `db:"slug"`
	Resource    string  `db:"resource"`
	Action      Action  `db:"action"`
	Description *string `db:"description"`
	IsSystem    bool    `db:"is_system"`
	entity.BaseEntity
}

type PaginatedPermissionEntity = entity.Pagination[Permission]
