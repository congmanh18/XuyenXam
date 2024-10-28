package category_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/repository/category"
)

type DeleteUseCase interface {
	Execute(ctx context.Context, id *string) error
}

type deleteUseCaseImpl struct {
	CategoryRepo category.CategoryRepository
}

func NewDeleteUseCase(categoryRepo category.CategoryRepository) DeleteUseCase {
	return &deleteUseCaseImpl{CategoryRepo: categoryRepo}
}

func (u *deleteUseCaseImpl) Execute(ctx context.Context, id *string) error {
	return u.CategoryRepo.DeleteCategory(ctx, id)
}

