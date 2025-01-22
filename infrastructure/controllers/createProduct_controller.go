package controllers

import (
	"net/http"
	"practica/aplication"
	"practica/domain"
	"practica/infrastructure"

	"github.com/gin-gonic/gin"
)

// Crear un nuevo recurso
func AddProduct(c *gin.Context) {
	
	// Recuperacion del body 
	var newProduct struct {
		Name string `json: name`
		Price float32 `json: price`
	}
	c.ShouldBindJSON(&newProduct)

	var product = domain.NewProduct(newProduct.Name, newProduct.Price)

	// Instancia de mysql
	var mysql = infrastructure.NewMySQL()

	// Inyeccion
	aplication.NewCreateProduct(mysql).Run(*product)



	c.JSON(http.StatusCreated, "Hola jeje")
}