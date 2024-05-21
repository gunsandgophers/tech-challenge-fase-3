package controllers

import (
	"net/http"
	httpserver "tech-challenge-fase-1/internal/adapter/driven/infra/http"
)

type HelloController struct {
}

func NewHelloController() *HelloController {
	return &HelloController{};
}

func (h *HelloController) Index(c httpserver.HTTPContext) {
	c.JSON(http.StatusOK, httpserver.Payload{"msg": "Hello World! :)"})
}
