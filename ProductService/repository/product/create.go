package product

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (p *productRepo) CreateProduct(ctx context.Context, product *entity.Product) error {
	return p.DB.Executor.WithContext(ctx).Create(product).Error
}
