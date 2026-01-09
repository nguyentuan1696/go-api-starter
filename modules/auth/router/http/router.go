package router

import (
	authHandler "go-api-starter/modules/auth/handler/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
)

type AuthHTTPRouter struct {
	handler *authHandler.AuthHTTPHandler
}

func NewAuthRouter(i do.Injector) (*AuthHTTPRouter, error) {
	h := do.MustInvoke[*authHandler.AuthHTTPHandler](i)
	return &AuthHTTPRouter{
		handler: h,
	}, nil
}

func (r *AuthHTTPRouter) Register(e *echo.Echo) {
	r.registerPublicRoutes(e)
	r.registerInternalRoutes(e)
}

func (r *AuthHTTPRouter) registerPublicRoutes(e *echo.Echo) {
	// group := e.Group("/api/v1/auth")
	// Add public routes here
}

func (r *AuthHTTPRouter) registerInternalRoutes(e *echo.Echo) {
	// group := e.Group("/internal/api/v1/auth")
	// Add internal routes here
}
