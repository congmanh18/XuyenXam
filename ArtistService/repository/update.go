package repository

import (
	"context"

	"github.com/congmanh18/XuyenXam/ArtistService/entity"
	"go.mongodb.org/mongo-driver/bson"
)

// Update cập nhật thông tin Artist
func (r *repoImpl) Update(ctx context.Context, doc *entity.Artist) error {
	collection := r.db.Database.Collection("artists")
	filter := bson.M{"_id": doc.ID}

	update := bson.M{
		"$set": doc, // Cập nhật toàn bộ thông tin
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
