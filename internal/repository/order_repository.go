package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/MCPutro/E-commerce/internal/entity"
	"github.com/MCPutro/E-commerce/pkg/constant"
)

type OrderRepository interface {
	Write(ctx context.Context, tx *sql.Tx, order *entity.Order) error
	ReadByID(ctx context.Context, tx *sql.Tx, id string) (*entity.Order, error)
}

type orderRepository struct{}

func (o *orderRepository) Write(ctx context.Context, tx *sql.Tx, order *entity.Order) error {
	now := time.Now().Format(constant.TimeFormat)
	query1 := "INSERT INTO e_commerce.orders (user_id,total_price,status,creation_time,update_time) VALUES (?,?,?,?,?);"
	result, err := tx.ExecContext(ctx, query1, order.UserId, order.TotalPrice, order.Status, now, now)
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

func (o *orderRepository) ReadByID(ctx context.Context, tx *sql.Tx, id string) (*entity.Order, error) {
	//TODO implement me
	panic("implement me")
}

func NewOrderRepository() OrderRepository {
	return &orderRepository{}
}
