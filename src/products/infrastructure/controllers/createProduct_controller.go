package controllers

import (
	"net/http"
	"practica/src/products/aplication"
	"practica/src/products/domain"
	"practica/src/products/infrastructure"

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

	// Instancia de mysql mover
	var mysql = infrastructure.NewMySQL()

	// Inyeccion mover
	aplication.NewCreateProduct(mysql).Run(*product)

	c.JSON(http.StatusCreated, "Hola jeje")
}

// Martin Fouler - Patrones de diseño para las arquitecturas