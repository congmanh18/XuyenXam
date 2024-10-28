package category_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/XuyenXam/ProductService/repository/category"
)

type GetAllUseCase interface {
	Execute(ctx context.Context) ([]entity.Category, error)
}

type getAllUseCaseImpl struct {
	CategoryRepo category.CategoryRepository
}

func NewGetAllUseCase(categoryRepo category.CategoryRepository) GetAllUseCase {
	return &getAllUseCaseImpl{CategoryRepo: categoryRepo}
}

func (u *getAllUseCaseImpl) Execute(ctx context.Context) ([]entity.Category, error) {
	return u.CategoryRepo.GetAllCategories(ctx)
}
