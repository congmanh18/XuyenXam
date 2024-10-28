package repository

import (
	"context"

	"github.com/congmanh18/XuyenXam/OrderService/entity"
)

func (o *orderRepo) FindByID(ctx context.Context, id *string) (*entity.Order, error) {
	var order entity.Order
	return &order, o.DB.Executor.WithContext(ctx).Where("id = ?", id).First(&order).Error
}
