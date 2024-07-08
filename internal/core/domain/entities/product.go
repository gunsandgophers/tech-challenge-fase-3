package entities

type Product struct {
	ID          int
	Name        string
	Category    string
	Price       float64
	Description string
	Image       string
}

func NewProduct(name string, category string, price float64, description string, image string) *Product {
	return &Product{
		Name:        name,
		Category:    category,
		Price:       price,
		Description: description,
		Image:       image,
	}
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetID() int {
	return p.ID
}

func (p *Product) GetCategory() string {
	return p.Category
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) GetDescription() string {
	return p.Description
}

func (p *Product) GetImage() string {
	return p.Image
}
