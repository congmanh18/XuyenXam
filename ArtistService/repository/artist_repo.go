package repository

import (
	"context"

	"github.com/congmanh18/XuyenXam/ArtistService/entity"
	"github.com/congmanh18/lucas-core/db/mongodb"
)

type Repo interface {
	Create(ctx context.Context, doc *entity.Artist) error
	FindByID(ctx context.Context, id *string) (*entity.Artist, error)
	FindAll(ctx context.Context) ([]entity.Artist, error)
	Update(ctx context.Context, doc *entity.Artist) error
	Delete(ctx context.Context, id *string) error
}

type repoImpl struct {
	db *mongodb.Database
}

func NewRepo(db *mongodb.Database) Repo {
	return &repoImpl{db: db}
}
