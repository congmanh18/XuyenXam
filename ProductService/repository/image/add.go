package image

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

// AddImage implements ImageRepository.
func (i *imageRepo) AddImage(ctx context.Context, image *entity.Image) error {
	return i.DB.Executor.WithContext(ctx).Create(image).Error
}
