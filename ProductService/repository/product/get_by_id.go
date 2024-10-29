package product

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (p *productRepo) GetProductByID(ctx context.Context, id *string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.Executor.WithContext(ctx).Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
