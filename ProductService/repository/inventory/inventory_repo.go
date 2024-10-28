package inventory

import (
	"context"

	"github.com/congmanh18/lucas-core/db/postgresql"
)

type InventoryRepository interface {
	CreateInventory(ctx context.Context, productID string, quantity int) error
	GetInventoryByProductID(ctx context.Context, productID string) (int, error)
	UpdateInventory(ctx context.Context, productID string, quantity int) error
	DeleteInventory(ctx context.Context, productID string) error
	BulkUpdateInventory(ctx context.Context, inventoryData map[string]int) error
}

type inventoryRepo struct {
	DB *postgresql.Database
}

func NewInventoryRepository(db *postgresql.Database) InventoryRepository {
	return &inventoryRepo{DB: db}
}
