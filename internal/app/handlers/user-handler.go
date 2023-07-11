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
	users, err := h.userService.GetUsers()

	if err != nil {
		c.Error(err)
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

	createdUser, err := h.userService.CreateUser(user)

	if err != nil {
		c.Error(err)
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

	user, err := h.userService.GetUser(userID)

	if err != nil {
		c.Error(err)
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

	updatedUser, err := h.userService.UpdateUser(userID, user)

	if err != nil {
		c.Error(err)
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

	if err := h.userService.DeleteUser(userID); err != nil {
		c.Error(err)
		return
	}

	//logging
	LoggingResponse(c, "DeleteUser", h.logger)

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
