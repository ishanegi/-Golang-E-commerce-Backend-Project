package app

import (
	"fmt"
	"go-app/internal/api"
	"go-app/internal/database"
	"go-app/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	DB     *database.DB
	API    *api.API
}

func NewApp() (*App, error) {
	r := gin.Default()

	db, err := database.ConnectDB()
	if err != nil {
		return nil, err
	}
	api := api.NewAPI(db)

	app := &App{
		Router: r,
		DB:     db,
		API:    api,
	}
	r.Use(func(c *gin.Context) {
		c.Set("db", db) // Set the database connection in the context
		c.Next()
	})
	app.configureMiddleware()
	app.configureRoutes()

	return app, nil
}

func (a *App) DBMiddleware(c *gin.Context) {
	c.Set("DB", a.DB) // Add this line to set the DB value in the context
	fmt.Println("Setting DB in context")
	c.Next()
}

func (a *App) ConfigureMiddleware() {
	a.Router.Use(gin.Logger())
	// a.Router.Use(middlewares.CORSMiddleware())
	a.Router.Use(a.DBMiddleware)
	// Add other middleware
}

func (a *App) Run(addr string) {
	a.Router.Run(addr)
}

func (a *App) configureRoutes() {
	api.SetupRoutes(a.Router, *&a.DB.DB)
}

func (a *App) configureMiddleware() {
	a.Router.Use(gin.Logger())
	a.Router.Use(middlewares.CORSMiddleware())
	a.Router.Use(middlewares.AuthMiddleware())
}
