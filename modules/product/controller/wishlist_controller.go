package controller

import (
	"go-api-starter/core/errors"
	"go-api-starter/core/params"
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/validator"

	"github.com/labstack/echo/v4"
)

func (controller *ProductController) PublicCreateWishlist(c echo.Context) error {
	ctx := c.Request().Context()

	requestData := new(dto.WishlistRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateWishlistRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	err := controller.ProductService.PublicCreateWishlist(ctx, requestData)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "create wishlist failed", err)
	}

	return controller.SuccessResponse(c, nil, "create wishlist success")
}

func (controller *ProductController) PublicGetWishlists(c echo.Context) error {
	ctx := c.Request().Context()

	params := params.NewQueryParams(c)
	wishlists, err := controller.ProductService.PublicGetWishlists(ctx, *params)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get wishlists failed", err)
	}

	return controller.SuccessResponse(c, wishlists, "get wishlists success")
}

func (controller *ProductController) PublicGetWishlistById(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	wishlist, err := controller.ProductService.PublicGetWishlistById(ctx, id)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get wishlist failed", err)
	}

	return controller.SuccessResponse(c, wishlist, "get wishlist success")
}

func (controller *ProductController) PublicGetWishlistByCustomerAndProduct(c echo.Context) error {
	ctx := c.Request().Context()

	customerID := c.QueryParam("customer_id")
	productID := c.QueryParam("product_id")

	if customerID == "" || productID == "" {
		return controller.BadRequest(errors.ErrInvalidInput, "customer_id and product_id are required", nil)
	}

	wishlist, err := controller.ProductService.PublicGetWishlistByCustomerAndProduct(ctx, customerID, productID)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get wishlist failed", err)
	}

	return controller.SuccessResponse(c, wishlist, "get wishlist success")
}

func (controller *ProductController) PublicDeleteWishlist(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	err := controller.ProductService.PublicDeleteWishlist(ctx, id)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "delete wishlist failed", err)
	}

	return controller.SuccessResponse(c, nil, "delete wishlist success")
}
