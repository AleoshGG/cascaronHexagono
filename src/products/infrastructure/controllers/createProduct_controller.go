package controllers

import (
	"net/http"
	"practica/dependences"
	"practica/src/products/aplication"
	"practica/src/products/domain"
	"practica/src/products/infrastructure"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct{
	mysql *infrastructure.MySQL
	app *aplication.CreateProduct
}

func NewCreateProductController() *CreateProductController {
	mysql := dependences.GetMySQL()
	app := aplication.NewCreateProduct(mysql)
	return &CreateProductController{
		mysql: mysql,
		app:   app,
	}
}

// Crear un nuevo recurso
func (cp_c *CreateProductController) AddProduct(c *gin.Context) {
	
	// Recuperacion del body 
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

	var product = domain.NewProduct(0, newProduct.Name, newProduct.Price)

	// Inyeccion mover
	id, err :=  cp_c.app.Run(*product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error": "No se pudo guardar el producto " + err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data": gin.H{
			"type": "products",
			"id": id,
			"attributes": gin.H{
			  "name": product.GetName(),
			  "price": product.GetPrice(),
			},
		},
	})
}

// Martin Fouler - Patrones de diseño para las arquitecturas