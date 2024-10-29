package category

import (
	"context"
	"reflect"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/lucas-core/gorm"
)

func (p *categoryRepo) UpdateCategory(ctx context.Context, category *entity.Category, id *string) error {
	category.ID = id
	omitFields := gorm.OmitFields(category, func(fieldValue reflect.Value) bool {
		return fieldValue.IsZero() ||
			fieldValue.String() == "id" ||
			fieldValue.String() == "created_at"
	})
	if err := p.DB.Executor.WithContext(ctx).
		Model(category).
		Omit(omitFields...).
		Where("id = ?", *id).
		Updates(category).Error; err != nil {
		return err
	}

	return nil
}
