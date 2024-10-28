package entity

import "github.com/congmanh18/lucas-core/record"

type Category struct {
	record.BaseEntity
	Name        *string `gorm:"size:255;not null"`
	Description *string
}

type ProductCategory struct {
	ProductID  *string `gorm:"foreignKey:ProductID"`
	CategoryID *string `gorm:"foreignKey:CategoryID"`
}
