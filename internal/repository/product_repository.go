package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/MCPutro/E-commerce/internal/entity"
	"github.com/MCPutro/E-commerce/pkg/constant"
	newError "github.com/MCPutro/E-commerce/pkg/error"
)

type ProductRepository interface {
	ReadAll(ctx context.Context, tx *sql.Tx) ([]*entity.Product, error)
	ReadByID(ctx context.Context, tx *sql.Tx, id uint) (*entity.Product, error)
	UpdateStock(ctx context.Context, tx *sql.Tx, productID uint, quantity int) error
	Write(ctx context.Context, tx *sql.Tx, product *entity.Product) error
}

type productRepository struct{}

func (p *productRepository) ReadAll(ctx context.Context, tx *sql.Tx) ([]*entity.Product, error) {
	query := "SELECT id, name, price, stock, creation_time, update_time, delete_time FROM e_commerce.products where delete_time is null FOR UPDATE;"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		var createdAt, updateTime, deleteTime sql.NullTime
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &createdAt, &updateTime, &deleteTime)
		if err != nil {
			return nil, err
		}

		if createdAt.Valid {
			product.CreationTime = createdAt.Time
		}

		if updateTime.Valid {
			product.UpdateTime = updateTime.Time
		}

		products = append(products, &product)
	}

	return products, nil
}

func (p *productRepository) ReadByID(ctx context.Context, tx *sql.Tx, id uint) (*entity.Product, error) {
	query := "SELECT id, name, price, stock, creation_time, update_time, delete_time FROM e_commerce.products where delete_time is null and id = ? LIMIT 1 FOR UPDATE;"
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var product entity.Product
		var createdAt, updateTime, deleteTime sql.NullTime
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &createdAt, &updateTime, &deleteTime)
		if err != nil {
			return nil, err
		}

		if createdAt.Valid {
			product.CreationTime = createdAt.Time
		}

		if updateTime.Valid {
			product.UpdateTime = updateTime.Time
		}

		return &product, nil
	}

	return nil, newError.ErrProductNotFound
}

func (p *productRepository) UpdateStock(ctx context.Context, tx *sql.Tx, productID uint, quantity int) error {

	query := "UPDATE e_commerce.products SET update_time = NOW(), stock = ? WHERE id = ?;"
	result, err := tx.ExecContext(ctx, query, quantity, productID)
	if err != nil {
		return err
	}

	res, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if res == 0 {
		return errors.New("no rows updated")
	} else {
		return nil
	}
}

func (p *productRepository) Write(ctx context.Context, tx *sql.Tx, product *entity.Product) error {
	now := time.Now().Format(constant.TimeFormat)

	query := "INSERT INTO e_commerce.products (name,price,stock,creation_time,update_time) VALUES (?,?,?,?,?);"

	result, err := tx.ExecContext(ctx, query,
		product.Name,
		product.Price,
		product.Stock,
		now,
		now,
	)
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

func NewProductRepository() ProductRepository {
	return &productRepository{}
}
