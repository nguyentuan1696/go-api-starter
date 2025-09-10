package controller

import (
	"go-api-starter/core/errors"
	"go-api-starter/core/params"
	"go-api-starter/core/utils"
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/validator"

	"github.com/labstack/echo/v4"
)

func (controller *ProductController) PrivateCreateShippingMethod(c echo.Context) error {
	ctx := c.Request().Context()

	requestData := new(dto.ShippingMethodRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateShippingMethodRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	err := controller.ProductService.PrivateCreateShippingMethod(ctx, requestData)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "create shipping method failed", err)
	}

	return controller.SuccessResponse(c, nil, "create shipping method success")

}

func (controller *ProductController) PrivateGetShippingMethods(c echo.Context) error {
	ctx := c.Request().Context()

	params := params.NewQueryParams(c)
	shippingMethods, err := controller.ProductService.PrivateGetShippingMethods(ctx, *params)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get shipping methods failed", err)
	}

	return controller.SuccessResponse(c, shippingMethods, "get shipping methods success")
}

func (controller *ProductController) PrivateGetShippingMethodById(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToNumber(c.Param("id"))
	shippingMethod, err := controller.ProductService.PrivateGetShippingMethodById(ctx, id)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get shipping method failed", err)
	}

	return controller.SuccessResponse(c, shippingMethod, "get shipping method success")
}

func (controller *ProductController) PrivateUpdateShippingMethod(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	requestData := new(dto.ShippingMethodRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateShippingMethodRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	err := controller.ProductService.PrivateUpdateShippingMethod(ctx, utils.ToNumber(id), requestData)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "update shipping method failed", err)
	}

	return controller.SuccessResponse(c, nil, "update shipping method success")
}

func (controller *ProductController) PrivateDeleteShippingMethod(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToNumber(c.Param("id"))
	err := controller.ProductService.PrivateDeleteShippingMethod(ctx, id)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "delete shipping method failed", err)
	}

	return controller.SuccessResponse(c, nil, "delete shipping method success")
}
