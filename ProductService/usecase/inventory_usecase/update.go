package inventory_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/XuyenXam/ProductService/repository/inventory"
)

type UpdateInventoryUseCase interface {
	Execute(ctx context.Context, entity *entity.Inventory) error
}

type updateInventoryUseCaseImpl struct {
	InventoryRepo inventory.InventoryRepository
}

func NewUpdateInventoryUseCase(inventoryRepo inventory.InventoryRepository) UpdateInventoryUseCase {
	return &updateInventoryUseCaseImpl{InventoryRepo: inventoryRepo}
}

func (u *updateInventoryUseCaseImpl) Execute(ctx context.Context, entity *entity.Inventory) error {
	return u.InventoryRepo.UpdateInventory(ctx, entity)
}
