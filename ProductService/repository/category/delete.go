package category

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (p *categoryRepo) DeleteCategory(ctx context.Context, id *string) error {
	err := p.DB.Executor.WithContext(ctx).Delete(&entity.Category{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
