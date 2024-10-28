package category

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (p *categoryRepo) UpdateCategory(ctx context.Context, category *entity.Category, id *string) error {
	return p.DB.Executor.WithContext(ctx).Where("id = ?", id).Save(category).Error
}
