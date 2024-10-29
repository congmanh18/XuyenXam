package image_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/XuyenXam/ProductService/repository/image"
)

type AddImageUseCase interface {
	Execute(ctx context.Context, image *entity.Image) error
}

type addImageUseCaseImpl struct {
	ImageRepo image.ImageRepository
}

func NewAddImageUseCase(imageRepo image.ImageRepository) AddImageUseCase {
	return &addImageUseCaseImpl{ImageRepo: imageRepo}
}

func (u *addImageUseCaseImpl) Execute(ctx context.Context, image *entity.Image) error {
	return u.ImageRepo.AddImage(ctx, image)
}
