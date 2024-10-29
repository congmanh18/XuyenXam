package inventory

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (i *inventoryRepo) GetInventoryByProductID(ctx context.Context, productID *string) (int, error) {
	var inventory entity.Inventory
	if err := i.DB.Executor.WithContext(ctx).Where("product_id = ?", productID).First(&inventory).Error; err != nil {
		return 0, err
	}
	return *inventory.Quantity, nil
}
