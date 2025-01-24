package controllers

import (
	"net/http"
	"practica/dependences"
	"practica/src/products/domain"

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

	// Inyeccion mover
	id, err := dependences.GetApp().Run(*product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No se pudo guardar el producto " + err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"msg": "Creado",
		"id": id,
	})
}

// Martin Fouler - Patrones de diseño para las arquitecturas