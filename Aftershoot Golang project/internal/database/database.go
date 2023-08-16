package database

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func ConnectDB() (*DB, error) {
	db, err := sql.Open("postgres", "user=postgres password=1234 dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	return &DB{db}, nil
}

// CheckAdminRole checks if a user has admin role
func CheckAdminRole(userID int) (bool, error) {
	db, err := ConnectDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	var isAdmin bool
	err = db.QueryRowContext(context.Background(), "SELECT is_admin FROM users WHERE id = $1", userID).Scan(&isAdmin)
	if err != nil {
		return false, err
	}

	return isAdmin, nil
}
