package product_category_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/XuyenXam/ProductService/repository/product_category"
)

type GetByCategoryIDUseCase interface {
	Execute(ctx context.Context, categoryID *string) ([]entity.Product, error)
}

type getByCategoryIDUseCaseImpl struct {
	ProductCategoryRepo product_category.ProductCategoryRepository
}

func NewGetByCategoryIDUseCase(productCategoryRepo product_category.ProductCategoryRepository) GetByCategoryIDUseCase {
	return &getByCategoryIDUseCaseImpl{ProductCategoryRepo: productCategoryRepo}
}

func (u *getByCategoryIDUseCaseImpl) Execute(ctx context.Context, categoryID *string) ([]entity.Product, error) {
	products, err := u.ProductCategoryRepo.GetProductsByCategory(ctx, categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}
