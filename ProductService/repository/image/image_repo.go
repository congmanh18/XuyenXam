package image

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/lucas-core/db/postgresql"
)

type ImageRepository interface {
	AddImage(ctx context.Context, image *entity.Image) error
	GetImageByProductID(ctx context.Context, productID string) ([]entity.Image, error)
	DeleteImage(ctx context.Context, id string) error
}

type imageRepo struct {
	DB *postgresql.Database
}

// AddImage implements ImageRepository.
func (i *imageRepo) AddImage(ctx context.Context, image *entity.Image) error {
	panic("unimplemented")
}

// DeleteImage implements ImageRepository.
func (i *imageRepo) DeleteImage(ctx context.Context, id string) error {
	panic("unimplemented")
}

// GetImageByProductID implements ImageRepository.
func (i *imageRepo) GetImageByProductID(ctx context.Context, productID string) ([]entity.Image, error) {
	panic("unimplemented")
}

func NewImageRepository(db *postgresql.Database) ImageRepository {
	return &imageRepo{DB: db}
}
