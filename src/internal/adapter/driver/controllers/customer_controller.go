package controllers

import (
	"net/http"
	"tech-challenge-fase-1/internal/core/applications/repositories"
	usecases "tech-challenge-fase-1/internal/core/applications/use_cases"
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	customerRepository repositories.CustomerRepositoryInterface
}

func NewCustomerController(customerRepository repositories.CustomerRepositoryInterface) *CustomerController {
	return &CustomerController{
		customerRepository: customerRepository,
	};
}

func (cc *CustomerController) GetCustomer(c *gin.Context) {
	cpf := c.Param("cpf")
	getCustomer := usecases.NewGetCustomer(cc.customerRepository)
	customer, err := getCustomer.Execute(cpf)
	if err != nil {
		c.String(http.StatusNotFound, "Customer not found")
		return
	}
	// Aprender a retorna o DTO como JSON
	c.String(http.StatusOK, "Fouded -> " + customer.Name)
}
