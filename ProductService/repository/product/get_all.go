package product

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

// hạn chế dùng, không có phân trang
func (p *productRepo) GetAllProducts(ctx context.Context) ([]entity.Product, error) {
	var products []entity.Product
	err := p.DB.Executor.WithContext(ctx).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
