package routers

import (
	infra "tech-challenge-fase-1/internal/adapter/driven/infra/repositories"
	"tech-challenge-fase-1/internal/adapter/driver/controllers"

	"github.com/gin-gonic/gin"
)

// Registra as rotas dos controllers
func RegisterRouters(router *gin.Engine) *gin.Engine {
	userRepository := infra.NewUserRepositoryDB()
	helloController := controllers.NewHelloController()
	userController := controllers.NewUserController(userRepository)

	router.GET("/", helloController.Index)
	router.GET("/users/", userController.Index)
	return router
}
