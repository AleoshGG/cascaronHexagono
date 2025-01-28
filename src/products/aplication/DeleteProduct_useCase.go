package aplication

import "practica/src/products/domain"

type DeleteProduct struct {
	db domain.IProduct
}

func NewDeleteProduct(db domain.IProduct) *DeleteProduct {
	return &DeleteProduct{db: db}
}

func (us *DeleteProduct) Run(id int) (uint64, error) {
	return us.db.Delete(id)
}