package domain

import (
	"time"
)

type UserAddress struct {
	UserId     string
	Seq        int
	Address    string
	City       string
	PostalCode string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

func (u *UserAddress) TableName() string {
	return "user_addresses"
}
