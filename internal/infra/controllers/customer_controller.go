package controllers

import (
	"net/http"
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/core/repositories"
	"tech-challenge-fase-1/internal/core/use_cases/customers"
	httpserver "tech-challenge-fase-1/internal/infra/http"
)

type CustomerController struct {
	customerRepository repositories.CustomerRepositoryInterface
}

func NewCustomerController(customerRepository repositories.CustomerRepositoryInterface) *CustomerController {
	return &CustomerController{
		customerRepository: customerRepository,
	}
}

// GetCustomer godoc
//
//	@Summary		Show a customer
//	@Description	get the customer by CPF
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			cpf	path		string	true	"Customer CPF"
//	@Success		200	{object}	dtos.CustomerDTO
//	@Failure		404	{string}	string	"when object not found"
//	@Router			/customer/{cpf}/ [get]
func (cc *CustomerController) GetCustomer(c httpserver.HTTPContext) {
	cpf := c.Param("cpf")
	getCustomer := customers.NewGetCustomerUseCase(cc.customerRepository)
	customer, err := getCustomer.Execute(cpf)
	if err != nil {
		sendError(c, http.StatusNotFound, "Customer not found")
		return
	}
	sendSuccess(c, http.StatusOK, "get-customer", customer)
}

// RegisterCustomer godoc
//
//	@Summary		Create a customer
//	@Description	register the customer information
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			customer	body		dtos.CreateCustomerDTO	true	"Insert Customer"
//	@Success		200			{object}	dtos.CustomerDTO
//	@Failure		400			{string}	string "when bad request"
//	@Failure		406			{string}	string "when invalid params or invalid object"
//	@Router			/customer/ [post]
func (cc *CustomerController) RegisterCustomer(c httpserver.HTTPContext) {
	request := RegiterCustomerRequest{}
	c.BindJSON(&request)
	if err := request.Validate(); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	insertCustomer := customers.NewInsertCustomerUseCase(cc.customerRepository)
	dto := &dtos.CreateCustomerDTO{
		Name:  request.Name,
		Email: request.Email,
		Cpf:   request.CPF,
	}

	customer, err := insertCustomer.Execute(dto)
	if err != nil {
		sendError(c, http.StatusNotAcceptable, err.Error())
		return
	}
	sendSuccess(c, http.StatusCreated, "register-customer", customer)
}
