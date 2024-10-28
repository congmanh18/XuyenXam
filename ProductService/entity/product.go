package entity

import (
	"github.com/congmanh18/lucas-core/record"
)

type Product struct {
	record.BaseEntity
	Name        *string `gorm:"size:255;not null"`
	Description *string
	Price       *float64 `gorm:"not null"`
}
