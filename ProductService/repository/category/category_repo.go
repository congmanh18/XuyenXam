package category

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/lucas-core/db/postgresql"
)

type CategoryRepository interface {
	GetAllCategories(ctx context.Context) ([]entity.Category, error)
	GetCategoryByID(ctx context.Context, id *string) (*entity.Category, error)
	CreateCategory(ctx context.Context, category *entity.Category) error
	UpdateCategory(ctx context.Context, category *entity.Category, id *string) error
	DeleteCategory(ctx context.Context, id *string) error
	GetProductsByCategory(ctx context.Context, categoryID *string) ([]entity.Product, error)
}

type categoryRepo struct {
	DB *postgresql.Database
}

func NewCategoryRepository(db *postgresql.Database) CategoryRepository {
	return &categoryRepo{DB: db}
}
