package api

import (
	"net/http"

	"go-app/internal/database"

	"github.com/gin-gonic/gin"
)

func AddCustomerHandler(c *gin.Context) {
	var customer database.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("DB").(*database.DB)
	err := db.AddCustomer(c, &customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add customer"})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusCreated, "Hello")
}

// GetTopCustomersHandler retrieves the top three customers who have placed the most orders
func GetTopCustomersHandler(c *gin.Context) {
	db := c.MustGet("DB").(*database.DB)
	customers, err := database.GetTopCustomers(c, db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch top customers"})
		return
	}
	c.JSON(http.StatusOK, customers)
}
