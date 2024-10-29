package product

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (p *productRepo) SearchProduct(ctx context.Context, query *string) ([]entity.Product, error) {
	var products []entity.Product
	if err := p.DB.Executor.WithContext(ctx).Where("name LIKE ?", "%"+*query+"%").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
