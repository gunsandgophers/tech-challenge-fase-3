package orders

import (
	"tech-challenge-fase-1/internal/core/repositories"
	"tech-challenge-fase-1/internal/core/dtos"
)

type AddOrderItemUseCaseRequest struct {
	ProductID string
	Quantity  int
	OrderID   string
}

type AddOrderItemUseCase struct {
	orderRepository   repositories.OrderRepositoryInterface
	productRepository repositories.ProductRepositoryInterface
}

func NewAddOrderItemUseCase(
	orderRepository repositories.OrderRepositoryInterface,
	productRepository repositories.ProductRepositoryInterface,
) *AddOrderItemUseCase {
	return &AddOrderItemUseCase{
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}
}

func (co *AddOrderItemUseCase) Execute(request *AddOrderItemUseCaseRequest) (*dtos.OrderDTO, error) {
	product, err := co.productRepository.FindProductByID(request.ProductID)
	if err != nil {
		return nil, err
	}
	order, err := co.orderRepository.FindOrderByID(request.OrderID)
	if err != nil {
		return nil, err
	}
	order.AddItem(product, request.Quantity)
	co.orderRepository.Update(order)
	return dtos.NewOrderDTOFromEntity(order), nil
}
