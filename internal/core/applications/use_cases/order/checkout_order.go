package usecases

import (
	"tech-challenge-fase-1/internal/core/applications/events"
	"tech-challenge-fase-1/internal/core/applications/repositories"
	"tech-challenge-fase-1/internal/core/applications/services"
	"tech-challenge-fase-1/internal/core/domain/dtos"
	"tech-challenge-fase-1/internal/core/domain/entities"
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

	order.Status = entities.AWAITING_PAYMENT

	checkout, err := c.paymentGateway.Execute(dtos.NewOrderDTOFromEntity(order), dtos.PIX)
	if err != nil {
		return nil, err
	}

	err = c.orderRepository.UpdateStatus(order)
	if err != nil {
		return nil, err
	}

	go func() {
		c.commandEventManager.Add("paid_out", func() {
			order.Status = entities.Paid
			c.orderRepository.UpdateStatus(order)
		})
	}()

	return checkout, nil
}
