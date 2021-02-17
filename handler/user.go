package handler

import (
	"crownfunding/helper"
	"crownfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

// NewUserHandler service
func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := helper.FormatValidationError(err)

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, false, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	formatter := user.FormatUser(newUser, "randomStringToken")

	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, false, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Account has been registered", http.StatusOK, true, formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := helper.FormatValidationError(err)

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, false, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	userLogin, err := h.userService.Login(input)
	formatter := user.FormatUser(userLogin, "randomStringToken")

	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, false, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Login success", http.StatusOK, true, formatter)

	c.JSON(http.StatusOK, response)

}
