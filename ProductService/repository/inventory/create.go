package inventory

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (i *inventoryRepo) CreateInventory(ctx context.Context, productID *string, quantity *int) error {
	err := i.DB.Executor.WithContext(ctx).Create(&entity.Inventory{ProductID: productID, Quantity: quantity}).Error
	if err != nil {
		return err
	}
	return nil
}
