package product

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (p *productRepo) GetProductByID(ctx context.Context, id *string) (*entity.Product, error) {
	var product entity.Product
	return &product, p.DB.Executor.WithContext(ctx).Where("id = ?", id).First(&product).Error
}
