package aplication

import "practica/domain"

type CreateProduct struct {
	db domain.IProduct
}

func NewCreateProduct (db domain.IProduct) *CreateProduct {
	return &CreateProduct{db: db}
}

// Run | Execute
func (us *CreateProduct) Run (product domain.Product) {
	us.db.Save(product)
}