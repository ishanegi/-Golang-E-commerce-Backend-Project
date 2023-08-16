package database

import (
	"context"
	"database/sql"
)

// GetProductByID retrieves a product by its ID
func GetProductByID(ctx context.Context, db *sql.DB, id int) (*Product, error) {
	var product Product
	query := `
        SELECT id, name, price
        FROM products
        WHERE id = $1
    `
	err := db.QueryRowContext(ctx, query, id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// AddProduct adds a new product to the database
func (db *DB) AddProduct(ctx context.Context, product *Product) error {
	_, err := db.ExecContext(ctx, `
        INSERT INTO products (name, price, quantity)
        VALUES ($1, $2, $3)
    `, product.Name, product.Price, product.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) CreateProductRating(rating Rating) error {
	query := `
        INSERT INTO ratings (product_id, rating)
        VALUES ($1, $2)
    `

	_, err := db.Exec(query, rating.ProductID, rating.Rating)
	if err != nil {
		return err
	}

	return nil
}

// ProductWithAvgRating represents a product along with its average rating
type ProductWithAvgRating struct {
	ID        int     `db:"id"`
	Name      string  `db:"name"`
	Price     float64 `db:"price"`
	Quantity  int     `db:"quantity"`
	AvgRating float64 `db:"avg_rating"`
}

// GetProductsWithAvgRatings fetches a list of products along with their average ratings
func GetProductsWithAvgRatings(ctx context.Context, db *sql.DB) ([]ProductWithAvgRating, error) {
	var products []ProductWithAvgRating
	query := `
        SELECT
            p.id,
            p.name,
            p.price,
            p.quantity,
            AVG(r.rating) AS avg_rating
        FROM products p
        LEFT JOIN ratings r ON p.id = r.product_id
        GROUP BY p.id, p.name, p.price, p.quantity
        ORDER BY avg_rating DESC
    `

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product ProductWithAvgRating
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity, &product.AvgRating); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (db *DB) GetProducts() ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price, quantity FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Quantity)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
