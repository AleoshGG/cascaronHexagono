package main

import (
	"practica/src/products/infrastructure/routes"
	"practica/src/core"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // Cargar las variables de entorno

	// Inicializar la conexión al iniciar el servidor
	db.InitializeDatabase()

	// Configurar Gin
	r := gin.Default()

	// Registrar rutas
	routes.RegisterRoutes(r)

	// Iniciar el servidor
	r.Run(":8080")
}