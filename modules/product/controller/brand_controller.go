package controller

import (
	"go-api-starter/core/errors"
	"go-api-starter/core/params"
	"go-api-starter/core/utils"
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/validator"

	"github.com/labstack/echo/v4"
)

func (controller *ProductController) PrivateCreateBrand(c echo.Context) error {
	ctx := c.Request().Context()

	requestData := new(dto.BrandRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateBrandRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	err := controller.ProductService.PrivateCreateBrand(ctx, requestData)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "create brand failed", err)
	}

	return controller.SuccessResponse(c, nil, "create brand success")

}

func (controller *ProductController) PrivateGetBrands(c echo.Context) error {
	ctx := c.Request().Context()

	params := params.NewQueryParams(c)
	brands, err := controller.ProductService.PrivateGetBrands(ctx, *params)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get brands failed", err)
	}

	return controller.SuccessResponse(c, brands, "get brands success")
}

func (controller *ProductController) PrivateGetBrandById(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))
	brand, errGet := controller.ProductService.PrivateGetBrandById(ctx, id)
	if errGet != nil {
		return controller.NotFound(errors.ErrNotFound, "Brand not found", errGet)
	}

	return controller.SuccessResponse(c, brand, "get brand success")
}

func (controller *ProductController) PrivateUpdateBrand(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))
	requestData := new(dto.BrandRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateBrandRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	errUpdate := controller.ProductService.PrivateUpdateBrand(ctx, requestData, id)
	if errUpdate != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "update brand failed", errUpdate)
	}

	return controller.SuccessResponse(c, nil, "update brand success")
}

func (controller *ProductController) PrivateDeleteBrand(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))
	errDelete := controller.ProductService.PrivateDeleteBrand(ctx, id)
	if errDelete != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "delete brand failed", errDelete)
	}

	return controller.SuccessResponse(c, nil, "delete brand success")
}
