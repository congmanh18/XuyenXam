package usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ArtistService/entity"
	"github.com/congmanh18/XuyenXam/ArtistService/repository"
)

type FindByIDUseCase interface {
	Execute(ctx context.Context, id *string) (*entity.Artist, error)
}

type findByIDUseCase struct {
	repo repository.Repo
}

func NewFindByIDUseCase(repo repository.Repo) FindByIDUseCase {
	return &findByIDUseCase{repo: repo}
}

func (uc *findByIDUseCase) Execute(ctx context.Context, id *string) (*entity.Artist, error) {
	return uc.repo.FindByID(ctx, id)
}
