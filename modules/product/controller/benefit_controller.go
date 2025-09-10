package controller

import (
	"go-api-starter/core/errors"
	"go-api-starter/core/params"
	"go-api-starter/core/utils"
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/validator"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (controller *ProductController) PrivateCreateBenefit(c echo.Context) error {
	ctx := c.Request().Context()

	requestData := new(dto.BenefitRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", nil)
	}

	validationResult := validator.ValidateBenefitRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	err := controller.ProductService.PrivateCreateBenefit(ctx, requestData)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "create benefit failed", err)
	}

	return controller.SuccessResponse(c, nil, "create benefit success")

}

func (controller *ProductController) PrivateGetBenefits(c echo.Context) error {
	ctx := c.Request().Context()

	params := params.NewQueryParams(c)
	benefits, err := controller.ProductService.PrivateGetBenefits(ctx, *params)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get benefits failed", err)
	}

	return controller.SuccessResponse(c, benefits, "get benefits success")
}

func (controller *ProductController) PrivateGetBenefitById(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))
	benefit, errGetBenefit := controller.ProductService.PrivateGetBenefitById(ctx, id)
	if errGetBenefit != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "get benefit failed", errGetBenefit)
	}

	return controller.SuccessResponse(c, benefit, "get benefit success")
}

func (controller *ProductController) PrivateUpdateBenefit(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid benefit id", err)
	}
	requestData := new(dto.BenefitRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Invalid request data", err)
	}

	validationResult := validator.ValidateBenefitRequest(requestData)
	if validationResult.HasError() {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid request data", validationResult)
	}

	errUpdateBenefit := controller.ProductService.PrivateUpdateBenefit(ctx, id, requestData)
	if errUpdateBenefit != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "update benefit failed", errUpdateBenefit)
	}

	return controller.SuccessResponse(c, nil, "update benefit success")
}

func (controller *ProductController) PrivateDeleteBenefit(c echo.Context) error {
	ctx := c.Request().Context()

	id := utils.ToUUID(c.Param("id"))
	errDeleteBenefit := controller.ProductService.PrivateDeleteBenefit(ctx, id)
	if errDeleteBenefit != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "delete benefit failed", errDeleteBenefit)
	}

	return controller.SuccessResponse(c, nil, "delete benefit success")
}
