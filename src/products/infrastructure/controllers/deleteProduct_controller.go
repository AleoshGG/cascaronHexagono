package controllers

import (
	"net/http"
	"practica/src/products/aplication"
	"practica/src/products/infrastructure/dependences"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	app *aplication.DeleteProduct
}

func NewDeleteProductController() *DeleteProductController {
	mysql := dependences.GetMySQL()
	app := aplication.NewDeleteProduct(mysql)

	return &DeleteProductController{app: app}
}

func (dp_c *DeleteProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	id_num, _ := strconv.ParseInt(id, 10, 64)

	rowsAffected, err := dp_c.app.Run(int(id_num))

	if err != nil || rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error": "No se pudo eliminar el producto " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Recurso eliminado",
	})
}