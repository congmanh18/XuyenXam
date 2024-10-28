package category_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/XuyenXam/ProductService/repository/category"
)

type CreateUseCase interface {
	Execute(ctx context.Context, category *entity.Category) error
}

type createUseCaseImpl struct {
	CategoryRepo category.CategoryRepository
}

func NewCreateUseCase(categoryRepo category.CategoryRepository) CreateUseCase {
	return &createUseCaseImpl{CategoryRepo: categoryRepo}
}

func (u *createUseCaseImpl) Execute(ctx context.Context, category *entity.Category) error {
	return u.CategoryRepo.CreateCategory(ctx, category)
}
