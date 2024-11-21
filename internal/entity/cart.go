package entity

import "time"

type Cart struct {
	Id           uint
	UserId       uint
	Items        []CartItem
	CreationTime time.Time
	UpdateTime   time.Time
	DeleteTime   time.Time
}
