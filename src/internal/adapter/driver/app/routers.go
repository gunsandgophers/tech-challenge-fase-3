package app

import (
	"tech-challenge-fase-1/internal/adapter/driver/controllers"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Registra as rotas dos controllers
func registerRouters(app *APIApp) {
	helloController := controllers.NewHelloController()
	customerController := controllers.NewCustomerController(app.customerRepository)
	productController := controllers.NewProductController(app.productRepository)
	orderController := controllers.NewOrderController(
		app.orderRepository,
		app.customerRepository,
		app.productRepository,
		app.mercadoPagoGateway,
		app.eventManager,
	)

	baseUrl := "/api/v1"
	app.httpServer.SetBasePath(baseUrl)
	app.httpServer.GET("/", helloController.Index)

	//customer
	app.httpServer.POST("/customer/", customerController.RegisterCustomer)
	app.httpServer.GET("/customer/:cpf/", customerController.GetCustomer)

	//products
	app.httpServer.POST("/product", productController.CreateProduct)
	app.httpServer.PUT("/product/:id", productController.UpdateProduct)
	app.httpServer.DELETE("/product/:id", productController.DeleteProduct)
	app.httpServer.GET("/products/:category", productController.ListProductsByCategory)

	//orders
	app.httpServer.POST("/order/open", orderController.OpenOrder)
	app.httpServer.POST("/order/:order_id/add/item", orderController.AddOrderItem)
	app.httpServer.POST("/order/:order_id/checkout", orderController.Checkout)

	app.httpServer.SetSwagger("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
