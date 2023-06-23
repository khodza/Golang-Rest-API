package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetId(ctx *gin.Context) (int, error) {
	id := ctx.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Id provided"})
		return 0, err
	}
	return ID, nil
}

func HandleJSONBinding(c *gin.Context, target interface{}) error {
	if err := c.ShouldBindJSON(&target); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return err
	}
	return nil
}
