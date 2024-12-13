package user

import (
	"context"
	"database/sql"

	"github.com/MCPutro/E-commerce/internal/domain"
)

type Repository interface {
	Write(cxt context.Context, tx *sql.Tx, user *domain.User) error
	ReadById(cxt context.Context, tx *sql.Tx, id string) (*domain.User, error)
	ReadByEmail(cxt context.Context, tx *sql.Tx, email string) (*domain.User, error)
	ReadAll(cxt context.Context, tx *sql.Tx) ([]domain.User, error)
	Update(cxt context.Context, tx *sql.Tx, user *domain.User) error
	Delete(cxt context.Context, tx *sql.Tx, id string) error
}
