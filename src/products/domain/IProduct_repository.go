package domain

type IProduct interface {
	Save(product Product) (uint64, error)
	GetAll() ([]Product, error)
	Update(product Product) (uint64, error)
}