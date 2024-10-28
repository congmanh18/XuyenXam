package product

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (p *productRepo) GetAllProducts(ctx context.Context) ([]entity.Product, error) {
	var products []entity.Product
	return products, p.DB.Executor.WithContext(ctx).Find(&products).Error
}
