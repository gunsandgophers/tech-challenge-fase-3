package controllers

import (
	"net/http"
	"tech-challenge-fase-1/internal/core/applications/repositories"
	usecases "tech-challenge-fase-1/internal/core/applications/use_cases"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository repositories.UserRepositoryInterface
}

func NewUserController(userRepository repositories.UserRepositoryInterface) *UserController {
	return &UserController{
		userRepository: userRepository,
	};
}

func (uc *UserController) Index(c *gin.Context) {

	listUsers := usecases.NewListUsers(uc.userRepository)
	users := listUsers.Execute()
	result := "Users List:"
	for _, user := range users {
		result += " " + user.Name
	}
	c.String(http.StatusOK, "User Controller -> " + result)
}
