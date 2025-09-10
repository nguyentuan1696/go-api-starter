package router

import (
	"go-api-starter/core/middleware"
	"go-api-starter/modules/auth/controller"

	"github.com/labstack/echo/v4"
)

type AuthRouter struct {
	AuthController controller.AuthController
}

func NewAuthRouter(authController controller.AuthController) *AuthRouter {
	return &AuthRouter{
		AuthController: authController,
	}
}

func (r *AuthRouter) Setup(e *echo.Echo, middleware *middleware.Middleware) {
	v1 := e.Group("/api/v1")

	privateRoutes := v1.Group("/private")
	publicRoutes := v1.Group("/public")

	authRoutes := privateRoutes.Group("/auth")
	authPublicRoutes := publicRoutes.Group("/auth")

	authPublicRoutes.POST("/register", r.AuthController.Register)
	authPublicRoutes.POST("/login", r.AuthController.Login)
	authPublicRoutes.POST("/logout", r.AuthController.Logout)

	authPublicRoutes.POST("/forgot-password", r.AuthController.ForgotPassword)
	authPublicRoutes.POST("/verify-otp", r.AuthController.VerifyOTP)
	authPublicRoutes.POST("/reset-password", r.AuthController.ResetPassword)
	authPublicRoutes.POST("/send-otp-change-password", r.AuthController.SendOTPChangePassword)
	authPublicRoutes.POST("/change-password", r.AuthController.ChangePassword)

	authPublicRoutes.POST("/update-password", r.AuthController.ChangePassword, middleware.AuthMiddleware())
	authPublicRoutes.PUT("user-profile", r.AuthController.UpdateUserProfile, middleware.AuthMiddleware())

	authRoutes.POST("/roles", r.AuthController.PrivateCreateRole, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	authRoutes.GET("/roles", r.AuthController.PrivateGetRoles, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	authRoutes.GET("/roles/:id", r.AuthController.PrivateGetRoleByID, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	authRoutes.PUT("/roles/:id", r.AuthController.PrivateUpdateRole, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	authRoutes.DELETE("/roles/:id", r.AuthController.PrivateDeleteRole, middleware.AuthMiddleware(), middleware.PermissionMiddleware())

	authRoutes.POST("/permissions", r.AuthController.PrivateCreatePermission, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	authRoutes.GET("/permissions", r.AuthController.PrivateGetPermissions, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	authRoutes.GET("/permissions/:id", r.AuthController.PrivateGetPermissionByID, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	authRoutes.PUT("/permissions/:id", r.AuthController.PrivateUpdatePermission, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	authRoutes.DELETE("/permissions/:id", r.AuthController.PrivateDeletePermission, middleware.AuthMiddleware(), middleware.PermissionMiddleware())

	authRoutes.POST("/user-roles", r.AuthController.PrivateAssignRoleToUser, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	authRoutes.POST("/role-permissions", r.AuthController.PrivateAssignPermissionToRole, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	authRoutes.POST("/user-permissions", r.AuthController.PrivateAssignPermissionToUser, middleware.AuthMiddleware(), middleware.PermissionMiddleware())

	authRoutes.GET("/users", r.AuthController.PrivateGetUsers, middleware.AuthMiddleware(), middleware.PermissionMiddleware("users:read"))
	authRoutes.GET("/users/:id", r.AuthController.PrivateGetUser, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	authRoutes.PUT("/users/:id", r.AuthController.PrivateUpdateUser, middleware.AuthMiddleware(), middleware.PermissionMiddleware())

}
