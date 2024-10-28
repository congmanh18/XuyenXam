package entity

import "github.com/congmanh18/lucas-core/record"

type OrderItem struct {
	record.BaseEntity
	OrderID     *string
	ProductName *string
	Quantity    *int
	Price       *float64
}
