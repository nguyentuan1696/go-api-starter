package service

import (
	"context"
	"sync"
	"go-api-starter/core/errors"
	"go-api-starter/core/params"
	"go-api-starter/modules/product/dto"
	"go-api-starter/modules/product/repository"

	"github.com/google/uuid"
)

type ProductService struct {
	mu   sync.Mutex
	repo repository.ProductRepositoryInterface
}

func NewProductService(repo repository.ProductRepositoryInterface) *ProductService {
	return &ProductService{repo: repo}
}

type ProductServiceInterface interface {

	// Private
	PrivateCreateCategory(ctx context.Context, req *dto.CategoryRequest) *errors.AppError
	PrivateGetCategoryById(ctx context.Context, id uuid.UUID) (*dto.CategoryResponse, *errors.AppError)
	PrivateUpdateCategory(ctx context.Context, req *dto.CategoryRequest, id uuid.UUID) *errors.AppError
	PrivateDeleteCategory(ctx context.Context, id uuid.UUID) *errors.AppError
	PrivateGetCategories(ctx context.Context, params params.QueryParams) (*dto.PaginatedCategoryResponse, *errors.AppError)

	PrivateCreateBrand(ctx context.Context, req *dto.BrandRequest) *errors.AppError
	PrivateGetBrands(ctx context.Context, params params.QueryParams) (*dto.PaginatedBrandResponse, *errors.AppError)
	PrivateGetBrandById(ctx context.Context, id uuid.UUID) (*dto.BrandResponse, *errors.AppError)
	PrivateUpdateBrand(ctx context.Context, req *dto.BrandRequest, id uuid.UUID) *errors.AppError
	PrivateDeleteBrand(ctx context.Context, id uuid.UUID) *errors.AppError

	PrivateCreateIngredient(ctx context.Context, req *dto.IngredientRequest) error
	PrivateGetIngredients(ctx context.Context, params params.QueryParams) (*dto.PaginatedIngredientDTO, *errors.AppError)
	PrivateGetIngredientById(ctx context.Context, id uuid.UUID) (*dto.IngredientResponse, *errors.AppError)
	PrivateUpdateIngredient(ctx context.Context, id uuid.UUID, req *dto.IngredientRequest) *errors.AppError
	PrivateDeleteIngredient(ctx context.Context, id uuid.UUID) *errors.AppError

	PrivateCreateTag(ctx context.Context, req *dto.TagRequest) *errors.AppError
	PrivateGetTags(ctx context.Context, params params.QueryParams) (*dto.PaginatedTagDTO, *errors.AppError)
	PrivateGetTagById(ctx context.Context, id uuid.UUID) (*dto.TagResponse, *errors.AppError)
	PrivateUpdateTag(ctx context.Context, id uuid.UUID, req *dto.TagRequest) *errors.AppError
	PrivateDeleteTag(ctx context.Context, id uuid.UUID) *errors.AppError

	PrivateCreateSkinType(ctx context.Context, req *dto.SkinTypeRequest) *errors.AppError
	PrivateGetSkinTypes(ctx context.Context, params params.QueryParams) (*dto.PaginatedSkinTypeDTO, *errors.AppError)
	PrivateGetSkinTypeById(ctx context.Context, id uuid.UUID) (*dto.SkinTypeResponse, *errors.AppError)
	PrivateUpdateSkinType(ctx context.Context, id uuid.UUID, req *dto.SkinTypeRequest) *errors.AppError
	PrivateDeleteSkinType(ctx context.Context, id uuid.UUID) *errors.AppError

	PrivateCreateBenefit(ctx context.Context, req *dto.BenefitRequest) *errors.AppError
	PrivateGetBenefits(ctx context.Context, params params.QueryParams) (*dto.PaginatedBenefitDTO, *errors.AppError)
	PrivateGetBenefitById(ctx context.Context, id uuid.UUID) (*dto.BenefitResponse, *errors.AppError)
	PrivateUpdateBenefit(ctx context.Context, id uuid.UUID, req *dto.BenefitRequest) *errors.AppError
	PrivateDeleteBenefit(ctx context.Context, id uuid.UUID) *errors.AppError

	PrivateCreateProduct(ctx context.Context, req *dto.ProductRequest) *errors.AppError
	PrivateGetProducts(ctx context.Context, params params.QueryParams) (*dto.PaginatedProductDTO, *errors.AppError)
	PrivateGetProductById(ctx context.Context, id uuid.UUID) (*dto.ProductResponse, *errors.AppError)
	PrivateUpdateProduct(ctx context.Context, id uuid.UUID, req *dto.ProductRequest) *errors.AppError
	PrivateDeleteProduct(ctx context.Context, id uuid.UUID) *errors.AppError

	PrivateCreatePaymentMethod(ctx context.Context, req *dto.PaymentMethodRequest) *errors.AppError
	PrivateGetPaymentMethods(ctx context.Context, params params.QueryParams) (*dto.PaginatedPaymentMethodDTO, *errors.AppError)
	PrivateGetPaymentMethodById(ctx context.Context, id int) (*dto.PaymentMethodResponse, *errors.AppError)
	PrivateUpdatePaymentMethod(ctx context.Context, id int, req *dto.PaymentMethodRequest) *errors.AppError
	PrivateDeletePaymentMethod(ctx context.Context, id int) *errors.AppError

	PrivateCreateShippingMethod(ctx context.Context, req *dto.ShippingMethodRequest) *errors.AppError
	PrivateGetShippingMethods(ctx context.Context, params params.QueryParams) (*dto.PaginatedShippingMethodDTO, *errors.AppError)
	PrivateGetShippingMethodById(ctx context.Context, id int) (*dto.ShippingMethodResponse, *errors.AppError)
	PrivateUpdateShippingMethod(ctx context.Context, id int, req *dto.ShippingMethodRequest) *errors.AppError
	PrivateDeleteShippingMethod(ctx context.Context, id int) *errors.AppError

	PrivateGetOrders(ctx context.Context, params params.QueryParams) (*dto.PaginatedOrderDTO, error)
	PrivateGetOrderById(ctx context.Context, id uuid.UUID) (*dto.OrderResponse, error)
	PrivateGetOrderDetailWithItems(ctx context.Context, orderID uuid.UUID) (*dto.OrderDetailWithItemsDTO, error)

	// Public
	PublicGetProductDetail(ctx context.Context, slug string) (*dto.ProductDetailResponse, *errors.AppError)
	PublicGetProductDetailWithFields(ctx context.Context, id uuid.UUID, field []string) (*dto.ProductDetailResponse, *errors.AppError)
	PublicGetProducts(ctx context.Context, params params.QueryParams) (*dto.PaginatedProductDTO, *errors.AppError)
	
	PublicCreateWishlist(ctx context.Context, req *dto.WishlistRequest) *errors.AppError
	PublicGetWishlists(ctx context.Context, params params.QueryParams) (*dto.PaginatedWishListDTO, *errors.AppError)
	PublicGetWishlistById(ctx context.Context, id string) (*dto.WishlistResponse, *errors.AppError)
	PublicGetWishlistByCustomerAndProduct(ctx context.Context, customerID string, productID string) (*dto.WishlistResponse, *errors.AppError)
	PublicDeleteWishlist(ctx context.Context, id string) *errors.AppError

	PublicPlaceOrder(ctx context.Context, req *dto.PlaceOrderRequest) (*dto.PlaceOrderResponse, *errors.AppError)
	PublicGetOrderDetailWithItems(ctx context.Context, orderID uuid.UUID) (*dto.OrderDetailWithItemsDTO, *errors.AppError)

	PublicGetProvinces(ctx context.Context, params params.QueryParams) (*dto.PaginatedProvinceDTO, *errors.AppError)
	PublicGetDistricts(ctx context.Context, params params.QueryParams) (*dto.PaginatedDistrictDTO, *errors.AppError)
	PublicGetWards(ctx context.Context, params params.QueryParams) (*dto.PaginatedWardDTO, *errors.AppError)
}
