package controller

import (
	"go-api-starter/core/errors"
	"go-api-starter/core/params"
	"go-api-starter/core/utils"
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/validator"

	"github.com/labstack/echo/v4"
)

func (controller *ProductController) PrivateCreateCategory(c echo.Context) error {
	ctx := c.Request().Context()

	requestData := new(dto.CategoryRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateCategoryRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	err := controller.ProductService.PrivateCreateCategory(ctx, requestData)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "create category failed", err)
	}

	return controller.SuccessResponse(c, nil, "create category success")

}

func (controller *ProductController) PrivateGetCategoryById(c echo.Context) error {
	ctx := c.Request().Context()

	categoryId := utils.ToUUID(c.Param("id"))

	category, errGet := controller.ProductService.PrivateGetCategoryById(ctx, categoryId)
	if errGet != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get category failed", errGet)
	}

	return controller.SuccessResponse(c, category, "get category success")

}

func (controller *ProductController) PrivateUpdateCategory(c echo.Context) error {
	ctx := c.Request().Context()

	categoryId := utils.ToUUID(c.Param("id"))

	requestData := new(dto.CategoryRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateCategoryRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	errUpdate := controller.ProductService.PrivateUpdateCategory(ctx, requestData, categoryId)
	if errUpdate != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "update category failed", errUpdate)
	}

	return controller.SuccessResponse(c, nil, "update category success")
}

func (controller *ProductController) PrivateDeleteCategory(c echo.Context) error {
	ctx := c.Request().Context()

	categoryId := utils.ToUUID(c.Param("id"))
	errDelete := controller.ProductService.PrivateDeleteCategory(ctx, categoryId)
	if errDelete != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "delete category failed", errDelete)
	}

	return controller.SuccessResponse(c, nil, "delete category success")
}

func (controller *ProductController) PrivateGetCategories(c echo.Context) error {
	ctx := c.Request().Context()

	queryParams := params.NewQueryParams(c)

	categories, err := controller.ProductService.PrivateGetCategories(ctx, *queryParams)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get categories failed", err)
	}

	return controller.SuccessResponse(c, categories, "get categories success")
}
