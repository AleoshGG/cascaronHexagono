package domain

import "fmt"

type Product struct {
	id    int32   
	name  string  
	price float32 	
}

func NewProduct(name string, price float32) *Product {
	// prod := Product{id: 1, name: name, price: price}
	return &Product{name: name, price: price}
}

func (p *Product) ViewProduct() string {
	return fmt.Sprintf("name: %s price: %.2f", p.name, p.price)
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