package user

import (
	"context"

	"github.com/MCPutro/E-commerce/internal/domain"
)

type UseCase interface {
	Registration(ctx context.Context, user *domain.User) error
	Login(ctx context.Context, email string, password string) (*domain.User, error)
	GetAllUsers(ctx context.Context) ([]domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id string) error
}
