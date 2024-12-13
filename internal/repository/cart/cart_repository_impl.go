package cart

import (
	"context"
	"database/sql"
	"time"

	"github.com/MCPutro/E-commerce/internal/domain"
	"github.com/MCPutro/E-commerce/pkg/constant"
	newError "github.com/MCPutro/E-commerce/pkg/error"
)

type cartRepository struct{}

func (c *cartRepository) Write(ctx context.Context, tx *sql.Tx, cart *domain.Cart) (*domain.Cart, error) {
	now := time.Now()
	query := "INSERT INTO e_commerce.carts (user_id, created_at, updated_at) VALUES (?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, cart.UserId, now.Format(constant.TimeFormat), now.Format(constant.TimeFormat))
	if err != nil {
		return nil, err
	}

	// Get the last inserted ID
	cartId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Set the ID in the cart object
	cart.Id = uint(cartId)
	cart.CreationTime = now // Set the creation time
	cart.UpdateTime = now   // Set the update time

	return cart, nil
}

func (c *cartRepository) ReadByUserId(ctx context.Context, tx *sql.Tx, userId string) (*domain.Cart, error) {
	query := `
		SELECT c.id, c.user_id, c.created_at, c.updated_at, 
		       ci.product_id, ci.quantity
		FROM e_commerce.carts c
		LEFT JOIN e_commerce.cart_items ci ON c.id = ci.cart_id
		WHERE c.user_id = ? AND c.deleted_at IS NULL`
	rows, err := tx.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cart domain.Cart
	cart.Items = []domain.CartItem{}

	for rows.Next() {
		var item domain.CartItem
		if err := rows.Scan(&cart.Id, &cart.UserId, &cart.CreationTime, &cart.UpdateTime, &item.ProductID, &item.Quantity); err != nil {
			return nil, err
		}
		if item.ProductID != 0 {
			cart.Items = append(cart.Items, item)
		}
	}

	if len(cart.Items) == 0 {
		return nil, newError.ErrCartNotFound
	}

	return &cart, nil
}

func (c *cartRepository) Remove(ctx context.Context, tx *sql.Tx, cartId string) error {
	now := time.Now().Format(constant.TimeFormat)
	query := "UPDATE e_commerce.carts SET deleted_at = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, now, cartId)
	return err
}

func (c *cartRepository) AddItem(ctx context.Context, tx *sql.Tx, cartId string, item *domain.CartItem) error {
	now := time.Now().Format(constant.TimeFormat)
	query := "INSERT INTO e_commerce.cart_items (cart_id, product_id, quantity, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, query, cartId, item.ProductID, item.Quantity, now, now)
	return err
}

func (c *cartRepository) RemoveItem(ctx context.Context, tx *sql.Tx, cartId string, productId string) error {
	now := time.Now().Format(constant.TimeFormat)
	query := "UPDATE e_commerce.cart_items SET deleted_at = ? WHERE cart_id = ? AND product_id = ?"
	_, err := tx.ExecContext(ctx, query, now, cartId, productId)
	return err
}

func (c *cartRepository) UpdateItem(ctx context.Context, tx *sql.Tx, cartId string, item *domain.CartItem) error {
	query := "UPDATE e_commerce.cart_items SET quantity = ? WHERE cart_id = ? AND product_id = ? AND deleted_at IS NULL"
	_, err := tx.ExecContext(ctx, query, item.Quantity, cartId, item.ProductID)
	return err
}

func NewCartRepository() Repository {
	return &cartRepository{}
}
