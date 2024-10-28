package repository

import (
	"context"

	"github.com/congmanh18/XuyenXam/OrderService/entity"
)

func (o *orderRepo) Delete(ctx context.Context, id *string) error {
	return o.DB.Executor.WithContext(ctx).Delete(&entity.Order{}, id).Error
}
