package handler

import (
	"go-api-starter/modules/auth/service"
	baseHandler "go-api-starter/pkg/handler"

	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
)

type AuthHTTPHandler struct {
	logger      *zerolog.Logger
	baseHandler baseHandler.BaseHandler
	service     service.AuthService
}

func NewAuthHTTPHandler(i do.Injector) (*AuthHTTPHandler, error) {
	logger := do.MustInvoke[*zerolog.Logger](i)
	service := do.MustInvoke[service.AuthService](i)
	return &AuthHTTPHandler{
		logger:  logger,
		service: service,
	}, nil
}
