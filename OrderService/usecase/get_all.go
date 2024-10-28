package usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/OrderService/entity"
	"github.com/congmanh18/XuyenXam/OrderService/repository"
)

type GetAllUseCase interface {
	Execute(ctx context.Context) ([]entity.Order, error)
}

type getAllUseCase struct {
	OrderRepo repository.OrderRepository
}

func NewGetAllUseCase(orderRepo repository.OrderRepository) GetAllUseCase {
	return &getAllUseCase{OrderRepo: orderRepo}
}

func (u *getAllUseCase) Execute(ctx context.Context) ([]entity.Order, error) {
	return u.OrderRepo.FindAll(ctx)
}
