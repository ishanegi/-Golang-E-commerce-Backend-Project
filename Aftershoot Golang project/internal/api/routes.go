package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {

	r.GET("/", Hello)
	// Customers
	r.POST("/customers", AddCustomerHandler) // test done

	// Orders
	r.POST("/orders", ProcessOrderHandler) // test done

	// orders with retry mechanism
	r.POST("/retry-orders", PlaceOrdersHandler) // test done

	// Products
	r.POST("/products", AddProductHandler) // test done

	// Get Products
	r.GET("/products", GetProducts) // test done

	// Create Product Rating
	r.POST("/product-rating", CreateProductRating) // test done

	r.GET("/products-ratings", GetProductsAndRatingsHandler) // test done
	r.GET("/top-customers", GetTopCustomersHandler)          // test done
	r.GET("/orders-details", GetOrdersWithDetailsHandler)    // test done
}
