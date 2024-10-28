package entity

import (
	"github.com/congmanh18/lucas-core/record"
)

type Payment struct {
	record.BaseEntity
	OrderID *string
	Amount  *float64
	Status  *string
}
