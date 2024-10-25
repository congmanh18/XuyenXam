package usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ArtistService/entity"
	"github.com/congmanh18/XuyenXam/ArtistService/repository"
)

type CreateUseCase interface {
	Execute(ctx context.Context, artist *entity.Artist) error
}

type createUseCase struct {
	repo repository.Repo
}

func NewCreateUseCase(repo repository.Repo) CreateUseCase {
	return &createUseCase{repo: repo}
}

func (uc *createUseCase) Execute(ctx context.Context, artist *entity.Artist) error {
	return uc.repo.Create(ctx, artist)
}
