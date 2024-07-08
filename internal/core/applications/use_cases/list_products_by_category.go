package usecases

import (
	"tech-challenge-fase-1/internal/core/applications/repositories"
	"tech-challenge-fase-1/internal/core/domain/dtos"
)

func NewListProductsByCategoryUseCase(productRepository repositories.ProductRepositoryInterface) *ListProductsByCategoryUseCase {
	return &ListProductsByCategoryUseCase{
		productRepository: productRepository,
	}
}

type ListProductsByCategoryUseCase struct {
	productRepository repositories.ProductRepositoryInterface
}

func (lpbcc *ListProductsByCategoryUseCase) Execute(category string, page, size int) ([]dtos.ProductDTO, error) {
	products, err := lpbcc.productRepository.FindByCategory(category, page, size)
	if err != nil {
		return nil, err
	}

	var productsDTO []dtos.ProductDTO
	for _, product := range products {
		productsDTO = append(productsDTO, dtos.ProductDTO{
			ID:          product.ID,
			Name:        product.Name,
			Category:    product.Category,
			Price:       product.Price,
			Description: product.Description,
			Image:       product.Image,
		})
	}

	return productsDTO, nil
}
