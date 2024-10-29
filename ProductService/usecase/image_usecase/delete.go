package image_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/repository/image"
)

type DeleteImageUseCase interface {
	Execute(ctx context.Context, id *string) error
}

type deleteImageUseCaseImpl struct {
	ImageRepo image.ImageRepository
}

func NewDeleteImageUseCase(imageRepo image.ImageRepository) DeleteImageUseCase {
	return &deleteImageUseCaseImpl{ImageRepo: imageRepo}
}

func (u *deleteImageUseCaseImpl) Execute(ctx context.Context, id *string) error {
	return u.ImageRepo.DeleteImage(ctx, id)
}
