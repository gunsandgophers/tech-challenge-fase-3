package products

import (
	"log"
	"strings"
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/core/repositories"
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
	productCategory := entities.ProductCategory(strings.ToUpper(category))
	log.Println(productCategory)
	products, err := lpbcc.productRepository.FindProductByCategory(productCategory, page, size)
	if err != nil {
		return nil, err
	}

	log.Println(products)

	var productsDTO []dtos.ProductDTO
	for _, product := range products {
		productsDTO = append(productsDTO, *dtos.NewProductDTOFromEntity(product))
	}

	return productsDTO, nil
}
