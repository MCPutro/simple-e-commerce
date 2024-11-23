package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/MCPutro/E-commerce/internal/entity"
	"github.com/MCPutro/E-commerce/internal/repository"
	newError "github.com/MCPutro/E-commerce/pkg/error"
)

type CartUseCase interface {
	AddToCart(ctx context.Context, userID uint, item *entity.CartItem) error
	GetCart(ctx context.Context, userID uint) (*entity.Cart, error)
}

type cartUsecase struct {
	cartRepo    repository.CartRepository
	productRepo repository.ProductRepository
	db          *sql.DB
}

func (c *cartUsecase) AddToCart(ctx context.Context, userID uint, item *entity.CartItem) error {
	if item == nil {
		return errors.New("cart item cannot be nil")
	}

	if userID == 0 {
		return errors.New("invalid user ID")
	}

	if item.Quantity <= 0 {
		return errors.New("quantity must be positive")
	}

	tx, err := c.db.Begin()
	if err != nil {
		return newError.ErrOpenTransactionWithDetails(err.Error())
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Cek ketersediaan produk dan stok
	product, err := c.productRepo.ReadByID(ctx, tx, item.ProductID)
	if err != nil {
		return newError.ErrProductNotFound
	}

	if product == nil {
		return newError.ErrProductNotFound
	}

	// if product.Stock < int(item.Quantity) {
	// 	return fmt.Errorf("stok tidak mencukupi")
	// }

	// Tambahkan ke keranjang
	err = c.cartRepo.WriteCart(ctx, tx, userID, item)
	if err != nil {
		return fmt.Errorf("gagal menambahkan ke keranjang: %w", err)
	}

	return tx.Commit()
}

func (c *cartUsecase) GetCart(ctx context.Context, userID uint) (*entity.Cart, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return nil, newError.ErrOpenTransactionWithDetails(err.Error())
	}
	defer tx.Rollback()

	cart, err := c.cartRepo.ReadCartByUserId(ctx, tx, userID)
	if err != nil {
		return nil, err
	}

	// Ambil items
	items, err := c.cartRepo.ReadCartItemsById(ctx, tx, cart.Id)
	if err != nil {
		return nil, err
	}

	cart.Items = items
	return cart, tx.Commit()
}

func NewCartUseCase(cartRepo repository.CartRepository, productRepo repository.ProductRepository, db *sql.DB) CartUseCase {
	return &cartUsecase{
		cartRepo:    cartRepo,
		productRepo: productRepo,
		db:          db,
	}
}
