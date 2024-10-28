package migration

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/congmanh18/XuyenXam/ArtistService/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Must(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Lấy tên collection từ tên struct (vd: Artist => "artists")
	collectionName := getCollectionName(&entity.Artist{})

	// Lấy database từ client
	db := client.Database(dbName)

	// Kiểm tra xem collection đã tồn tại chưa
	collections, err := db.ListCollectionNames(ctx, bson.M{"name": collectionName})
	if err != nil {
		panic(fmt.Sprintf("failed to list collections: %v", err))
	}

	// Tạo collection nếu chưa tồn tại
	if len(collections) == 0 {
		err := db.CreateCollection(ctx, collectionName)
		if err != nil {
			panic(fmt.Sprintf("failed to create collection %s: %v", collectionName, err))
		}
		fmt.Printf("Collection '%s' created successfully.\n", collectionName)
	} else {
		fmt.Printf("Collection '%s' already exists.\n", collectionName)
	}
}

// Lấy tên collection từ struct (vd: Artist => "artists")
func getCollectionName(model interface{}) string {
	typ := reflect.TypeOf(model)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ.Name()
}
