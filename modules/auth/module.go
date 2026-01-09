package auth

import (
	handler "go-api-starter/modules/auth/handler/http"
	repository "go-api-starter/modules/auth/repository"
	router "go-api-starter/modules/auth/router/http"
	service "go-api-starter/modules/auth/service"

	"github.com/samber/do/v2"
)

var Package = do.Package(
	do.Lazy(repository.NewAuthRepository),
	do.Lazy(service.NewAuthService),
	do.Lazy(handler.NewAuthHTTPHandler),
	do.Lazy(router.NewAuthRouter),
)
