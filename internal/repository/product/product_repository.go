package product

import (
	"context"
	"database/sql"

	"github.com/MCPutro/E-commerce/internal/domain"
)

type Repository interface {
	ReadAll(ctx context.Context, tx *sql.Tx) ([]*domain.Product, error)
	ReadByID(ctx context.Context, tx *sql.Tx, id uint) (*domain.Product, error)
	UpdateStock(ctx context.Context, tx *sql.Tx, productID uint, quantity int) error
	Write(ctx context.Context, tx *sql.Tx, product *domain.Product) error
	Update(ctx context.Context, tx *sql.Tx, product *domain.Product) error
	Delete(ctx context.Context, tx *sql.Tx, id string) error
}
