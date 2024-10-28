package product_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/XuyenXam/ProductService/repository/product"
)

type GetAllUseCase interface {
	Execute(ctx context.Context) ([]entity.Product, error)
}

type getAllUseCaseImpl struct {
	ProductRepo product.ProductRepository
}

func NewGetAllUseCase(productRepo product.ProductRepository) GetAllUseCase {
	return &getAllUseCaseImpl{ProductRepo: productRepo}
}

func (u *getAllUseCaseImpl) Execute(ctx context.Context) ([]entity.Product, error) {
	return u.ProductRepo.GetAllProducts(ctx)
}
