package product_category

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/lucas-core/db/postgresql"
)

type ProductCategoryRepository interface {
	AssignProductToCategory(ctx context.Context, productID, categoryID *string) error
	RemoveProductFromCategory(ctx context.Context, productID, categoryID *string) error
	GetCategoriesByProduct(ctx context.Context, productID *string) ([]entity.Category, error)
	GetProductsByCategory(ctx context.Context, categoryID *string) ([]entity.Product, error)
}

type productCategoryRepo struct {
	DB *postgresql.Database
}

func NewProductCategoryRepository(db *postgresql.Database) ProductCategoryRepository {
	return &productCategoryRepo{DB: db}
}
