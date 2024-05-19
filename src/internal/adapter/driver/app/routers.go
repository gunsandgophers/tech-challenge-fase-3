package app

import (
	"tech-challenge-fase-1/internal/adapter/driver/controllers"
)

// Registra as rotas dos controllers
func registerRouters(app *APIApp) {
	helloController := controllers.NewHelloController()
	userController := controllers.NewUserController(app.userRepository)
	customerController := controllers.NewCustomerController(app.customerRepository)
	app.httpServer.GET("/", helloController.Index)
	app.httpServer.GET("/users/", userController.Index)

	app.httpServer.GET("/api/v1/customer/:cpf/", customerController.GetCustomer)
}
