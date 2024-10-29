package product

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (p *productRepo) CreateProduct(ctx context.Context, product *entity.Product) error {
	err := p.DB.Executor.WithContext(ctx).Create(product).Error
	if err != nil {
		return err
	}
	return nil
}
