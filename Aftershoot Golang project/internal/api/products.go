package api

import (
	"net/http"

	"go-app/internal/database"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	db := c.MustGet("DB").(*database.DB)
	products, err := db.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products and ratings"})
		return
	}
	c.JSON(http.StatusOK, products)

}

func CreateProductRating(c *gin.Context) {
	var rating database.Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("DB").(*database.DB)

	err := db.CreateProductRating(rating)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product rating"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product Rating added successfully"})
}

func AddProductHandler(c *gin.Context) {
	var product database.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("DB").(*database.DB)

	err := db.AddProduct(c, &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added successfully"})
}

// GetProductsAndRatingsHandler retrieves a list of products along with their average ratings
func GetProductsAndRatingsHandler(c *gin.Context) {
	db := c.MustGet("DB").(*database.DB)
	products, err := database.GetProductsWithAvgRatings(c, db.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products and ratings"})
		return
	}
	c.JSON(http.StatusOK, products)
}
