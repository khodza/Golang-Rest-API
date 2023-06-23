package handlers

import (
	"fmt"
	"khodza/rest-api/internal/app/models"
	"khodza/rest-api/internal/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, CustomError := h.userService.GetUsers()

	if CustomError.StatusCode != 0 {
		c.JSON(CustomError.StatusCode, gin.H{"error": CustomError.Message})
		fmt.Println(CustomError.Err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := HandleJSONBinding(c, &user); err != nil {
		return
	}

	createdUser, CustomError := h.userService.CreateUser(user)

	if CustomError.StatusCode != 0 {
		c.JSON(CustomError.StatusCode, gin.H{"error": CustomError.Message})
		fmt.Println(CustomError.Err.Error())
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userID, err := GetId(c)
	if err != nil {
		return
	}

	user, CustomError := h.userService.GetUser(userID)

	if CustomError.StatusCode != 0 {
		c.JSON(CustomError.StatusCode, gin.H{"message": CustomError.Message})
		fmt.Println(CustomError.Err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	userID, err := GetId(c)
	if err != nil {
		return
	}

	var user models.User
	if err := HandleJSONBinding(c, &user); err != nil {
		return
	}

	updatedUser, CustomError := h.userService.UpdateUser(userID, user)
	if CustomError.StatusCode != 0 {
		c.JSON(CustomError.StatusCode, gin.H{"message": CustomError.Message})
		fmt.Println(CustomError.Err.Error())
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID, err := GetId(c)
	if err != nil {
		return
	}
	if CustomError := h.userService.DeleteUser(userID); CustomError.StatusCode != 0 {
		c.JSON(CustomError.StatusCode, gin.H{"message": CustomError.Message})
		fmt.Println(CustomError.Err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
