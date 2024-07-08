package usecases

import (
	"tech-challenge-fase-1/internal/core/applications/repositories"
	"tech-challenge-fase-1/internal/core/domain/dtos"
	"tech-challenge-fase-1/internal/core/domain/entities"
)

type UpdateProductUseCase struct {
	productRepository repositories.ProductRepositoryInterface
}

func NewUpdateProductUseCase(productRepository repositories.ProductRepositoryInterface) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		productRepository: productRepository,
	}
}

func (upuc *UpdateProductUseCase) Execute(productDTO *dtos.ProductDTO) (*entities.Product, error) {
	product, err := upuc.productRepository.FindByID(productDTO.ID)
	if err != nil {
		return nil, err
	}

	updateProductFromDTO(product, productDTO)

	updatedProduct, err := upuc.productRepository.UpdateProduct(product)
	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func updateProductFromDTO(product *entities.Product, productDTO *dtos.ProductDTO) {
	if productDTO.Name != "" {
		product.Name = productDTO.Name
	}

	if productDTO.Category != "" {
		product.Category = productDTO.Category
	}

	if productDTO.Price != 0 {
		product.Price = productDTO.Price
	}

	if productDTO.Description != "" {
		product.Description = productDTO.Description
	}

	if productDTO.Image != "" {
		product.Image = productDTO.Image
	}
}
