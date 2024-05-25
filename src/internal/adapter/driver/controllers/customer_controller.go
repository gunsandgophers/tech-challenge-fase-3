package controllers

import (
	"net/http"
	httpserver "tech-challenge-fase-1/internal/adapter/driven/infra/http"
	"tech-challenge-fase-1/internal/core/applications/repositories"
	usecases "tech-challenge-fase-1/internal/core/applications/use_cases"
	"tech-challenge-fase-1/internal/core/domain/dtos"
)

type CustomerController struct {
	customerRepository repositories.CustomerRepositoryInterface
}

func NewCustomerController(customerRepository repositories.CustomerRepositoryInterface) *CustomerController {
	return &CustomerController{
		customerRepository: customerRepository,
	}
}

func (cc *CustomerController) GetCustomer(c httpserver.HTTPContext) {
	cpf := c.Param("cpf")
	getCustomer := usecases.NewGetCustomer(cc.customerRepository)
	customer, err := getCustomer.Execute(cpf)
	if err != nil {
		sendError(c, http.StatusNotFound, "Customer not found")
		return
	}
	sendSuccess(c, http.StatusOK, "get-customer", customer)
}

func (cc *CustomerController) RegisterCustomer(c httpserver.HTTPContext) {
	request := RegiterCustomerRequest{}
	c.BindJSON(&request)
	if err := request.Validate(); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	insertCustomer := usecases.NewInsertCustomer(cc.customerRepository)
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
