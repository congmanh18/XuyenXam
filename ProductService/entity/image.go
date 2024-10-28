package entity

import (
	"github.com/congmanh18/lucas-core/record"
)

type Image struct {
	record.BaseEntity
	ProductID int    `gorm:"not null"`
	URL       string `gorm:"not null"`
}
