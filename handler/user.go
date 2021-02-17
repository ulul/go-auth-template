package handler

import (
	"crownfunding/auth"
	"crownfunding/helper"
	"crownfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

// NewUserHandler service
func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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

	token, err := h.authService.GenerateToken(newUser.ID)

	if err != nil {
		response := helper.APIResponse("Generate token failed", http.StatusBadRequest, false, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newUser, token)

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

	token, err := h.authService.GenerateToken(userLogin.ID)

	if err != nil {
		response := helper.APIResponse("Generate token failed", http.StatusBadRequest, false, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(userLogin, token)

	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusUnauthorized, false, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Login success", http.StatusOK, true, formatter)

	c.JSON(http.StatusOK, response)
}

// CheckAvailabilityEmail function
func (h *userHandler) CheckAvailabilityEmail(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessage := helper.FormatValidationError(err)

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, false, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	IsEmailAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {
		response := helper.APIResponse("Check email failed", http.StatusUnprocessableEntity, false, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_available": IsEmailAvailable,
	}
	metaMessage := "Email has been registered"

	if IsEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, true, data)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	// c.SaveUploadedFile()

	file, err := c.FormFile("avatar")

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar", http.StatusUnprocessableEntity, false, data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	avatarPath := "avatar/" + file.Filename
	err = c.SaveUploadedFile(file, avatarPath)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar", http.StatusUnprocessableEntity, false, data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID := 6 //

	userWithAvatar, err := h.userService.SaveAvatar(userID, avatarPath)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar", http.StatusUnprocessableEntity, false, data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Avatar uploaded", http.StatusOK, true, userWithAvatar)

	c.JSON(http.StatusOK, response)
}
