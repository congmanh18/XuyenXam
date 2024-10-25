package repository

import (
	"context"
	"errors"

	"github.com/congmanh18/XuyenXam/ArtistService/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// FindByID tìm kiếm Artist bằng ID
func (r *repoImpl) FindByID(ctx context.Context, id *string) (*entity.Artist, error) {
	collection := r.db.Database.Collection("artists")
	objectID, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	var artist entity.Artist
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&artist)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("artist not found")
		}
		return nil, err
	}
	return &artist, nil
}
