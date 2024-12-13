package payment

import (
	"context"
	"database/sql"
	"time"

	"github.com/MCPutro/E-commerce/internal/domain"
	"github.com/MCPutro/E-commerce/pkg/constant"
	newError "github.com/MCPutro/E-commerce/pkg/error"
)

type paymentRepository struct{}

func (r *paymentRepository) Write(ctx context.Context, tx *sql.Tx, payment *domain.Payment) error {
	now := time.Now().Format(constant.TimeFormat)
	query := "INSERT INTO e_commerce.payments (id, order_id, payment_date, amount, status, payment_method, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"
	_, err := tx.ExecContext(ctx, query, payment.Id, payment.OrderId, payment.Date, payment.Amount, payment.Status, payment.PaymentMethod, now, now)
	return err
}

func (r *paymentRepository) ReadById(ctx context.Context, tx *sql.Tx, id string) (*domain.Payment, error) {
	query := "SELECT id, order_id, payment_date, amount, status, payment_method, created_at, updated_at FROM e_commerce.payments WHERE id = ? AND deleted_at IS NULL;"
	row := tx.QueryRowContext(ctx, query, id)

	var payment domain.Payment
	var createdAt, updatedAt sql.NullTime
	if err := row.Scan(&payment.Id, &payment.OrderId, &payment.Date, &payment.Amount, &payment.Status, &payment.PaymentMethod, &createdAt, &updatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found
		}
		return nil, err
	}

	if createdAt.Valid {
		payment.CreatedAt = createdAt.Time
	}
	if updatedAt.Valid {
		payment.UpdatedAt = updatedAt.Time
	}

	return &payment, nil
}

func (r *paymentRepository) ReadAll(ctx context.Context, tx *sql.Tx) ([]domain.Payment, error) {
	query := "SELECT id, order_id, payment_date, amount, status, payment_method, created_at, updated_at FROM e_commerce.payments WHERE deleted_at IS NULL;"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []domain.Payment
	for rows.Next() {
		var payment domain.Payment
		var createdAt, updatedAt sql.NullTime
		if err := rows.Scan(&payment.Id, &payment.OrderId, &payment.Date, &payment.Amount, &payment.Status, &payment.PaymentMethod, &createdAt, &updatedAt); err != nil {
			return nil, err
		}
		if createdAt.Valid {
			payment.CreatedAt = createdAt.Time
		}
		if updatedAt.Valid {
			payment.UpdatedAt = updatedAt.Time
		}
		payments = append(payments, payment)
	}

	return payments, nil
}

func (r *paymentRepository) Update(ctx context.Context, tx *sql.Tx, payment *domain.Payment) error {
	query := "UPDATE e_commerce.payments SET order_id = ?, payment_date = ?, amount = ?, status = ?, payment_method = ?, updated_at = NOW() WHERE id = ?;"
	_, err := tx.ExecContext(ctx, query, payment.OrderId, payment.Date, payment.Amount, payment.Status, payment.PaymentMethod, payment.Id)
	return err
}

func (r *paymentRepository) Delete(ctx context.Context, tx *sql.Tx, id string) error {
	query := "UPDATE e_commerce.payments SET deleted_at = NOW() WHERE id = ?;"
	result, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return newError.ErrProductNotFound
	}
	return nil
}

func NewPaymentRepository() Repository {
	return &paymentRepository{}
}
