package order

import (
	"context"
	"database/sql"

	"github.com/MCPutro/E-commerce/internal/domain"
)

type Repository interface {
	Write(ctx context.Context, tx *sql.Tx, order *domain.Order) error
	ReadByID(ctx context.Context, tx *sql.Tx, id string) (*domain.Order, error)
}
