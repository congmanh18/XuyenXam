package image

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

// DeleteImage implements ImageRepository.
func (i *imageRepo) DeleteImage(ctx context.Context, id string) error {
	return i.DB.Executor.WithContext(ctx).Delete(&entity.Image{}, id).Error
}