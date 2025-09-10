package repository

import (
	"context"
	"go-api-starter/core/database"
	"go-api-starter/core/params"
	"go-api-starter/modules/product/entity"

	"github.com/google/uuid"
)

type ProductRepository struct {
	DB database.Database
}

func NewProductRepository(db database.Database) *ProductRepository {
	return &ProductRepository{DB: db}
}

type ProductRepositoryInterface interface {

	// Private
	PrivateCreateCategory(ctx context.Context, category *entity.Category) error
	PrivateGetCategoryById(ctx context.Context, id uuid.UUID) (*entity.Category, error)
	PrivateGetCategories(ctx context.Context, params params.QueryParams) (*entity.PaginatedCategoryResponse, error)
	PrivateUpdateCategory(ctx context.Context, category *entity.Category, id uuid.UUID) error
	PrivateDeleteCategory(ctx context.Context, id uuid.UUID) error

	PrivateCreateBrand(ctx context.Context, brand *entity.Brand) error
	PrivateGetBrandById(ctx context.Context, id uuid.UUID) (*entity.Brand, error)
	PrivateGetBrands(ctx context.Context, params params.QueryParams) (*entity.PaginatedBrandResponse, error)
	PrivateUpdateBrand(ctx context.Context, brand *entity.Brand, id uuid.UUID) error
	PrivateDeleteBrand(ctx context.Context, id uuid.UUID) error

	PrivateCreateIngredient(ctx context.Context, ingredient *entity.Ingredient) error
	PrivateGetIngredientById(ctx context.Context, id uuid.UUID) (*entity.Ingredient, error)
	PrivateGetIngredients(ctx context.Context, params params.QueryParams) (*entity.PaginatedIngredientEntity, error)
	PrivateUpdateIngredient(ctx context.Context, ingredient *entity.Ingredient, id uuid.UUID) error
	PrivateDeleteIngredient(ctx context.Context, id uuid.UUID) error

	PrivateCreateTag(ctx context.Context, tag *entity.Tag) error
	PrivateGetTagById(ctx context.Context, id uuid.UUID) (*entity.Tag, error)
	PrivateGetTags(ctx context.Context, params params.QueryParams) (*entity.PaginatedTagEntity, error)
	PrivateUpdateTag(ctx context.Context, tag *entity.Tag, id uuid.UUID) error
	PrivateDeleteTag(ctx context.Context, id uuid.UUID) error

	PrivateCreateSkinType(ctx context.Context, skinType *entity.SkinType) error
	PrivateGetSkinTypeById(ctx context.Context, id uuid.UUID) (*entity.SkinType, error)
	PrivateGetSkinTypes(ctx context.Context, params params.QueryParams) (*entity.PaginatedSkinTypeEntity, error)
	PrivateUpdateSkinType(ctx context.Context, skinType *entity.SkinType, id uuid.UUID) error
	PrivateDeleteSkinType(ctx context.Context, id uuid.UUID) error

	PrivateCreateBenefit(ctx context.Context, benefit *entity.Benefit) error
	PrivateGetBenefitById(ctx context.Context, id uuid.UUID) (*entity.Benefit, error)
	PrivateGetBenefits(ctx context.Context, params params.QueryParams) (*entity.PaginatedBenefitEntity, error)
	PrivateUpdateBenefit(ctx context.Context, benefit *entity.Benefit, id uuid.UUID) error
	PrivateDeleteBenefit(ctx context.Context, id uuid.UUID) error

	PrivateCreateProduct(ctx context.Context, product *entity.Product) error
	PrivateGetProducts(ctx context.Context, params params.QueryParams) (*entity.PaginatedProductEntity, error)
	PrivateGetProductById(ctx context.Context, id uuid.UUID) (*entity.Product, error)
	PrivateUpdateProduct(ctx context.Context, product *entity.Product, id uuid.UUID) error
	PrivateDeleteProduct(ctx context.Context, id uuid.UUID) error

	PrivateCreatePaymentMethod(ctx context.Context, paymentMethod *entity.PaymentMethod) error
	PrivateGetPaymentMethods(ctx context.Context, params params.QueryParams) (*entity.PaginatedPaymentMethodEntity, error)
	PrivateGetPaymentMethodById(ctx context.Context, id int) (*entity.PaymentMethod, error)
	PrivateUpdatePaymentMethod(ctx context.Context, paymentMethod *entity.PaymentMethod, id int) error
	PrivateDeletePaymentMethod(ctx context.Context, id int) error

	PrivateCreateShippingMethod(ctx context.Context, shippingMethod *entity.ShippingMethod) error
	PrivateGetShippingMethods(ctx context.Context, params params.QueryParams) (*entity.PaginatedShippingMethodEntity, error)
	PrivateGetShippingMethodById(ctx context.Context, id int) (*entity.ShippingMethod, error)
	PrivateUpdateShippingMethod(ctx context.Context, shippingMethod *entity.ShippingMethod, id int) error
	PrivateDeleteShippingMethod(ctx context.Context, id int) error

	PrivateGetOrderById(ctx context.Context, id uuid.UUID) (*entity.Order, error)
	PrivateGetOrders(ctx context.Context, params params.QueryParams) (*entity.PaginatedOrderEntity, error)
	PrivateCreateOrderItems(ctx context.Context, orderItems []*entity.OrderItem) error

	// Public

	PublicGetProductDetail(ctx context.Context, id uuid.UUID, fields []string) (*entity.ProductDetailEntity, error)
	PublicGetProductDetailBySlug(ctx context.Context, slug string, fields []string) (*entity.ProductDetailEntity, error)

	PublicCreateOrder(ctx context.Context, order *entity.Order) (*entity.Order, error)
	GetOrderDetailWithItems(ctx context.Context, orderID uuid.UUID) (*entity.OrderDetailWithItems, error)

	PublicCreateWishlist(ctx context.Context, wishlist *entity.Wishlist) error
	PublicGetWishlists(ctx context.Context, params params.QueryParams) (*entity.PaginatedWishlistEntity, error)
	PublicGetWishlistById(ctx context.Context, id string) (*entity.Wishlist, error)
	PublicGetWishlistByCustomerAndProduct(ctx context.Context, customerID string, productID string) (*entity.Wishlist, error)
	PublicDeleteWishlist(ctx context.Context, id string) error

	PublicGetProvinces(ctx context.Context, params params.QueryParams) (*entity.PaginatedProvinceEntity, error)
	PublicGetDistricts(ctx context.Context, params params.QueryParams) (*entity.PaginatedDistrictEntity, error)
	PublicGetWards(ctx context.Context, params params.QueryParams) (*entity.PaginatedWardEntity, error)
}
