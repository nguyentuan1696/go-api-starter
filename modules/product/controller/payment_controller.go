package controller

import (
	"go-api-starter/core/errors"
	"go-api-starter/core/params"
	"go-api-starter/core/utils"
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/validator"

	"github.com/labstack/echo/v4"
)

func (controller *ProductController) PrivateCreatePaymentMethod(c echo.Context) error {
	ctx := c.Request().Context()

	requestData := new(dto.PaymentMethodRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidatePaymentMethodRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	err := controller.ProductService.PrivateCreatePaymentMethod(ctx, requestData)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "create payment method failed", err)
	}

	return controller.SuccessResponse(c, nil, "create payment method success")

}

func (controller *ProductController) PrivateGetPaymentMethods(c echo.Context) error {
	ctx := c.Request().Context()

	params := params.NewQueryParams(c)
	paymentMethods, err := controller.ProductService.PrivateGetPaymentMethods(ctx, *params)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get payment methods failed", err)
	}

	return controller.SuccessResponse(c, paymentMethods, "get payment methods success")
}

func (controller *ProductController) PrivateGetPaymentMethodById(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToNumber(c.Param("id"))
	paymentMethod, err := controller.ProductService.PrivateGetPaymentMethodById(ctx, id)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get payment method failed", err)
	}

	return controller.SuccessResponse(c, paymentMethod, "get payment method success")
}

func (controller *ProductController) PrivateUpdatePaymentMethod(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToNumber(c.Param("id"))
	requestData := new(dto.PaymentMethodRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidatePaymentMethodRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	err := controller.ProductService.PrivateUpdatePaymentMethod(ctx, id, requestData)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "update payment method failed", err)
	}

	return controller.SuccessResponse(c, nil, "update payment method success")
}

func (controller *ProductController) PrivateDeletePaymentMethod(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToNumber(c.Param("id"))
	err := controller.ProductService.PrivateDeletePaymentMethod(ctx, id)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "delete payment method failed", err)
	}

	return controller.SuccessResponse(c, nil, "delete payment method success")
}
