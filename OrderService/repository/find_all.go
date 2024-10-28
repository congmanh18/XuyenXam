package repository

import (
	"context"

	"github.com/congmanh18/XuyenXam/OrderService/entity"
)

func (o *orderRepo) FindAll(ctx context.Context) ([]entity.Order, error) {
	var orders []entity.Order
	return orders, o.DB.Executor.WithContext(ctx).Find(&orders).Error
}
