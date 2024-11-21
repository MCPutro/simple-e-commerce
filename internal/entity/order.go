package entity

import "time"

type Order struct {
	Id           int
	UserId       uint
	TotalPrice   float64
	Items        []OrderItem
	Status       string
	CreationTime time.Time
	UpdateTIme   time.Time
	DeleteTime   time.Time
}
