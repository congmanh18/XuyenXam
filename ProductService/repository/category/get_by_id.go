package category

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (p *categoryRepo) GetCategoryByID(ctx context.Context, id *string) (*entity.Category, error) {
	var category entity.Category
	return &category, p.DB.Executor.WithContext(ctx).Where("id = ?", id).First(&category).Error
}
