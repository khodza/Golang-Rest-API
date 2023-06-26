package main

import (
	"khodza/rest-api/internal/app/handlers"
	"khodza/rest-api/internal/app/routers"
	"khodza/rest-api/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

func main() {
	//Loading env
	config.LoadEnv()
	//Init Logger
	Logger, err := config.CreateLogger()
	if err != nil {
		panic(err)
	}
	defer Logger.Sync()

	//Initialize DataBase
	err = config.InitDataBase()
	if err != nil {
		Logger.Fatal("Failed to connect to the database", zap.Error(err))
	}

	// Initialize dependencies
	handlersMap := config.InitDependencies(Logger)

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
		case "orders":
			orderHandler := handler.(*handlers.OrderHandler)
			routers.SetupOrderRouter(routeGroup, orderHandler)
		case "payments":
			paymentHandler := handler.(*handlers.PaymentHandler)
			routers.SetupPaymentRouter(routeGroup, paymentHandler)
		}
	}

	// Start the server
	port := ":8080"
	Logger.Info("Server starting", zap.String("port", port))
	if err := http.ListenAndServe(port, router); err != nil {
		Logger.Fatal("Failed to start the server", zap.Error(err))
	}
	// Run the server
	router.Run(":8080")

}
