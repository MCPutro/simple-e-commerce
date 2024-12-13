package constant

const (
	TimeFormat              = "2006-01-02 15:04:05"
	HeaderXRequestID string = "XRequestID"
)

type UserRole string

const (
	Customer UserRole = "Customer"
	Staff    UserRole = "Staff"
)

type OrderStatus string

const (
	Pending OrderStatus = "Pending"
)
