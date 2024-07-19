package controllers

import (
	"net/http"
	httpserver "tech-challenge-fase-1/internal/infra/http"
	"tech-challenge-fase-1/internal/core/events"
	"tech-challenge-fase-1/internal/core/queries"
	"tech-challenge-fase-1/internal/core/repositories"
	"tech-challenge-fase-1/internal/core/services"
	"tech-challenge-fase-1/internal/core/use_cases/orders"
)

type OrderController struct {
	orderRepository       repositories.OrderRepositoryInterface
	customerRepository    repositories.CustomerRepositoryInterface
	productRepository     repositories.ProductRepositoryInterface
	paymentGateway        services.PaymentGatewayInterface
	commandEventManager   events.ManagerEvent
	orderDisplayListQuery queries.OrderDisplayListQueryInterface
}

func NewOrderController(
	orderRepository repositories.OrderRepositoryInterface,
	customerRepository repositories.CustomerRepositoryInterface,
	productRepository repositories.ProductRepositoryInterface,
	paymentGateway services.PaymentGatewayInterface,
	commandEventManager events.ManagerEvent,
	orderDisplayListQuery queries.OrderDisplayListQueryInterface,
) *OrderController {
	return &OrderController{
		orderRepository:       orderRepository,
		customerRepository:    customerRepository,
		productRepository:     productRepository,
		paymentGateway:        paymentGateway,
		commandEventManager:   commandEventManager,
		orderDisplayListQuery: orderDisplayListQuery,
	}
}

// Checkout godoc
// @Summary      Do a order checkout
// @Description  do a checkout on a given order
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        order_id   path      string  true  "Open Order"
// @Success      200 {object} dtos.CheckoutDTO
// @Failure      400 {string} string "when invalid status"
// @Router       /order/{order_id}/add/item [post]
func (cc *OrderController) Checkout(c httpserver.HTTPContext) {
	request := CheckoutRequest{}
	c.BindJSON(&request)
	if err := request.Validate(); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	checkoutUseCase := orders.NewCheckoutOrderUseCase(
		cc.orderRepository,
		cc.customerRepository,
		cc.productRepository,
		cc.paymentGateway,
		cc.commandEventManager,
	)
	checkout, err := checkoutUseCase.Execute(request.CustomerId, request.ProductsIds)
	if err != nil {
		sendError(c, http.StatusNotAcceptable, err.Error())
		return
	}
	sendSuccess(c, http.StatusCreated, "checkout-order", checkout)
}

func (cc *OrderController) GetPaymentStatus(c httpserver.HTTPContext) {
	orderId := c.Param("order_id")
	getPaymentStatusUC := orders.NewGetPaymentStatusUseCase(cc.orderRepository)
	paymentStatus, err := getPaymentStatusUC.Execute(orderId)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	sendSuccess(c, http.StatusOK, "get-payment-status-order", paymentStatus)
}

func (cc *OrderController) Payment(c httpserver.HTTPContext) {
	request := &PaymentRequest{}
	c.BindJSON(request)
	if err := request.Validate(); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	paymentUseCase := orders.NewPaymentOrderUseCase(
		cc.orderRepository,
	)
	err := paymentUseCase.Execute(request.OrderId, request.PaymentStatus)
	if err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	sendSuccess(c, http.StatusNoContent, "payment-order", nil)
}

func (cc *OrderController) OrderDisplayList(c httpserver.HTTPContext) {
	orderDisplayListUseCase := orders.NewOrderDisplayListUseCase(
		cc.orderDisplayListQuery,
	)
	dtos, err := orderDisplayListUseCase.Execute()
	if err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	sendSuccess(c, http.StatusOK, "order-display-list", httpserver.Payload{
		"orders": dtos,
	})
}
