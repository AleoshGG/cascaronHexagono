package main

import (
	"practica/aplication"
	"practica/infrastructure"
)

func main() {
	mysql := infrastructure.NewMySQL()
	cp := aplication.NewCreateProduct(mysql)
	cp.Run()
}
