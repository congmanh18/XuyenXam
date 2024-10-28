package category

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (p *categoryRepo) GetAllCategories(ctx context.Context) ([]entity.Category, error) {
	var categories []entity.Category
	return categories, p.DB.Executor.WithContext(ctx).Find(&categories).Error
}
