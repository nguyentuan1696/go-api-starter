package controller

import (
	"go-api-starter/core/errors"
	"go-api-starter/core/params"
	"go-api-starter/core/utils"
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/validator"

	"github.com/labstack/echo/v4"
)

func (controller *ProductController) PrivateCreateSkinType(c echo.Context) error {
	ctx := c.Request().Context()

	requestData := new(dto.SkinTypeRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateSkinTypeRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	err := controller.ProductService.PrivateCreateSkinType(ctx, requestData)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "create skin type failed", err)
	}

	return controller.SuccessResponse(c, nil, "create skin type success")

}

func (controller *ProductController) PrivateGetSkinTypes(c echo.Context) error {
	ctx := c.Request().Context()

	params := params.NewQueryParams(c)
	skinTypes, err := controller.ProductService.PrivateGetSkinTypes(ctx, *params)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get skin types failed", err)
	}

	return controller.SuccessResponse(c, skinTypes, "get skin types success")
}

func (controller *ProductController) PrivateGetSkinTypeById(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))
	skinType, errGet := controller.ProductService.PrivateGetSkinTypeById(ctx, id)
	if errGet != nil {
		return controller.NotFound(errors.ErrNotFound, "Skin type not found", errGet)
	}

	return controller.SuccessResponse(c, skinType, "get skin type success")
}

func (controller *ProductController) PrivateUpdateSkinType(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))

	requestData := new(dto.SkinTypeRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateSkinTypeRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	errUpdate := controller.ProductService.PrivateUpdateSkinType(ctx, id, requestData)
	if errUpdate != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "update skin type failed", errUpdate)
	}

	return controller.SuccessResponse(c, nil, "update skin type success")
}

func (controller *ProductController) PrivateDeleteSkinType(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))

	errDelete := controller.ProductService.PrivateDeleteSkinType(ctx, id)
	if errDelete != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "delete skin type failed", errDelete)
	}

	return controller.SuccessResponse(c, nil, "delete skin type success")
}
