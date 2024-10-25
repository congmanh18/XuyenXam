package repository

import (
	"context"

	"github.com/congmanh18/XuyenXam/ArtistService/entity"
)

// Create thêm mới một Artist vào MongoDB
func (r *repoImpl) Create(ctx context.Context, doc *entity.Artist) error {
	collection := r.db.Database.Collection("artists")
	_, err := collection.InsertOne(ctx, doc)

	return err
}
