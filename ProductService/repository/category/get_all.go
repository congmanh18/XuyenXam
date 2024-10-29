package category

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (p *categoryRepo) GetAllCategories(ctx context.Context) ([]entity.Category, error) {
	var categories []entity.Category
	err := p.DB.Executor.WithContext(ctx).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}
