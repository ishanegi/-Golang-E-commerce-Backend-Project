package main

import (
	"fmt"
	"go-app/internal/api"
	"go-app/internal/app"
	"go-app/internal/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin instance
	r := gin.Default()

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	app, err := app.NewApp()
	if err != nil {
		return
	}
	r.Use(app.DBMiddleware)
	app.ConfigureMiddleware()
	api.SetupRoutes(r, db.DB) // dekhte h

	// Start the server on a specific port
	serverPort := ":8080"
	// Run the server
	fmt.Println("Server is running on", serverPort)
	err = r.Run(serverPort)
	if err != nil {
		return
	}

	port := 8080
	log.Printf("Server started on :%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
