package app

import (
	"tech-challenge-fase-1/internal/adapter/driver/controllers"
)

// Registra as rotas dos controllers
func registerRouters(app *APIApp) {
	helloController := controllers.NewHelloController()
	customerController := controllers.NewCustomerController(app.customerRepository)

	baseUrl := "/api/v1"
	v1 := app.httpServer.Group(baseUrl)
	{
		v1.GET("/", helloController.Index)
		v1.GET("/customer/inserir/", customerController.InsertCustomer)
		v1.GET("/customer/:cpf/", customerController.GetCustomer)
	}
}
