package handlers

import (
	"khodza/rest-api/internal/app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetId(c *gin.Context, logger *zap.Logger) (int, error) {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Id provided"})

		//logging
		logger.Error("Invalid Id provided",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", http.StatusBadRequest),
			zap.Error(err))
		return 0, err
	}
	return ID, nil
}

func HandleJSONBinding(c *gin.Context, target interface{}, logger *zap.Logger) error {
	if err := c.ShouldBindJSON(&target); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})

		//logging
		logger.Error("Failed to bind JSON",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", http.StatusBadRequest),
			zap.Error(err))
		return err
	}
	return nil
}

func SendCustomError(c *gin.Context, CustomError services.CustomError, zapMessage string, logger *zap.Logger) {
	c.JSON(CustomError.StatusCode, gin.H{"error": CustomError.Message})

	//logging
	logger.Error(zapMessage+"\n"+CustomError.Message,
		zap.String("method", c.Request.Method),
		zap.String("path", c.Request.URL.Path),
		zap.Int("status", CustomError.StatusCode),
		zap.Error(CustomError.Err))
}

func LoggingResponse(c *gin.Context, info string, logger *zap.Logger) {
	logger.Info(info,
		zap.String("method", c.Request.Method),
		zap.String("path", c.Request.URL.Path),
		zap.Int("status", http.StatusOK),
	)
}
