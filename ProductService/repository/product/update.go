package product

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (p *productRepo) UpdateProduct(ctx context.Context, product *entity.Product, id *string) error {
	return p.DB.Executor.WithContext(ctx).Where("id = ?", id).Save(product).Error
}
