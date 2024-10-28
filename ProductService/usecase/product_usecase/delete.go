package product_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/repository/product"
)

type DeleteUseCase interface {
	Execute(ctx context.Context, id *string) error
}

type deleteUseCaseImpl struct {
	ProductRepo product.ProductRepository
}

func NewDeleteUseCase(producRepo product.ProductRepository) DeleteUseCase {
	return &deleteUseCaseImpl{ProductRepo: producRepo}
}

func (u *deleteUseCaseImpl) Execute(ctx context.Context, id *string) error {
	return u.ProductRepo.DeleteProduct(ctx, id)
}
