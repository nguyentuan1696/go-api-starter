package product

import (
	"go-api-starter/core/cache"
	"go-api-starter/core/database"
	"go-api-starter/core/middleware"
	authRepository "go-api-starter/modules/auth/repository"
	authService "go-api-starter/modules/auth/service"
	"go-api-starter/modules/product/controller"
	"go-api-starter/modules/product/repository"
	"go-api-starter/modules/product/router"
	"go-api-starter/modules/product/service"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, db database.Database, cache cache.Cache) {
	repository := repository.NewProductRepository(db)
	service := service.NewProductService(repository)
	controller := controller.NewProductController(service)
	authRepository := authRepository.NewAuthRepository(db)
	authService := authService.NewAuthService(authRepository, cache)
	middleware := middleware.NewMiddleware(authService)

	router.NewProductRouter(*controller).Setup(e, middleware)
}
