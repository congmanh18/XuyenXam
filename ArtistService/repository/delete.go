package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Delete xóa một Artist bằng ID
func (r *repoImpl) Delete(ctx context.Context, id *string) error {
	collection := r.db.Database.Collection("artists")
	objectID, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	return nil
}
