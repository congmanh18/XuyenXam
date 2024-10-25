package usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ArtistService/repository"
)

type DeleteUseCase interface {
	Execute(ctx context.Context, id *string) error
}

type deleteUseCaseImpl struct {
	repo repository.Repo
}

func NewDeleteUseCase(repo repository.Repo) DeleteUseCase {
	return &deleteUseCaseImpl{repo: repo}
}

func (u *deleteUseCaseImpl) Execute(ctx context.Context, id *string) error {
	return u.repo.Delete(ctx, id)
}
