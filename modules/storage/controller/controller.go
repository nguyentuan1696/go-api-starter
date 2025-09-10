package controller

import (
	"go-api-starter/core/controller"
	"go-api-starter/modules/storage/service"
)

type StorageController struct {
	controller.BaseController
	service service.StorageServiceInterface
}

func NewStorageController(service service.StorageServiceInterface) *StorageController {
	return &StorageController{
		BaseController: controller.NewBaseController(),
		service:        service,
	}
}
