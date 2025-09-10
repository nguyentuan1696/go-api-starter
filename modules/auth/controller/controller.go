package controller

import (
	"go-api-starter/core/controller"
	"go-api-starter/modules/auth/service"
)

type AuthController struct {
	controller.BaseController
	AuthService service.AuthServiceInterface
}

func NewAuthController(service service.AuthServiceInterface) *AuthController {
	return &AuthController{
		BaseController: controller.NewBaseController(),
		AuthService:    service,
	}
}
