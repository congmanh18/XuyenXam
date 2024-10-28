package category

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (p *categoryRepo) CreateCategory(ctx context.Context, category *entity.Category) error {
	return p.DB.Executor.WithContext(ctx).Create(category).Error
}
