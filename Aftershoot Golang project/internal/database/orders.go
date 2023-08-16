package database

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"sync"
	"time"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusProcessed OrderStatus = "processed"
	OrderStatusCanceled  OrderStatus = "canceled"
)

// GetOrderByID retrieves an order by its ID
func GetOrderByID(ctx context.Context, db *sql.DB, id int) (*Order, error) {
	var order Order
	query := `
        SELECT id, customer_id, product_id, quantity
        FROM orders
        WHERE id = $1
    `
	err := db.QueryRowContext(ctx, query, id).Scan(&order.ID, &order.CustomerID, &order.ProductID, &order.Quantity)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// ProcessOrder processes a new order and updates the inventory
func ProcessOrder(ctx context.Context, order *Order, db *sql.DB) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Fetch the product's current quantity
	var currentQuantity int32
	err = tx.QueryRowContext(ctx, "SELECT quantity FROM products WHERE id = $1", order.ProductID).Scan(&currentQuantity)
	if err != nil {
		return err
	}

	// Check if there's enough quantity to fulfill the order
	if currentQuantity < order.Quantity {
		e := errors.New("inventory error")
		return e
	}

	// Update the product's quantity
	_, err = tx.ExecContext(ctx, "UPDATE products SET quantity = quantity - $1 WHERE id = $2", order.Quantity, order.ProductID)
	if err != nil {
		return err
	}

	// Insert the order
	_, err = tx.ExecContext(ctx, "INSERT INTO orders (product_id, customer_id, quantity) VALUES ($1, $2, $3)", order.ProductID, order.CustomerID, order.Quantity)
	if err != nil {
		return err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// GetOrderStatus retrieves the status of an order by order ID
func GetOrderStatus(ctx context.Context, db *sql.DB, orderID int) (OrderStatus, error) {
	var status OrderStatus
	err := db.QueryRowContext(ctx, "SELECT status FROM orders WHERE id = ?", orderID).Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
		return "", err
	}
	return status, nil
}

// OrderWithDetails represents an order along with details about the products and customers involved
type OrderWithDetails struct {
	OrderID         int    `json:"order_id"`
	OrderDate       string `json:"order_date"`
	CustomerID      int    `json:"customer_id"`
	CustomerName    string `json:"customer_name"`
	CustomerEmail   string `json:"customer_email"`
	ProductID       int    `json:"product_id"`
	ProductName     string `json:"product_name"`
	ProductPrice    int    `json:"product_price"`
	ProductQuantity int    `json:"product_quantity"`
}

func GetOrdersWithDetails(db *sql.DB) ([]Order, error) {
	query := `
        SELECT id, product_id, customer_id, quantity
        FROM orders
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order

	for rows.Next() {
		var order Order
		err := rows.Scan(&order.ID, &order.ProductID, &order.CustomerID, &order.Quantity)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func PlaceOrders(orderRequests []Order, maxRetries int, retryDelay time.Duration, db *sql.DB) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errors []error

	ctx := context.Background()

	for _, req := range orderRequests {
		wg.Add(1)
		go func(req Order) {
			defer wg.Done()

			err := createOrderWithRetries(ctx, req, maxRetries, retryDelay, db)
			if err != nil {
				mu.Lock()
				errors = append(errors, err)
				mu.Unlock()
			}
		}(req)
	}

	wg.Wait()

	if len(errors) > 0 {
		return errors[0] // For simplicity, returning the first error
	}

	return nil
}

func createOrderWithRetries(ctx context.Context, req Order, maxRetries int, retryDelay time.Duration, db *sql.DB) error {
	for attempt := 1; attempt <= maxRetries; attempt++ {
		err := ProcessOrder(ctx, &req, db)
		if err == nil {
			return nil
		}

		// Handle inventory constraint error and retry after delay
		if IsInventoryConstraintError(err) {
			if attempt < maxRetries {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case <-time.After(retryDelay):
				}
			} else {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func IsInventoryConstraintError(err error) bool {
	if err != nil && strings.Contains(err.Error(), "inventory error") {
		return true
	}
	return false
}
