package product_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	repository "github.com/congmanh18/XuyenXam/ProductService/repository/product"
)

type GetByIDUseCase interface {
	Execute(ctx context.Context, id *string) (*entity.Product, error)
}

type getByIDUseCaseImpl struct {
	ProductRepository repository.ProductRepository
}

func NewGetByIDUseCase(productRepository repository.ProductRepository) GetByIDUseCase {
	return &getByIDUseCaseImpl{ProductRepository: productRepository}
}

func (u *getByIDUseCaseImpl) Execute(ctx context.Context, id *string) (*entity.Product, error) {
	return u.ProductRepository.GetProductByID(ctx, id)
}
