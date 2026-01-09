package middleware

import (
	"go-api-starter/core/constants"
	"go-api-starter/core/controller"
	"go-api-starter/core/errors"
	"go-api-starter/core/logger"
	"go-api-starter/core/utils"
	"go-api-starter/modules/auth/service"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type Middleware struct {
	controller.BaseController
	AuthService service.AuthServiceInterface
}

func NewMiddleware(authService service.AuthServiceInterface) *Middleware {
	return &Middleware{
		BaseController: controller.NewBaseController(),
		AuthService:    authService,
	}
}
func (m *Middleware) AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get token from header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return m.Unauthorized(errors.ErrMissingAuthorizationHeader, "missing authorization header")
			}

			// Check Bearer token format
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return m.Unauthorized(errors.ErrInvalidTokenFormat, "invalid token format")
			}

			// Validate token
			claims, err := utils.ValidateAndParseToken(parts[1])
			if err != nil {
				logger.Error("AuthMiddleware:ValidateAndParseToken:Error:", err)
				return m.Unauthorized(errors.ErrInvalidTokenFormat, "invalid token: "+err.Error())
			}

			// Set user claims in context
			c.Set(constants.ContextTokenData, claims)
			return next(c)
		}
	}
}

func (m *Middleware) PermissionMiddleware(requiredPermissions ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Lấy thông tin user từ context (đã được set bởi AuthMiddleware)
			userData := c.Get(constants.ContextTokenData)
			if userData == nil {
				return m.Unauthorized(errors.ErrMissingAuthorizationHeader, "missing authorization header")
			}

			// Parse user claims từ token
			claims, ok := userData.(*utils.TokenClaims)
			if !ok {
				return m.Unauthorized(errors.ErrInvalidTokenFormat, "invalid token format")
			}

			// Nếu không có permission nào được yêu cầu, cho phép truy cập
			if len(requiredPermissions) == 0 {
				return next(c)
			}

			// Lấy danh sách permissions trực tiếp của user (bao gồm từ roles và permissions riêng)
			ctx := c.Request().Context()

			// Thử lấy từ cache trước
			userPermissions, err := m.AuthService.PrivateGetPermissionsByUserIDFromCache(ctx, claims.UserID)
			if err != nil {
				// Nếu cache không có hoặc lỗi, lấy từ database
				logger.Info("PermissionMiddleware:Cache miss, fetching from database:", err)
				userPermissions, err = m.AuthService.PrivateGetPermissionsByUserID(ctx, claims.UserID)
				if err != nil {
					logger.Error("PermissionMiddleware:PrivateGetPermissions:Error:", err)
					return m.Unauthorized(errors.ErrInternalServer, "internal http error")
				}
			}

			if userPermissions == nil {
				logger.Error("PermissionMiddleware:PrivateGetPermissions:Error: userPermissions is nil")
				return m.Unauthorized(errors.ErrInternalServer, "internal http error")
			}

			requiredSet := make(map[string]struct{}, len(requiredPermissions))
			for _, rp := range requiredPermissions {
				requiredSet[rp] = struct{}{}
			}

			for _, up := range *userPermissions {
				if _, ok := requiredSet[up.Resource+":"+string(up.Action)]; ok {
					return next(c)
				}
			}

			// Nếu không có permission nào khớp, trả về lỗi Forbidden
			return m.Unauthorized(errors.ErrForbidden, "forbidden")
		}
	}
}

func CORSMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
			c.Response().Header().Set("Access-Control-Allow-Credentials", "true")

			if c.Request().Method == "OPTIONS" {
				return c.NoContent(http.StatusOK)
			}

			return next(c)
		}
	}
}

func LoggerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Record start time
			start := time.Now()

			req := c.Request()
			res := c.Response()

			if err := next(c); err != nil {
				c.Error(err)
			}

			// Calculate latency
			latency := time.Since(start)

			// Log request details
			logger.Info("Request",
				"method", req.Method,
				"uri", req.RequestURI,
				"status", res.Status,
				"remote_ip", c.RealIP(),
				"user_agent", req.UserAgent(),
				"latency", latency.String(),
			)

			return nil
		}
	}
}
