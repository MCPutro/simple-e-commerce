package domain

import "time"

type Payment struct {
	Id            string
	OrderId       string
	Date          time.Time
	Amount        int
	Status        string
	PaymentMethod string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

func (u *Payment) TableName() string {
	return "payments"
}
