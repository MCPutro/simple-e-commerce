package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/MCPutro/E-commerce/internal/domain"
	"github.com/MCPutro/E-commerce/internal/repository/user"
	newError "github.com/MCPutro/E-commerce/pkg/error"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo user.Repository
	db       *sql.DB
}

func NewUserUseCase(userRepo user.Repository, db *sql.DB) UseCase {
	return &userUseCase{
		userRepo: userRepo,
		db:       db,
	}
}

func (uc *userUseCase) Registration(ctx context.Context, user *domain.User) error {
	if user == nil {
		return newError.ErrInvalidInput
	}

	// begin database trx
	trx, err := uc.db.Begin()
	if err != nil {
		return newError.ErrOpenTransactionWithDetails(err.Error())
	}
	defer func() {
		if err != nil {
			trx.Rollback()
		}
	}()

	// encrypt password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// update plain password to hash password
	user.Password = string(hashPassword)

	// write to repo
	err = uc.userRepo.Write(ctx, trx, user)
	if err != nil {
		return err
	}

	if err = trx.Commit(); err != nil {
		return newError.ErrCommitTransactionWithDetails(err.Error())
	}

	return nil
}

func (uc *userUseCase) Login(ctx context.Context, email string, password string) (*domain.User, error) {
	// begin database trx
	trx, err := uc.db.Begin()
	if err != nil {
		return nil, newError.ErrOpenTransactionWithDetails(err.Error())
	}
	defer func() {
		if err != nil {
			trx.Rollback()
		}
	}()

	// get user by email
	user, err := uc.userRepo.ReadByEmail(ctx, trx, email)
	if err != nil {
		return nil, err
	}

	// check input password and saved password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("email dan password tidak sesuai")
	}

	if err = trx.Commit(); err != nil {
		return nil, newError.ErrCommitTransactionWithDetails(err.Error())
	}

	return user, nil
}

func (uc *userUseCase) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	// begin database trx
	trx, err := uc.db.Begin()
	if err != nil {
		return nil, err
	}
	defer trx.Rollback()

	users, err := uc.userRepo.ReadAll(ctx, trx)
	if err != nil {
		return nil, err
	}

	trx.Commit()

	return users, nil
}

func (uc *userUseCase) UpdateUser(ctx context.Context, user *domain.User) error {
	if user == nil {
		return newError.ErrInvalidInput
	}

	// begin database trx
	trx, err := uc.db.Begin()
	if err != nil {
		return err
	}
	defer trx.Rollback()

	err = uc.userRepo.Update(ctx, trx, user)

	trx.Commit()

	return err
}

func (uc *userUseCase) DeleteUser(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("invalid user id")
	}

	// begin database trx
	trx, err := uc.db.Begin()
	if err != nil {
		return err
	}
	defer trx.Rollback()

	err = uc.userRepo.Delete(ctx, trx, id)
	trx.Commit()

	return err
}
