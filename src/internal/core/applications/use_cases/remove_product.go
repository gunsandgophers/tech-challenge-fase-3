package usecases

import (
	"tech-challenge-fase-1/internal/core/applications/repositories"
)

func NewDeleteProductUseCase(productRepository repositories.ProductRepositoryInterface) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		productRepository: productRepository,
	}
}

type DeleteProductUseCase struct {
	productRepository repositories.ProductRepositoryInterface
}

func (dpc *DeleteProductUseCase) Execute(productID int) error {
	_, err := dpc.productRepository.FindByID(productID)
	if err != nil {
		return err
	}

	return dpc.productRepository.DeleteProduct(productID)
}
