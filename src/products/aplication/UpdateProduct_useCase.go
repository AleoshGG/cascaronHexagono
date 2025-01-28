package aplication

import "practica/src/products/domain"

type UpdateProduct struct {
	db domain.IProduct
}

func NewUpdateProduct(db domain.IProduct) *UpdateProduct {
	return &UpdateProduct{db: db}
}

func (us *UpdateProduct) Run(product domain.Product) (uint64, error) {
	return us.db.Update(product)
}