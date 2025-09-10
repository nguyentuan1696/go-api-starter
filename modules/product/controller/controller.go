package controller

import (
	"go-api-starter/core/controller"
	"go-api-starter/modules/product/service"
)

type ProductController struct {
	controller.BaseController
	ProductService service.ProductServiceInterface
}

func NewProductController(service service.ProductServiceInterface) *ProductController {
	return &ProductController{
		BaseController: controller.NewBaseController(),
		ProductService: service,
	}
}
