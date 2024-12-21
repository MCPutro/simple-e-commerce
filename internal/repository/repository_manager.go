package repository

import (
	"sync"

	"github.com/MCPutro/E-commerce/internal/repository/cart"
	"github.com/MCPutro/E-commerce/internal/repository/order"
	"github.com/MCPutro/E-commerce/internal/repository/payment"
	"github.com/MCPutro/E-commerce/internal/repository/product"
	"github.com/MCPutro/E-commerce/internal/repository/user"
)

var (
	userRepo     user.Repository
	userRepoOnce sync.Once

	productRepo     product.Repository
	productRepoOnce sync.Once

	paymentRepo     payment.Repository
	paymentRepoOnce sync.Once

	orderRepo     order.Repository
	orderRepoOnce sync.Once

	cartRepo     cart.Repository
	cartRepoOnce sync.Once
)

type RepositoryManager interface {
	UserRepo() user.Repository
	ProductRepo() product.Repository
	PaymentRepo() payment.Repository
	OrderRepo() order.Repository
	CartRepo() cart.Repository
}

type repositoryManager struct{}

// CartRepo implements RepositoryManager.
func (r *repositoryManager) CartRepo() cart.Repository {
	cartRepoOnce.Do(func() {
		cartRepo = cart.NewCartRepository()
	})

	return cartRepo
}

// OrderRepo implements RepositoryManager.
func (r *repositoryManager) OrderRepo() order.Repository {
	orderRepoOnce.Do(func() {
		orderRepo = order.NewOrderRepository()
	})
	return orderRepo
}

// PaymentRepo implements RepositoryManager.
func (r *repositoryManager) PaymentRepo() payment.Repository {
	paymentRepoOnce.Do(func() {
		paymentRepo = payment.NewPaymentRepository()
	})
	return paymentRepo
}

// ProductRepo implements RepositoryManager.
func (r *repositoryManager) ProductRepo() product.Repository {
	productRepoOnce.Do(func() {
		productRepo = product.NewProductRepository()
	})
	return productRepo
}

// UserRepo implements RepositoryManager.
func (r *repositoryManager) UserRepo() user.Repository {
	userRepoOnce.Do(func() {
		userRepo = user.NewUserRepository()
	})
	return userRepo
}

func NewRepositoryManager() RepositoryManager {
	return &repositoryManager{}
}
