package domain

import "time"

type Product struct {
	Id           uint
	Name         string
	Price        float64
	Stock        int
	Description  string
	CreationTime time.Time
	UpdateTime   time.Time
	DeleteTime   time.Time
}

func (u *Product) TableName() string {
	return "products"
}
