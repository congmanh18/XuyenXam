package repository

import (
	"context"

	"github.com/congmanh18/XuyenXam/OrderService/entity"
	"github.com/congmanh18/lucas-core/db/postgresql"
)

type OrderRepository interface {
	Create(ctx context.Context, order *entity.Order) error
	FindByID(ctx context.Context, id *string) (*entity.Order, error)
	FindAll(ctx context.Context) ([]entity.Order, error)
	Update(ctx context.Context, order *entity.Order) error
	Delete(ctx context.Context, id *string) error
}

type orderRepo struct {
	DB *postgresql.Database
}

func NewOrderRepository(db *postgresql.Database) OrderRepository {
	return &orderRepo{DB: db}
}
