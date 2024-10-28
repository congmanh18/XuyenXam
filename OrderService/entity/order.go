package entity

import (
	"github.com/congmanh18/lucas-core/record"
)

type Order struct {
	record.BaseEntity
	CustomerID *string
	ArtistID   *string
	Status     *string // pending, confirmed, completed, canceled
	TotalPrice *float64
	Items      []OrderItem // Quan hệ 1-N với OrderItem
	Payment    *Payment    // Quan hệ 1-1 với Payment
}
