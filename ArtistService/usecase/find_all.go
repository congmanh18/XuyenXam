package usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ArtistService/entity"
	"github.com/congmanh18/XuyenXam/ArtistService/repository"
)

type FindAllUseCase interface {
	Execute(ctx context.Context) ([]entity.Artist, error)
}

type findAllUseCase struct {
	repo repository.Repo
}

func NewFindAllUseCase(repo repository.Repo) FindAllUseCase {
	return &findAllUseCase{repo: repo}
}
func (uc *findAllUseCase) Execute(ctx context.Context) ([]entity.Artist, error) {
	return uc.repo.FindAll(ctx)
}
