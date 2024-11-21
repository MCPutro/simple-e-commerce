package entity

import "time"

type OrderItem struct {
	TrxId        int
	Seq          uint
	ProductId    uint
	Quantity     uint
	TotalPrice   float64
	CreationTime time.Time
	UpdateTIme   time.Time
	DeleteTime   time.Time
}
