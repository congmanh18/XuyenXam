package inventory_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/XuyenXam/ProductService/repository/inventory"
)

type BulkUpdateInventoryUseCase interface {
	Execute(ctx context.Context, inventory []entity.Inventory) error
}

type bulkUpdateInventoryUseCaseImpl struct {
	InventoryRepo inventory.InventoryRepository
}

func NewBulkUpdateInventoryUseCase(inventoryRepo inventory.InventoryRepository) BulkUpdateInventoryUseCase {
	return &bulkUpdateInventoryUseCaseImpl{InventoryRepo: inventoryRepo}
}

func (u *bulkUpdateInventoryUseCaseImpl) Execute(ctx context.Context, inventory []entity.Inventory) error {
	return u.InventoryRepo.BulkUpdateInventory(ctx, inventory)
}
