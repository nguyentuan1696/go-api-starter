package router

import (
	"go-api-starter/core/middleware"
	"go-api-starter/modules/product/controller"

	"github.com/labstack/echo/v4"
)

type ProductRouter struct {
	ProductController controller.ProductController
}

func NewProductRouter(controller controller.ProductController) *ProductRouter {
	return &ProductRouter{
		ProductController: controller,
	}
}

func (r *ProductRouter) Setup(e *echo.Echo, middleware *middleware.Middleware) {
	// API version 1 group - nhóm route cho phiên bản API v1
	v1 := e.Group("/api/v1")

	// Private routes group - nhóm route riêng tư yêu cầu xác thực
	privateRoutes := v1.Group("/private")
	// Public routes group - nhóm route công khai không yêu cầu xác thực
	publicRoutes := v1.Group("/public")

	// Product routes groups
	products := privateRoutes.Group("/products")
	publicProducts := publicRoutes.Group("/products")

	// =====================================================
	// PRIVATE ROUTES - Yêu cầu xác thực
	// =====================================================

	// Benefits routes - Quản lý công dụng sản phẩm
	products.POST("/benefits", r.ProductController.PrivateCreateBenefit, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.GET("/benefits", r.ProductController.PrivateGetBenefits, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.GET("/benefits/:id", r.ProductController.PrivateGetBenefitById, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.PUT("/benefits/:id", r.ProductController.PrivateUpdateBenefit, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.DELETE("/benefits/:id", r.ProductController.PrivateDeleteBenefit, middleware.AuthMiddleware(), middleware.PermissionMiddleware())

	// Brands routes - Quản lý thương hiệu
	products.POST("/brands", r.ProductController.PrivateCreateBrand, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.GET("/brands", r.ProductController.PrivateGetBrands, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.GET("/brands/:id", r.ProductController.PrivateGetBrandById, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.PUT("/brands/:id", r.ProductController.PrivateUpdateBrand, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.DELETE("/brands/:id", r.ProductController.PrivateDeleteBrand, middleware.AuthMiddleware(), middleware.PermissionMiddleware())

	// Categories routes - Quản lý danh mục sản phẩm
	products.POST("/categories", r.ProductController.PrivateCreateCategory, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.GET("/categories", r.ProductController.PrivateGetCategories, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.GET("/categories/:id", r.ProductController.PrivateGetCategoryById, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.PUT("/categories/:id", r.ProductController.PrivateUpdateCategory, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.DELETE("/categories/:id", r.ProductController.PrivateDeleteCategory, middleware.AuthMiddleware(), middleware.PermissionMiddleware())

	// Ingredients routes - Quản lý thành phần
	products.POST("/ingredients", r.ProductController.PrivateCreateIngredient, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.GET("/ingredients", r.ProductController.PrivateGetIngredients, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.GET("/ingredients/:id", r.ProductController.PrivateGetIngredientById, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.PUT("/ingredients/:id", r.ProductController.PrivateUpdateIngredient, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.DELETE("/ingredients/:id", r.ProductController.PrivateDeleteIngredient, middleware.AuthMiddleware(), middleware.PermissionMiddleware())

	// Skin Types routes - Quản lý loại da
	products.POST("/skin-types", r.ProductController.PrivateCreateSkinType, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.GET("/skin-types", r.ProductController.PrivateGetSkinTypes, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.GET("/skin-types/:id", r.ProductController.PrivateGetSkinTypeById, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.PUT("/skin-types/:id", r.ProductController.PrivateUpdateSkinType, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.DELETE("/skin-types/:id", r.ProductController.PrivateDeleteSkinType, middleware.AuthMiddleware(), middleware.PermissionMiddleware())

	// Tags routes - Quản lý thẻ tag
	products.POST("/tags", r.ProductController.PrivateCreateTag, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.GET("/tags", r.ProductController.PrivateGetTags, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.GET("/tags/:id", r.ProductController.PrivateGetTagById, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.PUT("/tags/:id", r.ProductController.PrivateUpdateTag, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.DELETE("/tags/:id", r.ProductController.PrivateDeleteTag, middleware.AuthMiddleware(), middleware.PermissionMiddleware())

	// Products routes - Quản lý sản phẩm
	products.POST("/items", r.ProductController.PrivateCreateProduct, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.GET("/items", r.ProductController.PrivateGetProducts, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.GET("/items/:id", r.ProductController.PrivateGetProductById, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.PUT("/items/:id", r.ProductController.PrivateUpdateProduct, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.DELETE("/items/:id", r.ProductController.PrivateDeleteProduct, middleware.AuthMiddleware(), middleware.PermissionMiddleware())

	// Payment Methods routes - Quản lý phương thức thanh toán
	products.POST("/payment-methods", r.ProductController.PrivateCreatePaymentMethod, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.GET("/payment-methods", r.ProductController.PrivateGetPaymentMethods, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.GET("/payment-methods/:id", r.ProductController.PrivateGetPaymentMethodById, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.PUT("/payment-methods/:id", r.ProductController.PrivateUpdatePaymentMethod, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.DELETE("/payment-methods/:id", r.ProductController.PrivateDeletePaymentMethod, middleware.AuthMiddleware(), middleware.PermissionMiddleware())

	// Shipping Methods routes - Quản lý phương thức vận chuyển
	products.POST("/shipping-methods", r.ProductController.PrivateCreateShippingMethod, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.GET("/shipping-methods", r.ProductController.PrivateGetShippingMethods, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.GET("/shipping-methods/:id", r.ProductController.PrivateGetShippingMethodById, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.PUT("/shipping-methods/:id", r.ProductController.PrivateUpdateShippingMethod, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.DELETE("/shipping-methods/:id", r.ProductController.PrivateDeleteShippingMethod, middleware.AuthMiddleware(), middleware.PermissionMiddleware())

	// Orders routes - Quản lý đơn hàng
	products.GET("/orders", r.ProductController.PrivateGetOrders, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.GET("/orders/:id", r.ProductController.PrivateGetOrderById, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	products.GET("/orders/:id/items", r.ProductController.PrivateGetOrderDetailWithItems, middleware.AuthMiddleware(), middleware.PermissionMiddleware())

	// =====================================================
	// PUBLIC ROUTES - Không yêu cầu xác thực
	// =====================================================

	// Location routes - Quản lý địa điểm (tỉnh/thành, phường/xã)
	publicProducts.GET("/provinces", r.ProductController.PublicGetProvinces, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	publicProducts.GET("/districts", r.ProductController.PublicGetDistricts, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	publicProducts.GET("/wards", r.ProductController.PublicGetWards, middleware.AuthMiddleware(), middleware.PermissionMiddleware())

	// Product routes - Quản lý sản phẩm
	publicProducts.GET("/items/:slug", r.ProductController.PublicGetProductDetail, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	publicProducts.POST("/orders", r.ProductController.PublicPlaceOrder, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	publicProducts.GET("/orders/:id", r.ProductController.PublicGetOrderDetailWithItems, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	publicProducts.GET("/items", r.ProductController.PublicGetProductsList)

	// Wishlist routes - Quản lý danh sách
	publicProducts.POST("/wishlists", r.ProductController.PublicCreateWishlist, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	publicProducts.GET("/wishlists", r.ProductController.PublicGetWishlists, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	publicProducts.GET("/wishlists/:id", r.ProductController.PublicGetWishlistById, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	publicProducts.GET("/wishlists/:customerID/:productID", r.ProductController.PublicGetWishlistByCustomerAndProduct, middleware.AuthMiddleware(), middleware.PermissionMiddleware())
	publicProducts.DELETE("/wishlists/:id", r.ProductController.PublicDeleteWishlist, middleware.AuthMiddleware(), middleware.PermissionMiddleware())

}
