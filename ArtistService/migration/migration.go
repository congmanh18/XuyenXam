package migration

import (
	"github.com/congmanh18/XuyenXam/ArtistService/entity"
	"gorm.io/gorm"
)

func Must(db *gorm.DB) {
	db.Exec("CREATE SCHEMA IF NOT EXISTS user_service")
	err := db.Debug().AutoMigrate(
		&entity.Artist{},
	)
	if err != nil {
		panic(err)
	}
}
