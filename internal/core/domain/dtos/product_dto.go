package dtos

import "tech-challenge-fase-1/internal/core/domain/entities"

type ProductDTO struct {
	ID          int
	Name        string
	Category    string
	Price       float64
	Description string
	Image       string
}

func ConvertProductEntityToDTO(product *entities.Product) *ProductDTO {
	return &ProductDTO{
		ID:          product.GetID(),
		Name:        product.GetName(),
		Category:    product.GetCategory(),
		Price:       product.GetPrice(),
		Description: product.GetDescription(),
		Image:       product.GetImage(),
	}
}
