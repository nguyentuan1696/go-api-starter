package entity

import (
	"time"
	"go-api-starter/core/entity"

	"github.com/google/uuid"
)

type Storage struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	URL       string    `db:"url"`
	Size      int64     `db:"size"`
	Extension string    `db:"extension"`
	Uploaded  time.Time `db:"uploaded"`
}

type PaginatedStorageResponse = entity.Pagination[Storage]
