package validator

import (
	"go-api-starter/core/constants"
	"go-api-starter/core/utils"
	"go-api-starter/core/validation"
	"go-api-starter/modules/product/dto"

	"github.com/google/uuid"
)

func ValidatePlaceOrderRequest(req *dto.PlaceOrderRequest) *validation.ValidationResult {
	if req == nil {
		return nil
	}
	result := validation.NewValidationResult()

	if req.CustomerPhone != constants.DefaultEmptyString && !utils.IsValidPhone(req.CustomerPhone) {
		result.AddError("customer_phone", "Invalid customer phone")
	}

	if req.CustomerName == constants.DefaultEmptyString {
		result.AddError("customer_name", "Customer name is required")
	}

	if req.ShippingRecipientName == constants.DefaultEmptyString {
		result.AddError("shipping_recipient_name", "Shipping recipient name is required")
	}

	if req.ShippingRecipientPhone == constants.DefaultEmptyString && !utils.IsValidPhone(req.ShippingRecipientPhone) {
		result.AddError("shipping_recipient_phone", "Shipping recipient phone is required")
	}

	if req.ShippingAddress == constants.DefaultEmptyString {
		result.AddError("shipping_address", "Shipping address is required")
	}

	if req.ShippingMethodID <= constants.DefaultZeroValue {
		result.AddError("shipping_method_id", "Shipping method is required")
	}

	if len(req.OrderItems) <= constants.DefaultZeroValue {
		result.AddError("order_items", "Order items is required")
	}

	for _, item := range req.OrderItems {
		if item.ProductID == uuid.Nil {
			result.AddError("product_id", "Product ID is required")
		}

		if item.Price <= constants.DefaultZeroValue {
			result.AddError("price", "Price must be greater than 0")
		}

		if item.OriginalPrice <= constants.DefaultZeroValue {
			result.AddError("original_price", "Original price must be greater than 0")
		}

		if item.Quantity <= constants.DefaultZeroValue {
			result.AddError("quantity", "Quantity must be greater than 0")
		}
	}

	return result
}

func ValidatePaymentMethodRequest(req *dto.PaymentMethodRequest) *validation.ValidationResult {
	if req == nil {
		return nil
	}

	result := validation.NewValidationResult()

	if utils.IsEmpty(req.Name) {
		result.AddError("name", "Name is required")
	}

	return result
}

func ValidateCategoryRequest(req *dto.CategoryRequest) *validation.ValidationResult {
	if req == nil {
		return nil
	}
	result := validation.NewValidationResult()

	if utils.IsEmpty(req.Name) {
		result.AddError("name", "Name is required")
	}

	return result

}

func ValidateBrandRequest(req *dto.BrandRequest) *validation.ValidationResult {
	if req == nil {
		return nil
	}
	result := validation.NewValidationResult()

	if utils.IsEmpty(req.Name) {
		result.AddError("name", "Name is required")
	}

	return result
}

func ValidateIngredientRequest(req *dto.IngredientRequest) *validation.ValidationResult {
	if req == nil {
		return nil
	}
	result := validation.NewValidationResult()

	if utils.IsEmpty(req.Name) {
		result.AddError("name", "Name is required")
	}

	return result
}

func ValidateTagRequest(req *dto.TagRequest) *validation.ValidationResult {
	if req == nil {
		return nil
	}
	result := validation.NewValidationResult()

	if utils.IsEmpty(req.Name) {
		result.AddError("name", "Name is required")
	}

	return result
}

func ValidateSkinTypeRequest(req *dto.SkinTypeRequest) *validation.ValidationResult {
	if req == nil {
		return nil
	}
	result := validation.NewValidationResult()

	if utils.IsEmpty(req.Name) {
		result.AddError("name", "Name is required")
	}

	return result
}

func ValidateBenefitRequest(req *dto.BenefitRequest) *validation.ValidationResult {
	if req == nil {
		return nil
	}
	result := validation.NewValidationResult()

	if utils.IsEmpty(req.Name) {
		result.AddError("name", "Name is required")
	}

	return result
}

func ValidateProductRequest(req *dto.ProductRequest) *validation.ValidationResult {
	if req == nil {
		return nil
	}
	result := validation.NewValidationResult()

	if utils.IsEmpty(req.Name) {
		result.AddError("name", "Name is required")
	}

	if req.Price < 0 {
		result.AddError("price", "Price must be greater than or equal to 0")
	}

	if req.OriginalPrice < 0 {
		result.AddError("original_price", "Original price must be greater than or equal to 0")
	}

	if req.StockQuantity < 0 {
		result.AddError("stock_quantity", "Stock quantity must be greater than or equal to 0")
	}

	if req.MinStockLevel < 0 {
		result.AddError("min_stock_level", "Min stock level must be greater than or equal to 0")
	}

	return result
}

func ValidateShippingMethodRequest(req *dto.ShippingMethodRequest) *validation.ValidationResult {
	if req == nil {
		return nil
	}
	result := validation.NewValidationResult()

	if utils.IsEmpty(req.Name) {
		result.AddError("name", "Name is required")
	}

	if req.BaseCost < constants.DefaultZeroValue {
		result.AddError("base_cost", "Base cost must be greater than or equal to 0")
	}

	if req.CostPerKg < constants.DefaultZeroValue {
		result.AddError("cost_per_kg", "Cost per kg must be greater than or equal to 0")
	}

	if req.FreeShippingThreshold < constants.DefaultZeroValue {
		result.AddError("free_shipping_threshold", "Free shipping threshold must be greater than or equal to 0")
	}

	if req.EstimatedDaysMin < constants.DefaultZeroValue {
		result.AddError("estimated_days_min", "Estimated days min must be greater than or equal to 0")
	}

	if req.EstimatedDaysMax < constants.DefaultZeroValue {
		result.AddError("estimated_days_max", "Estimated days max must be greater than or equal to 0")
	}

	if req.EstimatedDaysMin > req.EstimatedDaysMax {
		result.AddError("estimated_days", "Estimated days min cannot be greater than estimated days max")
	}

	return result
}

func ValidateWishlistRequest(req *dto.WishlistRequest) *validation.ValidationResult {
	if req == nil {
		return nil
	}
	result := validation.NewValidationResult()

	if req.CustomerID == uuid.Nil {
		result.AddError("user_id", "User ID is required")
	}

	if utils.IsEmpty(req.ProductID) {
		result.AddError("product_id", "Product ID is required")
	}

	return result
}
