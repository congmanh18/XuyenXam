package inventory

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

// Cập nhật hàng loạt số lượng tồn kho
func (i *inventoryRepo) BulkUpdateInventory(ctx context.Context, inventoryData []entity.Inventory) error {
	// Extract product IDs from map keys
	productIDs := make([]string, 0, len(inventoryData))
	for _, inventory := range inventoryData {
		productIDs = append(productIDs, *inventory.ProductID)
	}
	if err := i.DB.Executor.WithContext(ctx).
		Model(&entity.Inventory{}).
		Where("product_id IN (?)", productIDs).
		Updates(inventoryData).Error; err != nil {
		return err
	}
	return nil
}
