package domain

import "time"

type User struct {
	ID         uint64
	Login      string
	Email      string
	Phone      string
	FirstName  string
	LastName   string
	MiddleName string
	Age        uint32
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
