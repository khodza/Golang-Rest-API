package main

import (
	"fmt"
	"khodza/rest-api/internal/app/routers"
	"khodza/rest-api/internal/config"
	"khodza/rest-api/internal/database"
	"khodza/rest-api/internal/dependencies"
	"net/http"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

func main() {
	//Loading env
	config.LoadEnv()

	//Initialize DataBase
	err := database.InitDataBase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
	}

	// Initialize dependencies
	globalErrorHandler, handlersMap, logger, err := dependencies.InitDependencies()
	if err != nil {
		fmt.Println("Failed to initialize dependencies")
	}
	// Initialize Gin router
	router := gin.Default()
	// Define the global error handler middleware
	router.Use(globalErrorHandler.HandleErrors())

	// Connect routers to handlers
	routers.ConnectRoutersToHandlers(router, handlersMap)

	// Start the server
	port := ":" + config.GetEnv("PORT", "8080")

	logger.Info("Server starting", zap.String("port", port))
	if err := http.ListenAndServe(port, router); err != nil {
		logger.Fatal("Failed to start the server", zap.Error(err))
	}
	// Run the server
	router.Run(port)
}
