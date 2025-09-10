package controller

import (
	"go-api-starter/core/errors"
	"go-api-starter/core/params"

	"github.com/labstack/echo/v4"
)

func (controller *ProductController) PublicGetProvinces(c echo.Context) error {
	ctx := c.Request().Context()

	params := params.NewQueryParams(c)
	result, err := controller.ProductService.PublicGetProvinces(ctx, *params)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get provinces failed", err)
	}
	return controller.SuccessResponse(c, result, "get provinces success")
}

func (controller *ProductController) PublicGetDistricts(c echo.Context) error {
	ctx := c.Request().Context()

	params := params.NewQueryParams(c)
	result, err := controller.ProductService.PublicGetDistricts(ctx, *params)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get districts failed", err)
	}
	return controller.SuccessResponse(c, result, "get districts success")
}

func (controller *ProductController) PublicGetWards(c echo.Context) error {
	ctx := c.Request().Context()

	params := params.NewQueryParams(c)
	result, err := controller.ProductService.PublicGetWards(ctx, *params)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get wards failed", err)
	}
	return controller.SuccessResponse(c, result, "get wards success")
}
