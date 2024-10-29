package product

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (p *productRepo) DeleteProduct(ctx context.Context, id *string) error {
	err := p.DB.Executor.WithContext(ctx).Delete(&entity.Product{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
