package custom_errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GlobalErrorHandler struct {
	logger *zap.Logger
}

func NewGlobalErrorHandler(logger *zap.Logger) *GlobalErrorHandler {
	return &GlobalErrorHandler{
		logger: logger,
	}
}
func (h *GlobalErrorHandler) HandleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check for errors in the request context
		if len(c.Errors) > 0 {

			// Determine the error status and message
			status := http.StatusInternalServerError
			message := "Internal Server Error"

			for _, err := range c.Errors.Errors() {
				if err == ErrEmailExist.Error() {
					status = http.StatusConflict
					message = "Email already exists"
					break
				}
				if err == ErrUserNotFound.Error() {
					status = http.StatusNotFound
					message = "User not found"
					break
				}
				if IsValidationErr(err) {
					status = http.StatusBadRequest
					message = err
					break
				}
				//other errors
			}

			// Log the error
			h.logger.Error(message,
				zap.String("method", c.Request.Method),
				zap.String("path", c.Request.URL.Path),
				zap.Int("status", status),
			)
			// Respond with the appropriate status code and error message
			c.JSON(status, gin.H{
				"error": message,
			})
		}
	}
}
