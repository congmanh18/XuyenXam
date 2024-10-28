package usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/OrderService/repository"
)

type DeleteUseCase interface {
	Execute(ctx context.Context, id *string) error
}

type deleteUseCaseImpl struct {
	OrderRepo repository.OrderRepository
}

func NewDeleteUseCase(orderRepo repository.OrderRepository) DeleteUseCase {
	return &deleteUseCaseImpl{OrderRepo: orderRepo}
}

func (u *deleteUseCaseImpl) Execute(ctx context.Context, id *string) error {
	return u.OrderRepo.Delete(ctx, id)
}
