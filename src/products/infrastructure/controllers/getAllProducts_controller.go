package controllers

import (
	"net/http"
	"practica/src/products/aplication"
	"practica/src/products/infrastructure/dependences"

	"github.com/gin-gonic/gin"
)

type GetAllProductsController struct{
	app *aplication.GetAllProducts
}

func NewGetAllProductsController() *GetAllProductsController {
	mysql := dependences.GetMySQL()
	app := aplication.NewGetAllProducts(mysql)
	return &GetAllProductsController{app: app}
}

func (gp_c *GetAllProductsController) GetAllProducts(c *gin.Context) {
	res, err := gp_c.app.Run()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error": "No se pudo recuperar los productos " + err.Error(),
		})
		return
	}

	type product struct {
		Id int64 `json: "id"`
		Name string `json: "name"`
		Price float64 `json: "price"`
	}

	var result []product

	for _, row := range res {
		var product product 
		product.Id = row.GetId()
		product.Name = row.GetName()
		product.Price = row.GetPrice()
		result = append(result, product)
	}


	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"links": gin.H{
			"self": "http://localhost:8080/products/",
		},
		"data": result,
	})
} 