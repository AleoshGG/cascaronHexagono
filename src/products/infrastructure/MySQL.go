package infrastructure

import (
	"fmt"
	"practica/src/products/domain"

	"github.com/go-mysql-org/go-mysql/client"
)

type MySQL struct{
	conn client.Conn
}
//pasar la conexion
func NewMySQL(conn client.Conn) *MySQL {
	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save(product domain.Product) (uint64, error) {
	
	query := "INSERT INTO products (name, price) VALUES (?, ?)"
	res, err := mysql.conn.Execute(query, product.GetName(), product.GetPrice())

	if err != nil {
		return 0, err
	}
	
	return res.InsertId, nil
}

func (mysql *MySQL) GetAll() {
	fmt.Println("Lista de productos")
}
