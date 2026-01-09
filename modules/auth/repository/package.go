package repository

import (
	"go-api-starter/pkg/database"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
)

type AuthRepository interface {
}

type authRepository struct {
	db     *pgxpool.Pool   `do:""`
	logger *zerolog.Logger `do:""`
}

func NewAuthRepository(injector do.Injector) (AuthRepository, error) {
	db := do.MustInvoke[*database.Postgresql](injector)
	logger := do.MustInvoke[*zerolog.Logger](injector)

	return &authRepository{db: db.Pool(), logger: logger}, nil
}
