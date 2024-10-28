package product

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (p *productRepo) DeleteProduct(ctx context.Context, id *string) error {
	return p.DB.Executor.WithContext(ctx).Delete(&entity.Product{}, id).Error
}
