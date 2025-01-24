package controllers

import (
	"fmt"
	"net/http"
	"practica/dependences"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	res, err := dependences.GetAppGetAll().Run()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error": "No se pudo recuperar los productos " + err.Error(),
		})
		return
	}

	fmt.Print(res)

	c.JSON(http.StatusOK, "Bien")
}