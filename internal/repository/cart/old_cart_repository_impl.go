package cart

// import (
// 	"context"
// 	"database/sql"
// 	"errors"
// 	"fmt"
// 	"time"

// 	"github.com/MCPutro/E-commerce/internal/domain"
// 	"github.com/MCPutro/E-commerce/pkg/constant"
// 	newError "github.com/MCPutro/E-commerce/pkg/error"
// )

// type oldcartRepository struct{}

// func (c *cartRepository) Write(ctx context.Context, tx *sql.Tx, userID uint, item *domain.CartItem) error {

// 	existingCart, err := c.ReadCartByUserId(ctx, tx, userID)
// 	if err != nil && !errors.Is(err, newError.ErrCartNotFound) {
// 		return err
// 	}

// 	now := time.Now().Format(constant.TimeFormat)
// 	var cartId int64

// 	if existingCart == nil {
// 		// Buat cart baru jika belum ada
// 		result, err := tx.ExecContext(ctx,
// 			"INSERT INTO e_commerce.carts (user_id,creation_time,update_time) VALUES (?,?,?);",
// 			userID, now, now)
// 		if err != nil {
// 			return err
// 		}
// 		cartId, err = result.LastInsertId()
// 		if err != nil {
// 			return err
// 		}
// 	} else {
// 		cartId = int64(existingCart.Id)
// 		// Cek item yang sudah ada
// 		var existingQty int
// 		err = tx.QueryRowContext(ctx,
// 			"SELECT quantity FROM e_commerce.cart_items WHERE cart_id = ? AND product_id = ? AND delete_time IS NULL",
// 			cartId, item.ProductID).Scan(&existingQty)

// 		if err == nil {
// 			// Update quantity jika item sudah ada
// 			_, err = tx.ExecContext(ctx,
// 				"UPDATE e_commerce.cart_items SET quantity = ?, update_time = ? WHERE cart_id = ? AND product_id = ?",
// 				existingQty+int(item.Quantity), now, cartId, item.ProductID)
// 			return err
// 		}

// 		if !errors.Is(err, sql.ErrNoRows) {
// 			return err
// 		}
// 	}

// 	// Insert item baru
// 	_, err = tx.ExecContext(ctx,
// 		"INSERT INTO e_commerce.cart_items (cart_id,product_id,quantity,creation_time,update_time) VALUES (?,?,?,?,?);",
// 		cartId, item.ProductID, item.Quantity, now, now)
// 	return err
// }

// func (c *oldcartRepository) ReadCartByUserId(ctx context.Context, tx *sql.Tx, userID uint) (*domain.Cart, error) {

// 	query := "SELECT id, user_id, creation_time, update_time, delete_time FROM e_commerce.carts WHERE user_id = ? AND delete_time IS NULL FOR UPDATE;"

// 	rows, err := tx.QueryContext(ctx, query, userID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	if rows.Next() {
// 		var cart domain.Cart
// 		var createdAt, updateTime, deleteTime sql.NullTime
// 		err := rows.Scan(&cart.Id, &cart.UserId, &createdAt, &updateTime, &deleteTime)
// 		if err != nil {
// 			return nil, err
// 		}

// 		if createdAt.Valid {
// 			cart.CreationTime = createdAt.Time
// 		}

// 		if updateTime.Valid {
// 			cart.UpdateTime = updateTime.Time
// 		}

// 		return &cart, nil
// 	}

// 	return nil, newError.ErrCartNotFound
// }

// func (c *oldcartRepository) ReadCartItemsById(ctx context.Context, tx *sql.Tx, cartID uint) ([]domain.CartItem, error) {
// 	if cartID == 0 {
// 		return nil, errors.New("invalid cart ID")
// 	}

// 	query2 := "SELECT cart_id, product_id, quantity, creation_time, update_time, delete_time FROM e_commerce.cart_items where cart_id = ? and delete_time is null FOR UPDATE;"
// 	row, err := tx.QueryContext(ctx, query2, cartID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer row.Close()

// 	var cartItems = make([]domain.CartItem, 0)
// 	for row.Next() {
// 		var item domain.CartItem
// 		var createdAt, updateTime, deleteTime sql.NullTime
// 		err := row.Scan(&item.CartId, &item.ProductID, &item.Quantity, &createdAt, &updateTime, &deleteTime)
// 		if err != nil {
// 			return nil, err
// 		}
// 		if createdAt.Valid {
// 			item.CreationTime = createdAt.Time
// 		}

