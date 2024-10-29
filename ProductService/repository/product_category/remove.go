package product_category

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (p *productCategoryRepo) RemoveProductFromCategory(ctx context.Context, productID, categoryID *string) error {
	if err := p.DB.Executor.WithContext(ctx).
		Delete(&entity.ProductCategory{ProductID: productID, CategoryID: categoryID}).Error; err != nil {
		return err
	}
	return nil
}
