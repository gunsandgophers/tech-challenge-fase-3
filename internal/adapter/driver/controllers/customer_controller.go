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

// GetCustomer godoc
// @Summary      Show a customer
// @Description  get the customer by CPF
// @Tags         customers
// @Accept       json
// @Produce      json
// @Param        cpf   path      string  true  "Customer CPF"
// @Success      200  {object}  dtos.CustomerDTO
// @Failure      404 {string} string "when object not found"
// @Router       /customer/{cpf}/ [get]
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

// RegisterCustomer godoc
// @Summary      Create a customer
// @Description  register the customer information
// @Tags         customers
// @Accept       json
// @Produce      json
// @Param        customer   body      dtos.CreateCustomerDTO  true  "Insert Customer"
// @Success      200  {object}  dtos.CustomerDTO
// @Failure      400 {string} string "when invalid params"
// @Failure      406 {string} string "when invalid status"
// @Router       /customer/ [post]
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
