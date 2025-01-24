package main

import (
	"practica/dependences"
	"practica/src/products/infrastructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	dependences.GoDependences()

	// Configurar Gin
	r := gin.Default()

	// Registrar rutas
	routes.RegisterRoutes(r)

	// Iniciar el servidor
	r.Run(":8080")
}