// 		if updateTime.Valid {
// 			item.UpdateTime = updateTime.Time
// 		}

// 		cartItems = append(cartItems, item)
// 	}
// 	return cartItems, nil
// }

// func (c *oldcartRepository) RemoveCart(ctx context.Context, tx *sql.Tx, cartID uint) error {
// 	now := time.Now().Format(constant.TimeFormat)
// 	queryRemoveCard := "UPDATE e_commerce.carts SET delete_time = ? WHERE id = ?"
// 	result, err := tx.ExecContext(ctx, queryRemoveCard, now, cartID)
// 	if err != nil {
// 		return err
// 	}

// 	res, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}

// 	if res == 0 {
// 		return errors.New("failed to clear cart")
// 	} else {
// 		return nil
// 	}
// }

// func (c *oldcartRepository) ReadAllCartId(ctx context.Context, tx *sql.Tx) ([]uint, error) {
// 	query := "SELECT id FROM e_commerce.carts where delete_time is null FOR UPDATE;"
// 	rows, err := tx.QueryContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var ids []uint
// 	for rows.Next() {
// 		var id uint
// 		err = rows.Scan(&id)
// 		if err != nil {
// 			return nil, err
// 		}
// 		ids = append(ids, id)
// 	}

// 	if len(ids) == 0 {
// 		return nil, newError.ErrCartNotFound
// 	}

// 	return ids, nil
// }

// func (c *oldcartRepository) UpdateCartItemQuantity(ctx context.Context, tx *sql.Tx, cartID uint, productID uint, quantity int) error {
// 	if quantity <= 0 {
// 		return errors.New("quantity must be positive")
// 	}

// 	now := time.Now().Format(constant.TimeFormat)
// 	query := "UPDATE e_commerce.cart_items  SET quantity = ?, update_time = ? 	WHERE cart_id = ? AND product_id = ? AND delete_time IS NULL;"

// 	result, err := tx.ExecContext(ctx, query, quantity, now, cartID, productID)
// 	if err != nil {
// 		return err
// 	}

// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}

// 	if rowsAffected == 0 {
// 		return newError.ErrCartItemNotFound
// 	}

// 	return nil
// }

// func (c *oldcartRepository) RemoveCartItem(ctx context.Context, tx *sql.Tx, cartID uint, productID uint) error {
// 	now := time.Now().Format(constant.TimeFormat)
// 	query := "UPDATE e_commerce.cart_items  SET delete_time = ? WHERE cart_id = ? AND product_id = ? AND delete_time IS NULL"

// 	result, err := tx.ExecContext(ctx, query, now, cartID, productID)
// 	if err != nil {
// 		return err
// 	}

// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}

// 	if rowsAffected == 0 {
// 		return errors.New("cart item not found")
// 	}

// 	return nil
// }

// func (c *oldcartRepository) ValidateCartItems(ctx context.Context, tx *sql.Tx, cartID uint) error {
// 	query := "SELECT ci.product_id, ci.quantity, p.stock FROM e_commerce.cart_items ci JOIN e_commerce.products p ON ci.product_id = p.id WHERE ci.cart_id = ? AND ci.delete_time IS NULL FOR UPDATE;"

// 	rows, err := tx.QueryContext(ctx, query, cartID)
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()

// 	var invalidItems []uint
// 	for rows.Next() {
// 		var productID uint
// 		var quantity, stock int
// 		// var status string

// 		err := rows.Scan(&productID, &quantity, &stock)
// 		if err != nil {
// 			return err
// 		}

// 		// // Validasi status produk
// 		// if status != "active" {
// 		// 	invalidItems = append(invalidItems, productID)
// 		// 	continue
// 		// }

// 		// Validasi stock
// 		if quantity > stock {
// 			invalidItems = append(invalidItems, productID)
// 		}
// 	}

// 	if len(invalidItems) > 0 {
// 		return fmt.Errorf("invalid items found in cart: %v", invalidItems)
// 	}

// 	return nil
// }

// // func NewoldcartRepository() Repository {
// // 	return &oldcartRepository{}
// // }
