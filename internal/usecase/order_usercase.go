package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/MCPutro/E-commerce/internal/entity"
	"github.com/MCPutro/E-commerce/internal/repository"
	newError "github.com/MCPutro/E-commerce/pkg/error"
	"github.com/MCPutro/E-commerce/pkg/logger"
)

type OrderUseCase interface {
	Checkout(ctx context.Context, userID uint) (*entity.Order, error)
	GetOrder(ctx context.Context, orderID string) (*entity.Order, error)
}

type orderUsecase struct {
	productRepo repository.ProductRepository
	cartRepo    repository.CartRepository
	orderRepo   repository.OrderRepository
	db          *sql.DB
}

func (o *orderUsecase) Checkout(ctx context.Context, userID uint) (*entity.Order, error) {
	logger.ContextLogger(ctx).Infof("Checkout userId: %d", userID)

	// Start transaction
	tx, err := o.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, newError.ErrOpenTransactionWithDetails(err.Error())
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Get cart with FOR UPDATE lock
	cart, err := o.cartRepo.GetCartByUserId(ctx, tx, userID)
	if err != nil {
		return nil, err
	}
	logger.ContextLogger(ctx).Infof("%+v", cart)

	// Get cart items with FOR UPDATE lock
	cartItems, err := o.cartRepo.GetCartItems(ctx, tx, cart.Id)
	if err != nil {
		return nil, err
	}

	if len(cartItems) == 0 {
		return nil, errors.New("cart is empty")
	}

	// Validate cart items and check stock
	err = o.cartRepo.ValidateCartItems(ctx, tx, cart.Id)
	if err != nil {
		return nil, fmt.Errorf("cart validation failed: %w", err)
	}

	// Calculate total price and prepare order items
	var totalPrice float64
	orderItems := make([]entity.OrderItem, 0)
	for _, item := range cartItems {
		// Get product with FOR UPDATE lock
		product, err := o.productRepo.GetProductByID(ctx, tx, item.ProductID)
		if err != nil {
			return nil, fmt.Errorf("failed to get product %d: %w", item.ProductID, err)
		}

		// Update stock
		newStock := product.Stock - int(item.Quantity)
		err = o.productRepo.UpdateStock(ctx, tx, product.Id, newStock)
		if err != nil {
			return nil, fmt.Errorf("failed to update stock for product %d: %w", product.Id, err)
		}

		// Calculate item total price
		itemTotal := float64(item.Quantity) * product.Price
		totalPrice += itemTotal

		// Create order item
		orderItems = append(orderItems, entity.OrderItem{
			ProductId:  product.Id,
			Quantity:   item.Quantity,
			TotalPrice: itemTotal,
		})
	}

	// Create order
	order := &entity.Order{
		UserId:     userID,
		TotalPrice: totalPrice,
		Items:      orderItems,
		Status:     "PENDING",
	}

	err = o.orderRepo.CreateOrder(ctx, tx, order)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	// Clear cart
	err = o.cartRepo.ClearCart(ctx, tx, cart.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to clear cart: %w", err)
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return order, nil
}

func (o *orderUsecase) GetOrder(ctx context.Context, orderID string) (*entity.Order, error) {
	tx, err := o.db.Begin()
	if err != nil {
		return nil, newError.ErrOpenTransactionWithDetails(err.Error())
	}
	defer tx.Rollback()

	order, err := o.orderRepo.GetOrderByID(ctx, tx, orderID)
	if err != nil {
		return nil, err
	}

	return order, tx.Commit()
}

func NewOrderUseCase(productRepo repository.ProductRepository, cartRepo repository.CartRepository, orderRepo repository.OrderRepository, db *sql.DB) OrderUseCase {
	return &orderUsecase{
		productRepo: productRepo,
		cartRepo:    cartRepo,
		orderRepo:   orderRepo,
		db:          db,
	}
}
