package handlers

import (
	"khodza/rest-api/internal/app/models"
	"khodza/rest-api/internal/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	userService services.UserService
	utils       HandlerUtilities
}

func NewUserHandler(userService services.UserService, utils HandlerUtilities) *UserHandler {
	return &UserHandler{
		userService: userService,
		utils:       utils,
	}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, CustomError := h.userService.GetUsers()

	if CustomError.StatusCode != 0 {
		h.utils.SendCustomError(c, CustomError, "Failed to get users")
		return
	}

	h.utils.logger.Info("GetUsers",
		zap.String("method", c.Request.Method),
		zap.String("path", c.Request.URL.Path),
		zap.Int("status", http.StatusOK),
	)

	//logging
	h.utils.LoggingResponse(c, "GetUsers")

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := h.utils.HandleJSONBinding(c, &user); err != nil {
		return
	}

	createdUser, CustomError := h.userService.CreateUser(user)

	if CustomError.StatusCode != 0 {
		h.utils.SendCustomError(c, CustomError, "Failed to create user")
		return
	}

	//logging
	h.utils.LoggingResponse(c, "CreateUser")

	c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userID, err := h.utils.GetId(c)
	if err != nil {
		return
	}

	user, CustomError := h.userService.GetUser(userID)

	if CustomError.StatusCode != 0 {
		h.utils.SendCustomError(c, CustomError, "Failed to get user")
		return
	}

	//logging
	h.utils.LoggingResponse(c, "GetUser")

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	userID, err := h.utils.GetId(c)
	if err != nil {
		return
	}

	var user models.User
	if err := h.utils.HandleJSONBinding(c, &user); err != nil {
		return
	}

	updatedUser, CustomError := h.userService.UpdateUser(userID, user)

	if CustomError.StatusCode != 0 {
		h.utils.SendCustomError(c, CustomError, "Failed to update user")
		return
	}

	//logging
	h.utils.LoggingResponse(c, "UpdateUser")

	c.JSON(http.StatusOK, updatedUser)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID, err := h.utils.GetId(c)
	if err != nil {
		return
	}

	if CustomError := h.userService.DeleteUser(userID); CustomError.StatusCode != 0 {
		h.utils.SendCustomError(c, CustomError, "Failed to delete user")
		return
	}

	//logging
	h.utils.LoggingResponse(c, "DeleteUser")

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
