package controllers

import (
	"net/http"
	"practica/src/products/aplication"
	"practica/src/products/domain"
	"practica/src/products/infrastructure"
	"practica/src/products/infrastructure/dependences"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateProductController struct {
	mysql *infrastructure.MySQL
	app *aplication.UpdateProduct
}


func NewUpdateProductController() *UpdateProductController {
	mysql := dependences.GetMySQL()
	app := aplication.NewUpdateProduct(mysql)

	return &UpdateProductController{mysql: mysql, app: app}
}

func (up_c *UpdateProductController) UpdateProduct (c *gin.Context) {

	id := c.Param("id")
	id_num, _ := strconv.ParseInt(id, 10, 64)
	 
	var newProduct struct {
		Name string `json: name`
		Price float64 `json: price`
	}

	if err := c.ShouldBindJSON(&newProduct); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
            "error": "Datos inválidos: " + err.Error(),
        })
        return
    }

	var product = domain.NewProduct(id_num, newProduct.Name, newProduct.Price)

	rowsAffected, err := up_c.app.Run(*product)

	if err != nil || rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error": "No se pudo actualizar el producto " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data": gin.H{
			"type": "products",
			"id": id_num,
			"attributes": gin.H{
			  "name": product.GetName(),
			  "price": product.GetPrice(),
			},
		},
	})
}