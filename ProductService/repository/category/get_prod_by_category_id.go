package category

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (p *categoryRepo) GetProductsByCategory(ctx context.Context, categoryID *string) ([]entity.Product, error) {
	var products []entity.Product
	err := p.DB.Executor.WithContext(ctx).Where("category_id = ?", categoryID).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
