package product_category

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
)

func (p *productCategoryRepo) AssignProductToCategory(ctx context.Context, productID, categoryID *string) error {
	err := p.DB.Executor.WithContext(ctx).Create(&entity.ProductCategory{ProductID: productID, CategoryID: categoryID}).Error
	if err != nil {
		return err
	}
	return nil
}
