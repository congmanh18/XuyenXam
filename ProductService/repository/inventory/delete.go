package inventory

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (i *inventoryRepo) DeleteInventory(ctx context.Context, productID *string) error {
	err := i.DB.Executor.WithContext(ctx).Delete(&entity.Inventory{}, productID).Error
	if err != nil {
		return err
	}
	return nil
}
