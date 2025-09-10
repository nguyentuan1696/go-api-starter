package controller

import (
	"go-api-starter/core/errors"
	"go-api-starter/core/params"
	"go-api-starter/core/utils"
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/validator"

	"github.com/labstack/echo/v4"
)

func (controller *ProductController) PrivateCreateIngredient(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(dto.IngredientRequest)
	if err := c.Bind(req); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateIngredientRequest(req)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	err := controller.ProductService.PrivateCreateIngredient(ctx, req)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "create resource failed", err)
	}

	return controller.SuccessResponse(c, nil, "Ingredient created successfully")
}

func (controller *ProductController) PrivateGetIngredients(c echo.Context) error {
	ctx := c.Request().Context()

	params := params.NewQueryParams(c)

	ingredients, err := controller.ProductService.PrivateGetIngredients(ctx, *params)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get ingredients failed", err)
	}

	return controller.SuccessResponse(c, ingredients, "get ingredients success")
}

func (controller *ProductController) PrivateGetIngredientById(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))

	ingredient, errGet := controller.ProductService.PrivateGetIngredientById(ctx, id)
	if errGet != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get ingredient failed", errGet)
	}

	return controller.SuccessResponse(c, ingredient, "get ingredient success")
}

func (controller *ProductController) PrivateUpdateIngredient(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))

	req := new(dto.IngredientRequest)
	if err := c.Bind(req); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateIngredientRequest(req)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	errUpdate := controller.ProductService.PrivateUpdateIngredient(ctx, id, req)
	if errUpdate != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "update ingredient failed", errUpdate)
	}

	return controller.SuccessResponse(c, nil, "update ingredient success")
}

func (controller *ProductController) PrivateDeleteIngredient(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))

	errDelete := controller.ProductService.PrivateDeleteIngredient(ctx, id)
	if errDelete != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "delete ingredient failed", errDelete)
	}

	return controller.SuccessResponse(c, nil, "delete ingredient success")
}
