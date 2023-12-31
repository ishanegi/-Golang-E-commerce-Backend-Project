// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package database

import (
	"database/sql"
)

type Customer struct {
	ID    int32
	Name  string
	Email string
}

type Order struct {
	ID          int32
	ProductID   int32
	CustomerID  int32
	Quantity    int32
	Price       int32
	Orderstatus sql.NullString
	OrderDate   sql.NullTime
}

type Product struct {
	ID       int32
	Name     string
	Price    string
	Quantity int32
}

type Rating struct {
	ID        int32
	ProductID int32
	Rating    string
}

type User struct {
	ID       int32
	Username string
	Password string
	Role     string
}
