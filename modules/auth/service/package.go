package service

import (
	"go-api-starter/modules/auth/repository"

	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
)

type AuthService interface {
}

type authService struct {
	logger         *zerolog.Logger
	authRepository repository.AuthRepository
}

func NewAuthService(i do.Injector) (AuthService, error) {
	logger := do.MustInvoke[*zerolog.Logger](i)
	authRepository := do.MustInvoke[repository.AuthRepository](i)
	return &authService{
		logger:         logger,
		authRepository: authRepository,
	}, nil
}
