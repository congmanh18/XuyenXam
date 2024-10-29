package inventory_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/XuyenXam/ProductService/repository/inventory"
)

type GetByProductIDUseCase interface {
	Execute(ctx context.Context, productID *string) (*entity.Inventory, error)
}

type getByProductIDUseCaseImpl struct {
	InventoryRepo inventory.InventoryRepository
}

func NewGetByProductIDUseCase(inventoryRepo inventory.InventoryRepository) GetByProductIDUseCase {
	return &getByProductIDUseCaseImpl{InventoryRepo: inventoryRepo}
}

func (u *getByProductIDUseCaseImpl) Execute(ctx context.Context, productID *string) (*entity.Inventory, error) {
	inventory, err := u.InventoryRepo.GetInventoryByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}
	return &entity.Inventory{ProductID: productID, Quantity: &inventory}, nil
}
