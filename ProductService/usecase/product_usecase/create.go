package product_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/XuyenXam/ProductService/repository/product"
)

type CreateUseCase interface {
	Execute(ctx context.Context, product *entity.Product) error
}

type creatUseCaseImpl struct {
	ProductRepo product.ProductRepository
}

func NewCreateUseCase(productRepo product.ProductRepository) CreateUseCase {
	return &creatUseCaseImpl{ProductRepo: productRepo}
}

func (u *creatUseCaseImpl) Execute(ctx context.Context, product *entity.Product) error {
	return u.ProductRepo.CreateProduct(ctx, product)
}
