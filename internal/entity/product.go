package entity

import "time"

type Product struct {
	Id           uint
	Name         string
	Price        float64
	Stock        int
	CreationTime time.Time
	UpdateTime   time.Time
	DeleteTime   time.Time
}
