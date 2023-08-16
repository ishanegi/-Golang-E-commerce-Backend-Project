package database

import (
	"context"
	"database/sql"
	_ "database/sql"
)

type CustomerWithOrderCount struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	OrderCount int    `json:"order_count"`
}

// GetCustomerByID retrieves a customer by their ID
func (db *DB) GetCustomerByID(ctx context.Context, id int) (*Customer, error) {
	var customer Customer
	query := `
        SELECT id, name, email
        FROM customers
        WHERE id = $1
    `
	err := db.QueryRowContext(ctx, query, id).Scan(&customer.ID, &customer.Name, &customer.Email)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

// AddCustomer adds a new customer to the database
func (db *DB) AddCustomer(ctx context.Context, customer *Customer) error {
	_, err := db.ExecContext(ctx, `
		INSERT INTO customers (name, email)
		VALUES ($1, $2)
	`, customer.Name, customer.Email)
	if err != nil {
		return err
	}

	return nil
}

// GetTopCustomers fetches the top three customers who have placed the most orders
func GetTopCustomers(ctx context.Context, db *sql.DB) ([]CustomerWithOrderCount, error) {
	query := `
        SELECT
            c.id,
            c.name,
            c.email,
            COUNT(o.id) AS order_count
        FROM
            customers c
        JOIN
            orders o ON c.id = o.customer_id
        GROUP BY
            c.id, c.name, c.email
        ORDER BY
            order_count DESC
        LIMIT
            3;
    `

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []CustomerWithOrderCount
	for rows.Next() {
		var customer CustomerWithOrderCount
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.OrderCount); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}
