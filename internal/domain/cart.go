package domain

import "time"

type Cart struct {
	Id           uint
	UserId       uint
	Items        []CartItem
	CreationTime time.Time
	UpdateTime   time.Time
	DeleteTime   time.Time
}

func (u *Cart) TableName() string {
	return "carts"
}
