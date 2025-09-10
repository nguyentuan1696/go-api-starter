package controller

import (
	"go-api-starter/core/errors"
	"go-api-starter/core/params"
	"go-api-starter/core/utils"
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/validator"

	"github.com/labstack/echo/v4"
)

func (controller *ProductController) PrivateCreateTag(c echo.Context) error {
	ctx := c.Request().Context()

	requestData := new(dto.TagRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateTagRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	err := controller.ProductService.PrivateCreateTag(ctx, requestData)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "create tag failed", err)
	}

	return controller.SuccessResponse(c, nil, "create tag success")

}

func (controller *ProductController) PrivateGetTags(c echo.Context) error {
	ctx := c.Request().Context()

	params := params.NewQueryParams(c)
	tags, err := controller.ProductService.PrivateGetTags(ctx, *params)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get tags failed", err)
	}

	return controller.SuccessResponse(c, tags, "get tags success")
}

func (controller *ProductController) PrivateGetTagById(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))
	tag, errGet := controller.ProductService.PrivateGetTagById(ctx, id)
	if errGet != nil {
		return controller.NotFound(errors.ErrNotFound, "Tag not found", errGet)
	}

	return controller.SuccessResponse(c, tag, "get tag success")
}

func (controller *ProductController) PrivateUpdateTag(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))

	requestData := new(dto.TagRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateTagRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	errUpdate := controller.ProductService.PrivateUpdateTag(ctx, id, requestData)
	if errUpdate != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "update tag failed", errUpdate)
	}

	return controller.SuccessResponse(c, nil, "update tag success")
}

func (controller *ProductController) PrivateDeleteTag(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))

	errDelete := controller.ProductService.PrivateDeleteTag(ctx, id)
	if errDelete != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "delete tag failed", errDelete)
	}

	return controller.SuccessResponse(c, nil, "delete tag success")
}
