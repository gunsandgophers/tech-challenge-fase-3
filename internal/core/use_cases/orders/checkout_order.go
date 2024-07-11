package orders

import (
	"tech-challenge-fase-1/internal/core/events"
	"tech-challenge-fase-1/internal/core/repositories"
	"tech-challenge-fase-1/internal/core/services"
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/core/entities"
)

type CheckoutOrderUseCase struct {
	orderRepository     repositories.OrderRepositoryInterface
	paymentGateway      services.PaymentGatewayInterface
	commandEventManager events.ManagerEvent
}

func NewCheckoutOrderUseCase(
	orderRepository repositories.OrderRepositoryInterface,
	paymentGateway services.PaymentGatewayInterface,
	commandEventManager events.ManagerEvent,
) *CheckoutOrderUseCase {
	return &CheckoutOrderUseCase{
		orderRepository:     orderRepository,
		paymentGateway:      paymentGateway,
		commandEventManager: commandEventManager,
	}
}

func (c *CheckoutOrderUseCase) Execute(orderID string) (*dtos.CheckoutDTO, error) {

	order, err := c.orderRepository.FindByID(&orderID)
	if err != nil {
		return nil, err
	}

	order.SetStatus(entities.AWAITING_PAYMENT)

	checkout, err := c.paymentGateway.Execute(dtos.NewOrderDTOFromEntity(order), dtos.PIX)
	if err != nil {
		return nil, err
	}

	// err = c.orderRepository.UpdateStatus(order)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// go func() {
	// 	c.commandEventManager.Add("paid_out", func() {
	// 		order.SetStatus(entities.PAID)
	// 		c.orderRepository.UpdateStatus(order)
	// 	})
	// }()

	return checkout, nil
}
