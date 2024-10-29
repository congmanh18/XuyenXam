package image_usecase

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/XuyenXam/ProductService/repository/image"
)

type GetImageUseCase interface {
	Execute(ctx context.Context, productID *string) ([]entity.Image, error)
}

type getImageUseCaseImpl struct {
	ImageRepo image.ImageRepository
}

func NewGetImageUseCase(imageRepo image.ImageRepository) GetImageUseCase {
	return &getImageUseCaseImpl{ImageRepo: imageRepo}
}

func (u *getImageUseCaseImpl) Execute(ctx context.Context, productID *string) ([]entity.Image, error) {
	return u.ImageRepo.GetImageByProductID(ctx, productID)
}
