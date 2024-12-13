package payment

import (
	"context"
	"database/sql"

	"github.com/MCPutro/E-commerce/internal/domain"
)

type Repository interface {
	Write(ctx context.Context, tx *sql.Tx, payment *domain.Payment) error
	ReadById(ctx context.Context, tx *sql.Tx, id string) (*domain.Payment, error)
	ReadAll(ctx context.Context, tx *sql.Tx) ([]domain.Payment, error)
	Update(ctx context.Context, tx *sql.Tx, payment *domain.Payment) error
	Delete(ctx context.Context, tx *sql.Tx, id string) error
}
