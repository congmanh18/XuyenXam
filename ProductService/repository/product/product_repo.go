package product

import (
	"context"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/lucas-core/db/postgresql"
)

type ProductRepository interface {
	GetAllProducts(ctx context.Context) ([]entity.Product, error)
	GetProductByID(ctx context.Context, id *string) (*entity.Product, error)
	CreateProduct(ctx context.Context, product *entity.Product) error
	UpdateProduct(ctx context.Context, product *entity.Product, id *string) error
	DeleteProduct(ctx context.Context, id *string) error
}

type productRepo struct {
	DB *postgresql.Database
}

func NewProductRepository(db *postgresql.Database) ProductRepository {
	return &productRepo{DB: db}
}
