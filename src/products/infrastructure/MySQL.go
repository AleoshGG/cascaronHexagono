package infrastructure

import (
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

func (mysql *MySQL) GetAll() ([]domain.Product, error) {

	query := "SELECT id, name, price FROM products"
	_, err := mysql.conn.Execute(query)
	if err != nil {
		return nil, err
	}

	var products []domain.Product

	/* for _, row := range res.Values {
		id, ok1 := row[0].(int64) // Cambié la conversión a int64, ya que MySQL generalmente devuelve int64 para id
		name, ok2 := row[1].(string) // Convierte el valor a string
		price, ok3 := row[2].(float64) // Convierte el valor a float64

		// Validar conversiones
		if !ok1 || !ok2 || !ok3 {
			return nil, fmt.Errorf("error al convertir los valores de la fila")
		}

		product := domain.NewProduct(id,name,price)
		// Agregar el producto al slice
		products = append(products, *product)
	} */

	return products, nil
}

