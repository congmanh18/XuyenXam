package repository

import (
	"context"

	"github.com/congmanh18/XuyenXam/OrderService/entity"
)

func (o *orderRepo) Update(ctx context.Context, order *entity.Order) error {
	return o.DB.Executor.WithContext(ctx).Save(&order).Error
}