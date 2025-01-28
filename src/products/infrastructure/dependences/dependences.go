package dependences

import (
	"practica/src/products/infrastructure"
)

var mysql *infrastructure.MySQL

func GoMySQL () {
	mysql = infrastructure.NewMySQL()
}

func GetMySQL() *infrastructure.MySQL {
	return mysql
}




