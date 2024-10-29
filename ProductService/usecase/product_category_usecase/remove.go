package product_category_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/repository/product_category"
)

type RemoveUseCase interface {
	Execute(ctx context.Context, productID *string, categoryID *string) error
}

type removeUseCaseImpl struct {
	ProductCategoryRepo product_category.ProductCategoryRepository
}

func NewRemoveUseCase(productCategoryRepo product_category.ProductCategoryRepository) RemoveUseCase {
	return &removeUseCaseImpl{ProductCategoryRepo: productCategoryRepo}
}

func (u *removeUseCaseImpl) Execute(ctx context.Context, productID *string, categoryID *string) error {
	return u.ProductCategoryRepo.RemoveProductFromCategory(ctx, productID, categoryID)
}
