package infrastructure

import (
	db "practica/src/core"
	"practica/src/products/domain"

	"github.com/go-mysql-org/go-mysql/client"
)

type MySQL struct{
	conn *client.Conn
}

func NewMySQL() *MySQL {
	conn := db.GetDBConnection()
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

func (mysql *MySQL) GetAll() ([]domain.Product, error) {

	query := "SELECT id, name, price FROM products"
	res, err := mysql.conn.Execute(query)

	if err != nil {
		return nil, err
	}

	var products []domain.Product

	// Itera sobre las filas del resultado
	for _, row := range res.Values {
		
		// Leer valores de las columnas
		id := row[0].AsInt64()      // Convertir a int64
		name := row[1].String()   // Convertir a string
		price := row[2].AsFloat64() // Convertir a float64

		product := domain.NewProduct(id, name, price)

		products = append(products, *product)
	}

	return products, nil
}

func (mysql *MySQL) Update(product domain.Product) (uint64, error) {

	query := "UPDATE products SET name = ?, price = ? WHERE id = ?"
	res, err := mysql.conn.Execute(query, product.GetName(), product.GetPrice(), product.GetId())

	if err != nil {
		return 0, err
	}

	return res.AffectedRows, nil
}

