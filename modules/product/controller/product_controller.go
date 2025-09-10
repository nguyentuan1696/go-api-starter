package controller

import (
	"go-api-starter/core/errors"
	"go-api-starter/core/params"
	"go-api-starter/core/utils"
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/validator"

	"github.com/labstack/echo/v4"
)

func (controller *ProductController) PublicGetProductsList(c echo.Context) error {
	ctx := c.Request().Context()

	params := params.NewQueryParams(c)
	products, err := controller.ProductService.PublicGetProducts(ctx, *params)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get products failed", err)
	}

	return controller.SuccessResponse(c, products, "get products success")
}

func (controller *ProductController) PublicGetProductDetail(c echo.Context) error {
	ctx := c.Request().Context()

	slug := c.Param("slug")
	product, errGet := controller.ProductService.PublicGetProductDetail(ctx, slug)
	if errGet != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get product failed", errGet)
	}

	return controller.SuccessResponse(c, product, "get product success")
}

func (controller *ProductController) PrivateCreateProduct(c echo.Context) error {
	ctx := c.Request().Context()

	requestData := new(dto.ProductRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateProductRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	err := controller.ProductService.PrivateCreateProduct(ctx, requestData)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "create product failed", err)
	}

	return controller.SuccessResponse(c, nil, "create product success")
}

func (controller *ProductController) PrivateGetProducts(c echo.Context) error {
	ctx := c.Request().Context()

	params := params.NewQueryParams(c)
	products, err := controller.ProductService.PrivateGetProducts(ctx, *params)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get products failed", err)
	}

	return controller.SuccessResponse(c, products, "get products success")
}

func (controller *ProductController) PrivateGetProductById(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))

	product, errGet := controller.ProductService.PrivateGetProductById(ctx, id)
	if errGet != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get product failed", errGet)
	}

	return controller.SuccessResponse(c, product, "get product success")
}

func (controller *ProductController) PrivateUpdateProduct(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))
	requestData := new(dto.ProductRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateProductRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	errUpdate := controller.ProductService.PrivateUpdateProduct(ctx, id, requestData)
	if errUpdate != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "update product failed", errUpdate)
	}

	return controller.SuccessResponse(c, nil, "update product success")
}

func (controller *ProductController) PrivateDeleteProduct(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))
	errDelete := controller.ProductService.PrivateDeleteProduct(ctx, id)
	if errDelete != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "delete product failed", errDelete)
	}

	return controller.SuccessResponse(c, nil, "delete product success")
}
