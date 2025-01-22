package domain

import "fmt"

type Product struct {
	id    int32   
	name  string  `json: "name"`
	price float32 	`json: "price"`
}

func NewProduct(name string, price float32) *Product {
	// prod := Product{id: 1, name: name, price: price}
	return &Product{id: 1, name: name, price: price}
}

func (p *Product) ViewProduct() string {
	return fmt.Sprintf("id: %d name: %s price: %.2f", p.id, p.name, p.price)
}

func (p *Product) GetId() int32 {
	return p.id
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) GetPrice() float32 {
	return p.price
}

func (p *Product) SetName(name string) {
	p.name = name
}  