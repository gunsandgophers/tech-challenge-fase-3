package products

import (
	"tech-challenge-fase-1/internal/core/repositories"
	"tech-challenge-fase-1/internal/core/dtos"
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
	products, err := lpbcc.productRepository.FindProductByCategory(category, page, size)
	if err != nil {
		return nil, err
	}

	var productsDTO []dtos.ProductDTO
	for _, product := range products {
		productsDTO = append(productsDTO, *dtos.NewProductDTOFromEntity(product))
	}

	return productsDTO, nil
}
