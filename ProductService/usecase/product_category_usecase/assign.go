package product_category_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/repository/product_category"
)

type AssignUseCase interface {
	Execute(ctx context.Context, productID *string, categoryID *string) error
}

type assignUseCaseImpl struct {
	ProductCategoryRepo product_category.ProductCategoryRepository
}

func NewAssignUseCase(productCategoryRepo product_category.ProductCategoryRepository) AssignUseCase {
	return &assignUseCaseImpl{ProductCategoryRepo: productCategoryRepo}
}

func (u *assignUseCaseImpl) Execute(ctx context.Context, productID *string, categoryID *string) error {
	return u.ProductCategoryRepo.AssignProductToCategory(ctx, productID, categoryID)
}
