package db

import (
	"fmt"
	"log"
	"os"

	"github.com/go-mysql-org/go-mysql/client"
)

var dbConn *client.Conn

// Función para inicializar la conexión
func InitializeDatabase() {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")

	// Realizar la conexión directa a la base de datos
	conn, err := client.Connect(host, user, password, database)
	if err != nil {
		log.Fatal("Error al conectar con la base de datos: ", err)
	}

	dbConn = conn

	// Imprimir un mensaje si la conexión fue exitosa
	fmt.Println("Conexión a la base de datos establecida exitosamente")
}

// Obtener la conexión a la base de datos
func GetDBConnection() *client.Conn {
	return dbConn
}

