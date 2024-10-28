package inventory

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (i *inventoryRepo) UpdateInventory(ctx context.Context, productID string, quantity int) error {
	return i.DB.Executor.WithContext(ctx).Model(&entity.Inventory{}).Where("product_id = ?", productID).Update("quantity", quantity).Error
}
