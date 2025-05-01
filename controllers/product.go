package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jayeshjadhav/mobile-store/config"
	"github.com/jayeshjadhav/mobile-store/models"
)

func CreateProduct(c *gin.Context) {
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := c.GetUint("userID")
	input.CreatedByID = userID

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func GetAllProducts(c *gin.Context) {
	adminName := c.Query("admin")

	var products []models.Product
	if adminName != "" {
		var user models.User
		if err := config.DB.Where("name = ?", adminName).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
			return
		}
		config.DB.Where("created_by_id = ?", user.ID).Find(&products)
	} else {
		config.DB.Find(&products)
	}
	c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Ensure only creator can edit
	userID := c.GetUint("userID")
	if product.CreatedByID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only edit your own products"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	userID := c.GetUint("userID")
	if product.CreatedByID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own products"})
		return
	}

	config.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
