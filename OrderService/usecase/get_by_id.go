package usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/OrderService/entity"
	"github.com/congmanh18/XuyenXam/OrderService/repository"
)

// type OrderUseCase interface {
// 	GetAllOrders(ctx context.Context) ([]entity.Order, error)
// }

type GetByIDUseCase interface {
	Execute(ctx context.Context, id *string) (*entity.Order, error)
}

type getOrderByIDUseCase struct {
	OrderRepo repository.OrderRepository
}

func NewGetByIDUseCase(orderRepo repository.OrderRepository) GetByIDUseCase {
	return &getOrderByIDUseCase{OrderRepo: orderRepo}
}

func (u *getOrderByIDUseCase) Execute(ctx context.Context, id *string) (*entity.Order, error) {
	return u.OrderRepo.FindByID(ctx, id)
}
