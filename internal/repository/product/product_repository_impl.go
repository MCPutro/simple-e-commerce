package product

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/MCPutro/E-commerce/internal/domain"
	"github.com/MCPutro/E-commerce/pkg/constant"
	newError "github.com/MCPutro/E-commerce/pkg/error"
)

type ProductRepository interface {
	ReadAll(ctx context.Context, tx *sql.Tx) ([]*domain.Product, error)
	ReadByID(ctx context.Context, tx *sql.Tx, id uint) (*domain.Product, error)
	UpdateStock(ctx context.Context, tx *sql.Tx, productID uint, quantity int) error
	Write(ctx context.Context, tx *sql.Tx, product *domain.Product) error
	Update(ctx context.Context, tx *sql.Tx, product *domain.Product) error
	Delete(ctx context.Context, tx *sql.Tx, id uint) error
}

type productRepository struct{}

func (p *productRepository) ReadAll(ctx context.Context, tx *sql.Tx) ([]*domain.Product, error) {
	query := "SELECT id, name, price, stock, created_at, updated_at FROM e_commerce.products WHERE deleted_at IS NULL FOR UPDATE;"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		var product domain.Product
		var createdAt, updatedAt sql.NullTime
		if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &createdAt, &updatedAt); err != nil {
			return nil, err
		}

		if createdAt.Valid {
			product.CreationTime = createdAt.Time
		}
		if updatedAt.Valid {
			product.UpdateTime = updatedAt.Time
		}

		products = append(products, &product)
	}

	return products, nil
}

func (p *productRepository) ReadByID(ctx context.Context, tx *sql.Tx, id uint) (*domain.Product, error) {
	query := "SELECT id, name, price, stock, created_at, updated_at FROM e_commerce.products WHERE deleted_at IS NULL AND id = ? LIMIT 1 FOR UPDATE;"
	row := tx.QueryRowContext(ctx, query, id)

	var product domain.Product
	var createdAt, updatedAt sql.NullTime
	if err := row.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &createdAt, &updatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, newError.ErrProductNotFound
		}
		return nil, err
	}

	if createdAt.Valid {
		product.CreationTime = createdAt.Time
	}
	if updatedAt.Valid {
		product.UpdateTime = updatedAt.Time
	}

	return &product, nil
}

func (p *productRepository) UpdateStock(ctx context.Context, tx *sql.Tx, productID uint, quantity int) error {
	query := "UPDATE e_commerce.products SET stock = ?, updated_at = NOW() WHERE id = ?;"
	result, err := tx.ExecContext(ctx, query, quantity, productID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows updated")
	}
	return nil
}

func (p *productRepository) Write(ctx context.Context, tx *sql.Tx, product *domain.Product) error {
	now := time.Now().Format(constant.TimeFormat)
	query := "INSERT INTO e_commerce.products (name, price, stock, created_at, updated_at) VALUES (?, ?, ?, ?, ?);"

	result, err := tx.ExecContext(ctx, query, product.Name, product.Price, product.Stock, now, now)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	product.Id = uint(id)
	return nil
}

// Delete implements Repository.
func (p *productRepository) Delete(ctx context.Context, tx *sql.Tx, id string) error {
	query := "UPDATE e_commerce.products SET deleted_at = NOW() WHERE id = ?;"
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

// Update implements Repository.
func (p *productRepository) Update(ctx context.Context, tx *sql.Tx, product *domain.Product) error {
	query := "UPDATE e_commerce.products SET name = ?, price = ?, stock = ?, updated_at = NOW() WHERE id = ?;"
	result, err := tx.ExecContext(ctx, query, product.Name, product.Price, product.Stock, product.Id)
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

func NewProductRepository() Repository {
	return &productRepository{}
}
