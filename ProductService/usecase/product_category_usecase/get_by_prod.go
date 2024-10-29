package product_category_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/XuyenXam/ProductService/repository/product_category"
)

type GetByProductIDUseCase interface {
	Execute(ctx context.Context, productID *string) ([]entity.Category, error)
}

type getByProductIDUseCaseImpl struct {
	ProductCategoryRepo product_category.ProductCategoryRepository
}

func NewGetByProductIDUseCase(productCategoryRepo product_category.ProductCategoryRepository) GetByProductIDUseCase {
	return &getByProductIDUseCaseImpl{ProductCategoryRepo: productCategoryRepo}
}

func (u *getByProductIDUseCaseImpl) Execute(ctx context.Context, productID *string) ([]entity.Category, error) {
	categories, err := u.ProductCategoryRepo.GetCategoriesByProduct(ctx, productID)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
