package image

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

// GetImageByProductID implements ImageRepository.
func (i *imageRepo) GetImageByProductID(ctx context.Context, productID *string) ([]entity.Image, error) {
	var images []entity.Image

	if err := i.DB.Executor.WithContext(ctx).Where("product_id = ?", productID).Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}
