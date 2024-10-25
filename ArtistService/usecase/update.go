package usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ArtistService/entity"
	"github.com/congmanh18/XuyenXam/ArtistService/repository"
)

type UpdateUseCase interface {
	Execute(ctx context.Context, artist *entity.Artist) error
}

type updateUseCase struct {
	repo repository.Repo
}

func NewUpdateUseCase(repo repository.Repo) UpdateUseCase {
	return &updateUseCase{repo: repo}
}

func (uc *updateUseCase) Execute(ctx context.Context, artist *entity.Artist) error {
	return uc.repo.Update(ctx, artist)
}
