package image

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/lucas-core/db/postgresql"
)

type ImageRepository interface {
	AddImage(ctx context.Context, image *entity.Image) error
	GetImageByProductID(ctx context.Context, productID *string) ([]entity.Image, error)
	DeleteImage(ctx context.Context, id *string) error
}

type imageRepo struct {
	DB *postgresql.Database
}

func NewImageRepository(db *postgresql.Database) ImageRepository {
	return &imageRepo{DB: db}
}
