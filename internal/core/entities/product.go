package entities

import "github.com/google/uuid"

type Product struct {
	id          string
	name        string
	category    string
	price       float64
	description string
	image       string
}

func CreateProduct(name string, category string, price float64, description string, image string) *Product {
	return RestoreProduct(
		uuid.NewString(),
		name,
		category,
		price,
		description,
		image,
	)
}

func RestoreProduct(id string, name string, category string, price float64, description string, image string) *Product {
	return &Product{
		id:          id,
		name:        name,
		category:    category,
		price:       price,
		description: description,
		image:       image,
	}
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) SetName(name string) {
	p.name = name
}

func (p *Product) GetId() string {
	return p.id
}

func (p *Product) GetCategory() string {
	return p.category
}

func (p *Product) SetCategory(category string) {
	p.category = category
}

func (p *Product) GetPrice() float64 {
	return p.price
}

func (p *Product) SetPrice(price float64) {
	p.price = price
}

func (p *Product) GetDescription() string {
	return p.description
}

func (p *Product) SetDescription(description string) {
	p.description = description
}

func (p *Product) GetImage() string {
	return p.image
}

func (p *Product) SetImage(image string) {
	p.image = image
}
