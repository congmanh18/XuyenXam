package product_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/XuyenXam/ProductService/repository/product"
)

type SearchUseCase interface {
	Execute(ctx context.Context, query *string) ([]entity.Product, error)
}

type searchUseCaseImpl struct {
	ProductRepo product.ProductRepository
}

func NewSearchUseCase(productRepo product.ProductRepository) SearchUseCase {
	return &searchUseCaseImpl{ProductRepo: productRepo}
}

func (u *searchUseCaseImpl) Execute(ctx context.Context, query *string) ([]entity.Product, error) {
	return u.ProductRepo.SearchProduct(ctx, query)
}
