package product_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/XuyenXam/ProductService/repository/product"
	"github.com/congmanh18/lucas-core/record"
)

type GetPaginateUseCase interface {
	Execute(ctx context.Context, pagination *record.Pagination) ([]entity.Product, error)
}

type getPaginateUseCaseImpl struct {
	ProductRepo product.ProductRepository
}

func NewGetPaginateUseCase(productRepo product.ProductRepository) GetPaginateUseCase {
	return &getPaginateUseCaseImpl{ProductRepo: productRepo}
}

func (u *getPaginateUseCaseImpl) Execute(ctx context.Context, pagination *record.Pagination) ([]entity.Product, error) {
	return u.ProductRepo.GetProducts(ctx, pagination)
}
