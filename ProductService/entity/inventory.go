package entity

import "github.com/congmanh18/lucas-core/record"

type Inventory struct {
	record.BaseEntity
	ProductID *string `gorm:"not null"`
	Quantity  *int    `gorm:"not null"` // số lượng tồn kho
}
