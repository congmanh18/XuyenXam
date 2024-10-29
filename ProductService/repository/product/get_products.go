package product

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/lucas-core/gorm"
	"github.com/congmanh18/lucas-core/record"
)

// Nên dùng cái này, có phân trang
func (p *productRepo) GetProducts(ctx context.Context, pagination *record.Pagination) ([]entity.Product, error) {
	var products []entity.Product

	// Tạo query để đếm tổng số bản ghi
	countQuery := p.DB.Executor.Model(&entity.Product{})

	err := p.DB.Executor.WithContext(ctx).
		Scopes(gorm.Paginate(pagination, countQuery)).
		Find(&products).Error

	return products, err
}
