package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jayeshjadhav/mobile-store/controllers"
	"github.com/jayeshjadhav/mobile-store/middleware"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		api.GET("/products", controllers.GetAllProducts)
		api.GET("/products/:id", controllers.GetProductByID)

		auth := api.Group("/")
		auth.Use(middleware.AuthMiddleware(), middleware.AdminOnly())
		{
			auth.POST("/products", controllers.CreateProduct)
			auth.PUT("/products/:id", controllers.UpdateProduct)
			auth.DELETE("/products/:id", controllers.DeleteProduct)
		}
	}
}
