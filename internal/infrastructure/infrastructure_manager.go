package infrastructure

import (
	"sync"

	"github.com/MCPutro/E-commerce/internal/infrastructure/db"
	"github.com/MCPutro/E-commerce/internal/infrastructure/middleware"
	"github.com/MCPutro/E-commerce/internal/infrastructure/payment_gateway"
)

var (
	middlewareInstance middleware.Middleware
	middlewareOnce     sync.Once

	paymentGatewayInstance payment_gateway.PaymentGateway
	paymentGatewayOnce     sync.Once
)

type InfrastructureManager interface {
	DB() *db.DB
	Middleware() middleware.Middleware
	PaymentGateway() payment_gateway.PaymentGateway
}

type infrastructureManager struct{}

// DB implements InfrastructureManager.
func (im *infrastructureManager) DB() *db.DB {
	dbOnce.Do(func() {
		var err error
		dbInstance, err = db.GetMysqlConnection()
		if err != nil {
			panic("Failed to connect to database: " + err.Error())
		}
	})
	return dbInstance
}

// Middleware implements InfrastructureManager.
func (im *infrastructureManager) Middleware() middleware.Middleware {
	middlewareOnce.Do(func() {
		middlewareInstance = middleware.NewMiddleware()
	})
	return middlewareInstance
}

// PaymentGateway implements InfrastructureManager.
func (im *infrastructureManager) PaymentGateway() payment_gateway.PaymentGateway {
	paymentGatewayOnce.Do(func() {
		paymentGatewayInstance = payment_gateway.NewPaymentGateway()
	})
	return paymentGatewayInstance
}

// NewInfrastructureManager creates a new instance of InfrastructureManager.
func NewInfrastructureManager() InfrastructureManager {
	return &infrastructureManager{}
}
