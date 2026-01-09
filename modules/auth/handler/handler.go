package handler

import (
	"go-api-starter/modules/auth/service"
	"go-api-starter/pkg/controller"
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
