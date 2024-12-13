package domain

import (
	"time"

	"github.com/MCPutro/E-commerce/pkg/constant"
)

type User struct {
	Id        string
	Name      string
	Email     string
	Password  string
	Role      constant.UserRole
	Address   []UserAddress
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (u *User) TableName() string {
	return "users"
}
