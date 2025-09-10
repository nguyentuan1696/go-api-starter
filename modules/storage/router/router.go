package router

import (
	"go-api-starter/core/middleware"
	"go-api-starter/modules/storage/controller"

	"github.com/labstack/echo/v4"
)

type StorageRouter struct {
	controller *controller.StorageController
}

func NewStorageRouter(controller *controller.StorageController) *StorageRouter {
	return &StorageRouter{controller: controller}
}

func (r *StorageRouter) Setup(e *echo.Echo, middleware *middleware.Middleware) {

	// API version 1 group - nhóm route cho phiên bản API v1
	v1 := e.Group("/api/v1")

	// Private routes group - nhóm route riêng tư yêu cầu xác thực
	privateRoutes := v1.Group("/private")

	// Storage routes - route cho module storage
	privateRoutes.POST("/storage/upload", r.controller.UploadToR2, middleware.AuthMiddleware(), middleware.PermissionMiddleware())

}
