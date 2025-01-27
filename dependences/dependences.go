package dependences

import (
	"practica/src/products/infrastructure"
)

func GetMySQL() *infrastructure.MySQL {
	return infrastructure.NewMySQL()
}



