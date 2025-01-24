package domain

import "fmt"

type Product struct {
	id    int32   
	name  string  
	price float32 	
}

func NewProduct(id int32, name string, price float32) *Product {
	// prod := Product{id: 1, name: name, price: price}
	return &Product{id: id, name: name, price: price}
}

func (p *Product) ViewProduct() string {
	return fmt.Sprintf(" name: %s price: %.2f", p.name, p.price)
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) GetPrice() float32 {
	return p.price
}

func (p *Product) SetId(id int32) {
	p.id = id
} 

func (p *Product) SetName(name string) {
	p.name = name
} 

func (p *Product) SetPrice(price float32) {
	p.price = price
} 