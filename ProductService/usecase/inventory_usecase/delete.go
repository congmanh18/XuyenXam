package inventory_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/repository/inventory"
)

type DeleteInventoryUseCase interface {
	Execute(ctx context.Context, productID *string) error
}

type deleteInventoryUseCaseImpl struct {
	InventoryRepo inventory.InventoryRepository
}

func NewDeleteInventoryUseCase(inventoryRepo inventory.InventoryRepository) DeleteInventoryUseCase {
	return &deleteInventoryUseCaseImpl{InventoryRepo: inventoryRepo}
}

func (u *deleteInventoryUseCaseImpl) Execute(ctx context.Context, productID *string) error {
	return u.InventoryRepo.DeleteInventory(ctx, productID)
}
