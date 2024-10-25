package repository

import (
	"context"

	"github.com/congmanh18/XuyenXam/ArtistService/entity"
	"go.mongodb.org/mongo-driver/bson"
)

// FindAll trả về tất cả Artist từ MongoDB
func (r *repoImpl) FindAll(ctx context.Context) ([]entity.Artist, error) {
	collection := r.db.Database.Collection("artists")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var artists []entity.Artist
	for cursor.Next(ctx) {
		var artist entity.Artist
		if err := cursor.Decode(&artist); err != nil {
			return nil, err
		}
		artists = append(artists, artist)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return artists, nil
}
