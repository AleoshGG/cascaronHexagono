package routes

import (
	"practica/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	productsRoutes := r.Group("/products")
	{
		productsRoutes.POST("/add", controllers.AddProduct)
	}
}