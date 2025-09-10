package auth

import (
	"go-api-starter/core/cache"
	"go-api-starter/core/database"
	"go-api-starter/core/middleware"
	"go-api-starter/modules/auth/controller"
	"go-api-starter/modules/auth/repository"
	"go-api-starter/modules/auth/router"
	"go-api-starter/modules/auth/service"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, db database.Database, cache cache.Cache) {
	repository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(repository, cache)
	controller := controller.NewAuthController(authService)
	middleware := middleware.NewMiddleware(authService)

	router.NewAuthRouter(*controller).Setup(e, middleware)
}
