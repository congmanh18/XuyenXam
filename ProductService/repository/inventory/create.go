package inventory

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (i *inventoryRepo) CreateInventory(ctx context.Context, productID string, quantity int) error {
	return i.DB.Executor.WithContext(ctx).Create(&entity.Inventory{ProductID: &productID, Quantity: &quantity}).Error
}
