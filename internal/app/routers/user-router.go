package routers

import (
	"khodza/rest-api/internal/app/handlers"

	"github.com/gin-gonic/gin"
)

func SetupUserRouter(router *gin.RouterGroup, userHandler *handlers.UserHandler) {
	userGroup := router.Group("")
	userGroup.GET("/", userHandler.GetUsers)
	userGroup.POST("/", userHandler.CreateUser)
	userGroup.GET("/:id", userHandler.GetUser)
	userGroup.PATCH("/:id", userHandler.UpdateUser)
	userGroup.DELETE("/:id", userHandler.DeleteUser)
}
