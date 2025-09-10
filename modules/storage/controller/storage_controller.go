package controller

import (
	"go-api-starter/core/errors"
	"go-api-starter/core/utils"

	"github.com/labstack/echo/v4"
)

func (controller *StorageController) UploadToR2(c echo.Context) error {
	ctx := c.Request().Context()

	// Lấy file từ form
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return controller.BadRequest(errors.ErrInvalidRequestData, "Missing image field", nil)
	}

	// Validate file
	if err = utils.ValidateUploadFile(fileHeader); err != nil {
		return controller.BadRequest(errors.ErrInvalidInput, "Invalid file", err.Error())
	}

	// Open the uploaded file
	src, err := fileHeader.Open()
	if err != nil {
		return controller.BadRequest(errors.ErrInvalidInput, "Failed to open uploaded file", nil)
	}

	result, err := controller.service.SaveStorageWithRollback(ctx, src, fileHeader)
	if err != nil {
		return controller.InternalServerError(errors.ErrInternalServer, "Upload file to R2 failed", err)
	}

	// Return success response
	return controller.SuccessResponse(c, result, "Upload file to R2 successfully")
}
