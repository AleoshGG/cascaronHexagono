package controllers

import (
	"fmt"
	"net/http"
	"practica/src/products/infrastructure/dependences"
	"practica/src/products/aplication"
	"practica/src/products/infrastructure"

	"github.com/gin-gonic/gin"
)

type GetAllProductsController struct{
	mysql *infrastructure.MySQL
	app *aplication.GetAllProducts
}

func NewGetAllProductsController() *GetAllProductsController {
	mysql := dependences.GetMySQL()
	app := aplication.NewGetAllProducts(mysql)
	return &GetAllProductsController{
		mysql: mysql,
		app:   app,
	}
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
	
	fmt.Print(res)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"links": gin.H{
			"self": "http://localhost:8080/products/",
		},
		"data": res,
	})
} 