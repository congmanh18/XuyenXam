package product

import (
	"context"
	"reflect"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/lucas-core/gorm"
)

func (p *productRepo) UpdateProduct(ctx context.Context, product *entity.Product, id *string) error {
	omitFields := gorm.OmitFields(product, func(fieldValue reflect.Value) bool {
		return fieldValue.IsZero() ||
			fieldValue.String() == "id" ||
			fieldValue.String() == "created_at"
	})

	if err := p.DB.Executor.WithContext(ctx).
		Model(product).
		Omit(omitFields...).
		Where("id = ?", *id).
		Updates(product).Error; err != nil {
		return err
	}

	return nil
}
