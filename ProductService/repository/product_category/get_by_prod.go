package product_category

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (p *productCategoryRepo) GetCategoriesByProduct(ctx context.Context, productID *string) ([]entity.Category, error) {
	var categoryIDs []string
	if err := p.DB.Executor.WithContext(ctx).
		Model(&entity.ProductCategory{}).
		Where("product_id = ?", productID).
		Pluck("category_id", &categoryIDs).Error; err != nil {
		return nil, err
	}

	var categories []entity.Category
	for _, categoryID := range categoryIDs {
		if err := p.DB.Executor.WithContext(ctx).
			Where("id = (?)", categoryID).
			Find(&categories).Error; err != nil {
			return nil, err
		}
	}

	return categories, nil
}
