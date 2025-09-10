package storage

import (
	"go-api-starter/core/cache"
	"go-api-starter/core/database"
	"go-api-starter/core/middleware"
	"go-api-starter/modules/storage/controller"
	"go-api-starter/modules/storage/repository"
	"go-api-starter/modules/storage/router"
	"go-api-starter/modules/storage/service"

	authRepository "go-api-starter/modules/auth/repository"
	authService "go-api-starter/modules/auth/service"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, db database.Database, r2Client *s3.Client, cache cache.Cache) {
	repository := repository.NewStorageRepository(db)
	storageService := service.NewStorageService(repository, r2Client)
	controller := controller.NewStorageController(storageService)
	authRepository := authRepository.NewAuthRepository(db)
	authService := authService.NewAuthService(authRepository, cache)
	middleware := middleware.NewMiddleware(authService)

	router.NewStorageRouter(controller).Setup(e, middleware)
}
