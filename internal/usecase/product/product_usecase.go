package product

// import (
// 	"context"
// 	"database/sql"
// 	"errors"
// 	"fmt"

// 	"github.com/MCPutro/E-commerce/internal/domain"
// 	"github.com/MCPutro/E-commerce/internal/repository/product"
// 	newError "github.com/MCPutro/E-commerce/pkg/error"
// )

// type ProductUseCase interface {
// 	UpdateProduct(ctx context.Context, product *domain.Product) error
// 	GetProductByID(ctx context.Context, productID uint) (*domain.Product, error)
// 	GetProducts(ctx context.Context) ([]*domain.Product, error)
// 	CreateProduct(ctx context.Context, product *domain.Product) error
// }

// type productUsecase struct {
// 	productRepo product.Repository
// 	db          *sql.DB
// }

// func (p *productUsecase) UpdateProduct(ctx context.Context, product *domain.Product) error {
// 	if product.Stock < 0 {
// 		return errors.New("stock quantity cannot be negative")
// 	}

// 	tx, err := p.db.Begin()
// 	if err != nil {
// 		return newError.ErrOpenTransactionWithDetails(err.Error())
// 	}
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	// Cek apakah produk ada
// 	existingProduct, err := p.productRepo.ReadByID(ctx, tx, product.Id)
// 	if err != nil {
// 		return fmt.Errorf("produk tidak ditemukan: %w", err)
// 	}

// 	if existingProduct == nil {
// 		return newError.ErrProductNotFound
// 	}

// 	// Update stok
// 	err = p.productRepo.UpdateStock(ctx, tx, product.Id, product.Stock)
// 	if err != nil {
// 		return fmt.Errorf("gagal mengupdate stok: %w", err)
// 	}

// 	// Update informasi produk lainnya jika diperlukan
// 	// ...

// 	return tx.Commit()
// }

// func (p *productUsecase) GetProductByID(ctx context.Context, productID uint) (*domain.Product, error) {
// 	tx, err := p.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return nil, newError.ErrOpenTransactionWithDetails(err.Error())
// 	}
// 	defer tx.Rollback()

// 	product, err := p.productRepo.ReadByID(ctx, tx, productID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return product, tx.Commit()
// }

// func (p *productUsecase) GetProducts(ctx context.Context) ([]*domain.Product, error) {
// 	tx, err := p.db.Begin()
// 	if err != nil {
// 		return nil, newError.ErrOpenTransactionWithDetails(err.Error())
// 	}
// 	defer tx.Rollback()

// 	products, err := p.productRepo.ReadAll(ctx, tx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	tx.Commit()
// 	return products, nil
// }

// func (p *productUsecase) CreateProduct(ctx context.Context, product *domain.Product) error {

// 	tx, err := p.db.Begin()
// 	if err != nil {
// 		return newError.ErrOpenTransactionWithDetails(err.Error())
// 	}
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	err = p.productRepo.Write(ctx, tx, product)
// 	if err != nil {
// 		return fmt.Errorf("gagal membuat produk: %w", err)
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return fmt.Errorf("gagal menyimpan produk: %w", err)
// 	}

// 	return nil
// }

// func NewProductUseCase(productRepo product.Repository, db *sql.DB) ProductUseCase {
// 	return &productUsecase{
// 		productRepo: productRepo,
// 		db:          db,
// 	}
// }
