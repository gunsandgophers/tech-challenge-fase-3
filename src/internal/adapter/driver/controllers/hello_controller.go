package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type HelloController struct {
}

func NewHelloController() *HelloController {
	return &HelloController{};
}

func (h *HelloController) Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World! :)")
}
