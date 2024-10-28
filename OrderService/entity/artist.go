package entity

import "github.com/congmanh18/lucas-core/record"

type Artist struct {
	record.BaseEntity
	Name           *string
	Specialization *string
	Orders         []Order // Quan hệ 1-N với Order
}
