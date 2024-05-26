package controllers

import (
	"net/http"
	httpserver "tech-challenge-fase-1/internal/adapter/driven/infra/http"
	"tech-challenge-fase-1/internal/adapter/driver/controllers/request"
	"tech-challenge-fase-1/internal/core/applications/events"
	"tech-challenge-fase-1/internal/core/applications/repositories"
	"tech-challenge-fase-1/internal/core/applications/services"
	usecases "tech-challenge-fase-1/internal/core/applications/use_cases/order"
)

type OrderController struct {
	orderRepository    repositories.OrderRepositoryInterface
	customerRepository repositories.CustomerRepositoryInterface
	productRepository  repositories.ProductRepositoryInterface

	paymentGateway services.PaymentGatewayInterface

	commandEventManager events.ManagerEvent
}

func NewOrderController(
	orderRepository repositories.OrderRepositoryInterface,
	customerRepository repositories.CustomerRepositoryInterface,
	productRepository repositories.ProductRepositoryInterface,
	paymentGateway services.PaymentGatewayInterface,
	commandEventManager events.ManagerEvent,
) *OrderController {
	return &OrderController{
		orderRepository:     orderRepository,
		customerRepository:  customerRepository,
		productRepository:   productRepository,
		paymentGateway:      paymentGateway,
		commandEventManager: commandEventManager,
	}
}

func (cc *OrderController) OpenOrder(c httpserver.HTTPContext) {

	request := request.OpenOrderRequest{}
	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	openOrderUseCase := usecases.NewOpenOrderUseCase(
		cc.orderRepository,
		cc.customerRepository,
	)

	order, err := openOrderUseCase.Execute(request.CustomerID)

	if err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	sendSuccess(c, http.StatusCreated, "open-order", order)
}

func (cc *OrderController) AddOrderItem(c httpserver.HTTPContext) {

	request := request.AddOrderItemRequest{}
	c.BindJSON(&request)

	request.OrderID = c.Param("order_id")

	if err := request.Validate(); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	addItemOrderUseCase := usecases.NewAddOrderItemUseCase(
		cc.orderRepository,
		cc.productRepository,
	)

	order, err := addItemOrderUseCase.Execute(&usecases.AddOrderItemUseCaseRequest{
		ProductID: request.ProductID,
		Quantity:  request.Quantity,
		OrderID:   request.OrderID,
	})

	if err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	sendSuccess(c, http.StatusCreated, "add-order-item", order)
}

func (cc *OrderController) Checkout(c httpserver.HTTPContext) {

	orderID := c.Param("order_id")

	checkoutUseCase := usecases.NewCheckoutOrderUseCase(
		cc.orderRepository,
		cc.paymentGateway,
		cc.commandEventManager)

	checkout, err := checkoutUseCase.Execute(orderID)
	if err != nil {
		sendError(c, 400, err.Error())
		return
	}

	sendSuccess(c, http.StatusCreated, "checkout-order", checkout)

}
