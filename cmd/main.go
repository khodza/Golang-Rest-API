package main

import (
	"fmt"
	"khodza/rest-api/internal/app/handlers"
	"khodza/rest-api/internal/app/routers"
	"khodza/rest-api/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	//Loading env
	config.LoadEnv()

	//Initialize DataBase
	config.InitDataBase()

	// Initialize dependencies
	handlersMap := config.InitDependencies()

	// Initialize Gin router
	router := gin.Default()

	// Connect routers to handlers
	for route, handler := range handlersMap {
		routeGroup := router.Group("/" + route)
		switch route {
		case "users":
			userHandler := handler.(*handlers.UserHandler)
			routers.SetupUserRouter(routeGroup, userHandler)
		case "products":
			productHandler := handler.(*handlers.ProductHandler)
			routers.SetupProductRouter(routeGroup, productHandler)
		}

	}
	// Run the server
	router.Run(":8080")
	fmt.Println(router)
}
