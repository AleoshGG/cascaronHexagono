package aplication

import "practica/src/products/domain"

type CreateProduct struct {
	db domain.IProduct
}

func NewCreateProduct(db domain.IProduct) *CreateProduct {
	return &CreateProduct{db: db}
}

// Run | Execute
func (us *CreateProduct) Run(product domain.Product) (uint64, error) {
	return us.db.Save(product)
}