package entity

import "github.com/congmanh18/lucas-core/record"

type Customer struct {
	record.BaseEntity
	Name        *string
	Email       *string
	PhoneNumber *string
}
