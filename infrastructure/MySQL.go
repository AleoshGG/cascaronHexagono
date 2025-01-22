package infrastructure

import (
	"fmt"
	"log"
	"practica/db"
	"practica/domain"
)

type MySQL struct{}

func NewMySQL() *MySQL {
	return &MySQL{}
}

func (mysql *MySQL) Save(product domain.Product) {
	
	
	conn := db.GetDBConnection()
	query := "INSERT INTO products (id, name, price) VALUES (?, ?, ?)"
	_, err := conn.Execute(query, product.GetId(), product.GetName(), product.GetPrice())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Producto salvado" + product.ViewProduct())
}

func (mysql *MySQL) GetAll() {
	fmt.Println("Lista de productos")
}
