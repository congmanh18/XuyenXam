package category

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (p *categoryRepo) DeleteCategory(ctx context.Context, id *string) error {
	return p.DB.Executor.WithContext(ctx).Delete(&entity.Category{}, id).Error
}
