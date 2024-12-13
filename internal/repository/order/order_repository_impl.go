package order

import (
	"context"
	"database/sql"
	"time"

	"github.com/MCPutro/E-commerce/internal/domain"
	"github.com/MCPutro/E-commerce/pkg/constant"
	newError "github.com/MCPutro/E-commerce/pkg/error"
)

type orderRepository struct{}

func (o *orderRepository) Write(ctx context.Context, tx *sql.Tx, order *domain.Order) error {
	now := time.Now().Format(constant.TimeFormat)
	query1 := "INSERT INTO e_commerce.orders (user_id,total_price,status,creation_time,update_time) VALUES (?,?,?,?,?);"
	result, err := tx.ExecContext(ctx, query1, order.UserId, order.TotalAmount, order.Status, now, now)
	if err != nil {
		return err
	}

	orderId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, "INSERT INTO e_commerce.order_items (trx_id,seq,product_id,quantity,total_price,creation_time,update_time) VALUES (?,?,?,?,?,?,?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for i, item := range order.Items {
		_, err2 := stmt.ExecContext(ctx, orderId, i+1, item.ProductId, item.Quantity, item.TotalPrice, now, now)
		if err2 != nil {
			return err2
		}
	}
	return nil
}

func (o *orderRepository) ReadByID(ctx context.Context, tx *sql.Tx, id string) (*domain.Order, error) {
	query := "SELECT id, user_id, total_amount, status, creation_time, update_time FROM e_commerce.orders WHERE id = ? AND deleted_at IS NULL;"
	row := tx.QueryRowContext(ctx, query, id)

	var order domain.Order
	var createdAt, updatedAt sql.NullTime
	if err := row.Scan(&order.Id, &order.UserId, &order.TotalAmount, &order.Status, &createdAt, &updatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, newError.ErrProductNotFound
		}
		return nil, err
	}

	if createdAt.Valid {
		order.CreationTime = createdAt.Time
	}
	if updatedAt.Valid {
		order.UpdateTIme = updatedAt.Time
	}

	return &order, nil
}

func NewOrderRepository() Repository {
	return &orderRepository{}
}
