package dependences

import (
	db "practica/src/core"
	"practica/src/products/aplication"
	"practica/src/products/infrastructure"

	"github.com/joho/godotenv"
)

var appCreate *aplication.CreateProduct
var appGetAll *aplication.GetAllProducts

func GoDependences() {
	godotenv.Load() // Cargar las variables de entorno

	// Inicializar la conexión al iniciar el servidor
	db.InitializeDatabase()

	// Instancia de mysql mover
	conn := db.GetDBConnection()
	mysql := infrastructure.NewMySQL(*conn)
	appCreate = aplication.NewCreateProduct(mysql)
	appGetAll = aplication.NewGetAllProducts(mysql)
}

func GetAppCreate() *aplication.CreateProduct {
	return appCreate
}

func GetAppGetAll() *aplication.GetAllProducts {
	return appGetAll
}

