package api

import (
	"net/http"
	"time"

	"go-app/internal/database"

	"github.com/gin-gonic/gin"
)

func ProcessOrderHandler(c *gin.Context) {
	var order database.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("DB").(*database.DB)
	err := database.ProcessOrder(c, &order, db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process order"})
		return
	}

	c.JSON(http.StatusCreated, order)

}

// GetOrdersWithDetailsHandler retrieves orders along with details about the products and customers involved
func GetOrdersWithDetailsHandler(c *gin.Context) {
	db := c.MustGet("DB").(*database.DB)
	orders, err := database.GetOrdersWithDetails(db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders with details"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func PlaceOrdersHandler(c *gin.Context) {
	db := c.MustGet("DB").(*database.DB)
	var orderRequests []database.Order
	if err := c.ShouldBindJSON(&orderRequests); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	maxRetries := 2               // Maximum number of retries for order processing
	retryDelay := 1 * time.Second // Delay between retries

	err := database.PlaceOrders(orderRequests, maxRetries, retryDelay, db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to place orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Orders placed successfully"})
}
