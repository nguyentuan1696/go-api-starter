package entity

import (
	"go-api-starter/core/entity"
)

type Role struct {
	Name        string  `db:"name"`
	Slug        string  `db:"slug"`
	Description *string `db:"description"`
	IsSystem    bool    `db:"is_system"`
	IsActive    bool    `db:"is_active"`
	entity.BaseEntity
}

type PaginatedRoleEntity = entity.Pagination[Role]
