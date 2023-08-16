// api/api.go

package api

import (
	"go-app/internal/database"
)

type API struct {
	DB *database.DB
}

func NewAPI(db *database.DB) *API {
	return &API{DB: db}
}
