package usecases

import (
	"tech-challenge-fase-1/internal/core/applications/repositories"
	"tech-challenge-fase-1/internal/core/domain/dtos"
	"tech-challenge-fase-1/internal/core/domain/entities"
)

type CreateProduct struct {
	productRepository repositories.ProductRepositoryInterface
}

func NewCreateProductUseCase(productRepository repositories.ProductRepositoryInterface) *CreateProduct {
	return &CreateProduct{
		productRepository: productRepository,
	}
}

func (cp *CreateProduct) Execute(productDTO *dtos.ProductDTO) (*dtos.ProductDTO, error) {
	product := entities.NewProduct(productDTO.Name, productDTO.Category, productDTO.Price, productDTO.Description, productDTO.Image)

	err := cp.productRepository.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return dtos.ConvertProductEntityToDTO(product), nil
}
