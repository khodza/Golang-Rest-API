package handlers

import (
	"khodza/rest-api/internal/app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HandlerUtilities struct {
	logger *zap.Logger
}

func NewHandlerUtilities(logger *zap.Logger) *HandlerUtilities {
	return &HandlerUtilities{
		logger: logger,
	}
}

func (u *HandlerUtilities) GetId(c *gin.Context) (int, error) {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Id provided"})

		//logging
		u.logger.Error("Invalid Id provided",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", http.StatusBadRequest),
			zap.Error(err))
		return 0, err
	}
	return ID, nil
}

func (u *HandlerUtilities) HandleJSONBinding(c *gin.Context, target interface{}) error {
	if err := c.ShouldBindJSON(&target); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})

		//logging
		u.logger.Error("Failed to bind JSON",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", http.StatusBadRequest),
			zap.Error(err))
		return err
	}
	return nil
}

func (u *HandlerUtilities) SendCustomError(c *gin.Context, CustomError services.CustomError, zapMessage string) {
	c.JSON(CustomError.StatusCode, gin.H{"error": CustomError.Message})

	//logging
	u.logger.Error(zapMessage+"\n"+CustomError.Message,
		zap.String("method", c.Request.Method),
		zap.String("path", c.Request.URL.Path),
		zap.Int("status", CustomError.StatusCode),
		zap.Error(CustomError.Err))
}

func (u *HandlerUtilities) LoggingResponse(c *gin.Context, info string) {
	u.logger.Info(info,
		zap.String("method", c.Request.Method),
		zap.String("path", c.Request.URL.Path),
		zap.Int("status", http.StatusOK),
	)
}
