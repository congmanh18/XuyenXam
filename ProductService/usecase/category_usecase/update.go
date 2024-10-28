package category_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/XuyenXam/ProductService/repository/category"
)

type UpdateUseCase interface {
	Execute(ctx context.Context, category *entity.Category, id *string) error
}

type updateUseCaseImpl struct {
	CategoryRepo category.CategoryRepository
}

func NewUpdateUseCase(categoryRepo category.CategoryRepository) UpdateUseCase {
	return &updateUseCaseImpl{CategoryRepo: categoryRepo}
}

func (u *updateUseCaseImpl) Execute(ctx context.Context, category *entity.Category, id *string) error {
	return u.CategoryRepo.UpdateCategory(ctx, category, id)
}
