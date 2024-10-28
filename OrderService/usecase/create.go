package usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/OrderService/entity"
	"github.com/congmanh18/XuyenXam/OrderService/repository"
)

type CreateUseCase interface {
	Execute(ctx context.Context, order *entity.Order) error
}

type creatUseCaseImpl struct {
	OrderRepo repository.OrderRepository
}

func NewCreateUseCase(orderRepo repository.OrderRepository) CreateUseCase {
	return &creatUseCaseImpl{OrderRepo: orderRepo}
}

func (u *creatUseCaseImpl) Execute(ctx context.Context, order *entity.Order) error {
	return u.OrderRepo.Create(ctx, order)
}
