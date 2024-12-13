package cart

import (
	"context"
	"database/sql"

	"github.com/MCPutro/E-commerce/internal/domain"
)

type Repository interface {
	Write(ctx context.Context, tx *sql.Tx, cart *domain.Cart) (*domain.Cart, error)
	ReadByUserId(ctx context.Context, tx *sql.Tx, userId string) (*domain.Cart, error)
	Remove(ctx context.Context, tx *sql.Tx, cardId string) error
	AddItem(ctx context.Context, tx *sql.Tx, cartId string, item *domain.CartItem) error
	RemoveItem(ctx context.Context, tx *sql.Tx, cartId string, productId string) error
	UpdateItem(ctx context.Context, tx *sql.Tx, cartId string, item *domain.CartItem) error
}
