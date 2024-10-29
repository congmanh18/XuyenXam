package inventory

import (
	"context"
	"reflect"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/lucas-core/gorm"
)

func (i *inventoryRepo) UpdateInventory(ctx context.Context, entity *entity.Inventory) error {
	omitFields := gorm.OmitFields(entity, func(fieldValue reflect.Value) bool {
		return fieldValue.IsZero() ||
			fieldValue.String() == "id" ||
			fieldValue.String() == "created_at" ||
			fieldValue.String() == "product_id"
	})
	err := i.DB.Executor.WithContext(ctx).
		Model(entity).
		Omit(omitFields...).
		Where("product_id = ?", entity.ProductID).
		Update("quantity", entity.Quantity).Error
	if err != nil {
		return err
	}
	return nil
}
