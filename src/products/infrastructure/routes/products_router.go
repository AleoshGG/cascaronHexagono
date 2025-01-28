package routes

import (
	"practica/src/products/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	productsRoutes := r.Group("/products")
	{
		productsRoutes.POST("/add", controllers.NewCreateProductController().AddProduct)
		productsRoutes.GET("/", controllers.NewGetAllProductsController().GetAllProducts)
		productsRoutes.PUT("/:id", controllers.NewUpdateProductController().UpdateProduct)
		productsRoutes.DELETE("/:id", controllers.NewDeleteProductController().DeleteProduct)
	}
}