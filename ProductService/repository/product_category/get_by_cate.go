package product_category

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (p *productCategoryRepo) GetProductsByCategory(ctx context.Context, categoryID *string) ([]entity.Product, error) {
	var productIDs []string
	if err := p.DB.Executor.WithContext(ctx).
		Model(&entity.ProductCategory{}).
		Where("category_id = ?", categoryID).
		Pluck("product_id", &productIDs).Error; err != nil {
		return nil, err
	}

	var products []entity.Product
	for _, productID := range productIDs {
		if err := p.DB.Executor.WithContext(ctx).
			Where("id = (?)", productID).
			Find(&products).Error; err != nil {
			return nil, err
		}
	}

	return products, nil
}
