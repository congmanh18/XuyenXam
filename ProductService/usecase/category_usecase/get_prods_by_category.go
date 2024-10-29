package category_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/XuyenXam/ProductService/repository/category"
)

type GetProductsByCategoryUseCase interface {
	Execute(ctx context.Context, categoryID *string) ([]entity.Product, error)
}

type getProductsByCategoryUseCaseImpl struct {
	CategoryRepo category.CategoryRepository
}

func NewGetProductsByCategoryUseCase(categoryRepo category.CategoryRepository) GetProductsByCategoryUseCase {
	return &getProductsByCategoryUseCaseImpl{CategoryRepo: categoryRepo}
}

func (u *getProductsByCategoryUseCaseImpl) Execute(ctx context.Context, categoryID *string) ([]entity.Product, error) {
	return u.CategoryRepo.GetProductsByCategory(ctx, categoryID)
}
