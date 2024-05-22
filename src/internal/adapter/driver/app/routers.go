package app

import (
	"tech-challenge-fase-1/internal/adapter/driver/controllers"
)

// Registra as rotas dos controllers
func registerRouters(app *APIApp) {
	helloController := controllers.NewHelloController()
	customerController := controllers.NewCustomerController(app.customerRepository)
	productController := controllers.NewProductController(app.productRepository)

	baseUrl := "/api/v1"
	app.httpServer.SetBasePath(baseUrl)
	app.httpServer.GET("/", helloController.Index)

	//customer
	app.httpServer.POST("/customer/", customerController.RegisterCustomer)
	app.httpServer.GET("/customer/:cpf/", customerController.GetCustomer)

	app.httpServer.POST("/product", productController.CreateProduct)
	app.httpServer.PUT("/product/:id", productController.UpdateProduct)
	app.httpServer.DELETE("/product/:id", productController.DeleteProduct)
	app.httpServer.GET("/products/:category", productController.ListProductsByCategory)
}
