package dependences

import (
	db "practica/src/core"
	"practica/src/products/aplication"
	"practica/src/products/infrastructure"

	"github.com/joho/godotenv"
)

var app *aplication.CreateProduct

func GoDependences() {
	godotenv.Load() // Cargar las variables de entorno

	// Inicializar la conexión al iniciar el servidor
	db.InitializeDatabase()

	// Instancia de mysql mover
	conn := db.GetDBConnection()
	app = aplication.NewCreateProduct(infrastructure.NewMySQL(*conn))
}

func GetApp() *aplication.CreateProduct {
	return app
}

