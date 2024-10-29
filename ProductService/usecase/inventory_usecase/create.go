package inventory_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/repository/inventory"
)

type CreateInventoryUseCase interface {
	Execute(ctx context.Context, productID *string, quantity *int) error
}

type createInventoryUseCaseImpl struct {
	InventoryRepo inventory.InventoryRepository
}

func NewCreateInventoryUseCase(inventoryRepo inventory.InventoryRepository) CreateInventoryUseCase {
	return &createInventoryUseCaseImpl{InventoryRepo: inventoryRepo}
}

func (u *createInventoryUseCaseImpl) Execute(ctx context.Context, productID *string, quantity *int) error {
	return u.InventoryRepo.CreateInventory(ctx, productID, quantity)
}
