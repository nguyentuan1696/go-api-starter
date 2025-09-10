package controller

import (
	"go-api-starter/core/errors"
	"go-api-starter/core/logger"
	"go-api-starter/core/params"
	"go-api-starter/core/utils"
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/validator"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (controller *ProductController) PrivateGetOrderDetailWithItems(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))

	orderDetail, err := controller.ProductService.PrivateGetOrderDetailWithItems(ctx, id)
	if err != nil {
		logger.Error("ProductController:PrivateGetOrderDetailWithItems error: %v", err)
		return controller.BadRequest(errors.ErrNotFound, "Order not found")
	}

	controller.SuccessResponse(c, orderDetail, "Get order detail with items success")
	return nil
}

func (controller *ProductController) PublicGetOrderDetailWithItems(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))

	orderDetail, err := controller.ProductService.PublicGetOrderDetailWithItems(ctx, id)
	if err != nil {
		logger.Error("ProductController:PublicGetOrderDetailWithItems error: %v", err)
		return controller.BadRequest(errors.ErrNotFound, "Order not found")
	}

	controller.SuccessResponse(c, orderDetail, "Get order detail with items success")
	return nil
}

func (controller *ProductController) PublicPlaceOrder(c echo.Context) error {
	ctx := c.Request().Context()

	requestData := new(dto.PlaceOrderRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidatePlaceOrderRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	orderInfo, err := controller.ProductService.PublicPlaceOrder(ctx, requestData)
	if err != nil {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", nil)
	}

	controller.SuccessResponse(c, orderInfo, "Place order success")
	return nil
}

func (controller *ProductController) PrivateGetOrders(c echo.Context) error {
	ctx := c.Request().Context()

	params := params.NewQueryParams(c)

	orders, err := controller.ProductService.PrivateGetOrders(ctx, *params)
	if err != nil {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", nil)
	}

	controller.SuccessResponse(c, orders, "Get orders success")
	return nil
}

func (controller *ProductController) PrivateGetOrderById(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", nil)
	}

	order, err := controller.ProductService.PrivateGetOrderById(ctx, id)
	if err != nil {
		return controller.BadRequest(errors.ErrInvalidInput, "Order not found", nil)
	}

	controller.SuccessResponse(c, order, "Get order success")
	return nil
}
