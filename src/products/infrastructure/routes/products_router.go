package routes

import (
	"practica/src/products/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	productsRoutes := r.Group("/products")
	{
		productsRoutes.POST("/add", controllers.AddProduct)
	}
}