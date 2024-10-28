package inventory

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (i *inventoryRepo) BulkUpdateInventory(ctx context.Context, inventoryData map[string]int) error {
	// Extract product IDs from map keys
	productIDs := make([]string, 0, len(inventoryData))
	for id := range inventoryData {
		productIDs = append(productIDs, id)
	}

	return i.DB.Executor.WithContext(ctx).Model(&entity.Inventory{}).Where("product_id IN (?)", productIDs).Updates(inventoryData).Error
}
