package domain

import "time"

type OrderItem struct {
	OrderId      string
	Seq          uint
	ProductId    uint
	Quantity     uint
	UnitPrice    float64
	TotalPrice   float64
	CreationTime time.Time
	UpdateTIme   time.Time
	DeleteTime   time.Time
}

func (u *OrderItem) TableName() string {
	return "order_items"
}
