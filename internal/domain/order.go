package domain

import (
	"time"

	"github.com/MCPutro/E-commerce/pkg/constant"
)

type Order struct {
	Id           string
	OrderDate    time.Time
	TotalAmount  float64
	UserId       uint
	AddressSeq   string
	Items        []OrderItem
	Status       constant.OrderStatus
	CreationTime time.Time
	UpdateTIme   time.Time
	DeleteTime   time.Time
}

func (u *Order) TableName() string {
	return "orders"
}
