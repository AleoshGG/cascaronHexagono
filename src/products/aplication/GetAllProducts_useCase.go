package aplication

import "practica/src/products/domain"

type GetAllProducts struct {
	db domain.IProduct
}

func NewGetAllProducts(db domain.IProduct) *GetAllProducts {
	return &GetAllProducts{db: db}
}

func (us *GetAllProducts) Run() ([]domain.Product, error) {
	return us.db.GetAll()
}