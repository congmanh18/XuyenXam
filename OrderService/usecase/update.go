package usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/OrderService/entity"
	"github.com/congmanh18/XuyenXam/OrderService/repository"
)

type UpdateUseCase interface {
	Execute(ctx context.Context, order *entity.Order) error
}

type updateUseCaseImpl struct {
	OrderRepo repository.OrderRepository
}

func NewUpdateUseCase(orderRepo repository.OrderRepository) UpdateUseCase {
	return &updateUseCaseImpl{OrderRepo: orderRepo}
}

func (u *updateUseCaseImpl) Execute(ctx context.Context, order *entity.Order) error {
	return u.OrderRepo.Update(ctx, order)
}
