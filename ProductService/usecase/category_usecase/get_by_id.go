package category_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/XuyenXam/ProductService/repository/category"
)

type GetByIDUseCase interface {
	Execute(ctx context.Context, id *string) (*entity.Category, error)
}

type getByIDUseCaseImpl struct {
	CategoryRepo category.CategoryRepository
}

func NewGetByIDUseCase(categoryRepo category.CategoryRepository) GetByIDUseCase {
	return &getByIDUseCaseImpl{CategoryRepo: categoryRepo}
}

func (u *getByIDUseCaseImpl) Execute(ctx context.Context, id *string) (*entity.Category, error) {
	return u.CategoryRepo.GetCategoryByID(ctx, id)
}
