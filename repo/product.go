
package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgUrl      string  `json:"imageUrl" db:"img_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productId int) (*Product, error)
	List() ([]*Product, error)
	Delete(productId int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

// Create creates a new product in PostgreSQL.
func (r *productRepo) Create(p Product) (*Product, error) {
	query := `
		INSERT INTO products (
			title,
			description,
			price,
			img_url
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		p.Title,
		p.Description,
		p.Price,
		p.ImgUrl,
	).Scan(&p.ID)

	if err != nil {
		return nil, fmt.Errorf("create product: %w", err)
	}

	return &p, nil
}

// Get retrieves a product by ID.
func (r *productRepo) Get(productID int) (*Product, error) {
	var p Product

	query := `
		SELECT id, title, description, price, img_url
		FROM products
		WHERE id = $1
	`

	err := r.db.Get(&p, query, productID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("get product: %w", err)
	}

	return &p, nil
}

// List retrieves all products.
func (r *productRepo) List() ([]*Product, error) {
	products := make([]*Product, 0)

	query := `
		SELECT id, title, description, price, img_url
		FROM products
		ORDER BY id
	`

	err := r.db.Select(&products, query)
	if err != nil {
		return nil, fmt.Errorf("list products: %w", err)
	}

	return products, nil
}

// Update updates an existing product.
func (r *productRepo) Update(p Product) (*Product, error) {
	query := `
		UPDATE products
		SET title = $1,
			description = $2,
			price = $3,
			img_url = $4
		WHERE id = $5
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		p.Title,
		p.Description,
		p.Price,
		p.ImgUrl,
		p.ID,
	).Scan(&p.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("update product: %w", err)
	}

	return &p, nil
}

// Delete deletes a product by ID.
func (r *productRepo) Delete(productID int) error {
	query := `
		DELETE FROM products
		WHERE id = $1
	`

	result, err := r.db.Exec(query, productID)
	if err != nil {
		return fmt.Errorf("delete product: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("check deleted product: %w", err)
	}

	if affected == 0 {
		return fmt.Errorf("product with ID %d not found", productID)
	}

	return nil
}
