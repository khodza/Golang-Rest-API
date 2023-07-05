package handlers

import (
	"khodza/rest-api/internal/app/models"
	"khodza/rest-api/internal/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	userService services.UserServiceInterface
	logger      *zap.Logger
}

func NewUserHandler(userService services.UserServiceInterface, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      logger,
	}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, CustomError := h.userService.GetUsers()

	if CustomError.StatusCode != 0 {
		SendCustomError(c, CustomError, "Failed to get users", h.logger)
		return
	}

	//logging
	LoggingResponse(c, "GetUsers", h.logger)

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := HandleJSONBinding(c, &user, h.logger); err != nil {
		return
	}

	createdUser, CustomError := h.userService.CreateUser(user)

	if CustomError.StatusCode != 0 {
		SendCustomError(c, CustomError, "Failed to create user", h.logger)
		return
	}

	//logging
	LoggingResponse(c, "CreateUser", h.logger)

	c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userID, err := GetId(c, h.logger)
	if err != nil {
		return
	}

	user, CustomError := h.userService.GetUser(userID)

	if CustomError.StatusCode != 0 {
		SendCustomError(c, CustomError, "Failed to get user", h.logger)
		return
	}

	//logging
	LoggingResponse(c, "GetUser", h.logger)

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	userID, err := GetId(c, h.logger)
	if err != nil {
		return
	}

	var user models.User
	if err := HandleJSONBinding(c, &user, h.logger); err != nil {
		return
	}

	updatedUser, CustomError := h.userService.UpdateUser(userID, user)

	if CustomError.StatusCode != 0 {
		SendCustomError(c, CustomError, "Failed to update user", h.logger)
		return
	}

	//logging
	LoggingResponse(c, "UpdateUser", h.logger)

	c.JSON(http.StatusOK, updatedUser)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID, err := GetId(c, h.logger)
	if err != nil {
		return
	}

	if CustomError := h.userService.DeleteUser(userID); CustomError.StatusCode != 0 {
		SendCustomError(c, CustomError, "Failed to delete user", h.logger)
		return
	}

	//logging
	LoggingResponse(c, "DeleteUser", h.logger)

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
