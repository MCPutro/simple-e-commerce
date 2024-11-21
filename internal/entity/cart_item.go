package entity

import "time"

type CartItem struct {
	CartId       uint
	ProductID    uint
	Quantity     uint
	Product      *Product
	CreationTime time.Time
	UpdateTime   time.Time
	DeleteTime   time.Time
}
