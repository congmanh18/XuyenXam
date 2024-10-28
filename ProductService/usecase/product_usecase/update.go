package product_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	repository "github.com/congmanh18/XuyenXam/ProductService/repository/product"
)

type UpdateUseCase interface {
	Execute(ctx context.Context, product *entity.Product, id *string) error
}

type updateProductImpl struct {
	ProductRepo repository.ProductRepository
}

func NewUpdateUseCase(repo repository.ProductRepository) UpdateUseCase {
	return &updateProductImpl{ProductRepo: repo}
}

func (u *updateProductImpl) Execute(ctx context.Context, product *entity.Product, id *string) error {
	return u.ProductRepo.UpdateProduct(ctx, product, id)
}
