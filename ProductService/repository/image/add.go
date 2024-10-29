package image

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

// AddImage implements ImageRepository.
func (i *imageRepo) AddImage(ctx context.Context, image *entity.Image) error {
	err := i.DB.Executor.WithContext(ctx).Create(image).Error
	if err != nil {
		return err
	}
	return nil
}
